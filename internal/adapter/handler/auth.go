package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
)

type AuthHandler struct {
	svc port.AuthService
	conf *config.JWT
}

func InitAuthHandler(conf *config.JWT, svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
		conf: conf,
	}
}

type LoginReq struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and password are required",
		})
		return
	}

	user, err := ah.svc.Login(c, &domain.User{
		Email: req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create jwt token
	mySigningKey := []byte(ah.conf.Secret)

	tokenDuration, err := strconv.Atoi(ah.conf.Duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// create custom JWT claims
	claims := &domain.JWT{
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenDuration) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	maxAge := tokenDuration * 60

	c.SetCookie("jwt_token", ss, maxAge, "/", "", ah.conf.Env == "production", true)

	c.JSON(http.StatusOK, gin.H{
		"jwt_token": ss,
	})
}

func (ah *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt_token", "", -1, "/", "", ah.conf.Env == "production", true)

	c.JSON(http.StatusOK, gin.H{
		"message": "user logged out successfully",
	})
}
