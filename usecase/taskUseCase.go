package usecase

import (
	"Clean_Architecture/domain"
	"context"
	"time"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUseCase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (taskUC *taskUsecase) Create(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, taskUC.contextTimeout)
	defer cancel()

	return taskUC.taskRepository.Create(ctx, task)
}

func (taskUC *taskUsecase) FetchByUserID(c context.Context, taskID string) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, taskUC.contextTimeout)
	defer cancel()

	return taskUC.taskRepository.FetchByUserID(ctx, taskID)
}
