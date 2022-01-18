package functionaloption

import (
	"context"

	"gorm.io/gorm"
)

type Task struct {
	Id    int
	Title string
}

type DBOption func(*gorm.DB) *gorm.DB

type ITaskRepo interface {
	WithById(id int) DBOption
	WithByTitle(title string) DBOption
	GetTask(ctx context.Context, opts ...DBOption) (*Task, error)
}
