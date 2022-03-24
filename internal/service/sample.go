package service

import (
	"context"
	"go-database/internal/repository"

	_ "github.com/go-sql-driver/mysql"
)

type SampleService struct {
	repo repository.IRepository
}

func NewSampleService(repo repository.IRepository) *SampleService {
	return &SampleService{repo: repo}
}

func (ss *SampleService) Create(ctx context.Context) error {
	return ss.repo.Create(ctx)
}

func (ss *SampleService) Get(ctx context.Context, id string) (interface{}, error) {
	return ss.repo.Get(ctx, id)
}

func (ss *SampleService) List(ctx context.Context) (interface{}, error) {
	return ss.repo.List(ctx)
}
