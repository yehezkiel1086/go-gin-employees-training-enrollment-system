package domain

type TrainingStatistics struct {
	TotalAvailableTrainings int64 `json:"total_available_trainings"`
	EnrolledTrainings       int64 `json:"enrolled_trainings"`
	CompletedTrainings      int64 `json:"completed_trainings"`
}

type TrainingsByCategory struct {
	CategoryName   string `json:"category_name"`
	TotalTrainings int64  `json:"total_trainings"`
}
