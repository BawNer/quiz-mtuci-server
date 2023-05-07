package entity

type Task struct {
	Name string `json:"name" binding:"required" gorm:"name"`
}

type NewTaskRequest struct {
	Task
}

type NewTaskResponse struct {
	Status string `json:"status"`
}

type TasksResponse struct {
	Data []*Task `json:"data"`
	Meta Meta    `json:"meta"`
}

type Meta struct {
	TotalCount int `json:"totalCount"`
}
