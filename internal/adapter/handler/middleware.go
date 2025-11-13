package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
)

// validates JWT token from cookie
func AuthMiddleware(jwtConfig *config.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt_token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No token provided"})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
			c.Abort()
			return
		}

		claims := &domain.JWT{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtConfig.Secret), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token signature"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
			c.Abort()
			return
		}

		// store user claims in context for subsequent handlers
		c.Set("userEmail", claims.Email)
		c.Set("userRole", claims.Role)
		c.Set("userName", claims.Name)

		c.Next()
	}
}

// check if logged in user is the only one able to access resource
func CheckEmailParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "email parameter is required",
			})
			c.Abort()
			return
		}
		userEmail, exists := c.Get("userEmail")
		if !exists || userEmail != email {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: only loggedin user is able to access this resource"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// check if authenticated user had ADMIN_ROLE
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole") // sent from AuthMiddleware
		if !exists || (role.(domain.Role) != domain.ADMIN_ROLE) {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insuficient privilege"})
			c.Abort()
			return
		}
		c.Next()
	}
}
