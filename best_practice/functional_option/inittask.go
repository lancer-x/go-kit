package functionaloption

import "go-kit/entity"

type taskOpt func(*entity.Task)

func WithId(id int) taskOpt {
	return func(t *entity.Task) {
		t.Id = id
	}
}

func WithDesc(desc string) taskOpt {
	return func(t *entity.Task) {
		t.Desc = desc
	}
}

func WithCreator(creator string) taskOpt {
	return func(t *entity.Task) {
		t.Creator = creator
	}
}

// 假定任务必须有一个title
func NewTask(title string, opts ...taskOpt) *entity.Task {
	task := &entity.Task{
		Title: title,
	}
	for _, option := range opts {
		option(task)
	}
	return task
}
