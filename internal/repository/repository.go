package repository

import (
	"context"
)

type IRepository interface {
	Create(ctx context.Context) error
	List(ctx context.Context) (interface{}, error) // 다양한 테스트를 위해 interface타입으로 둠
	Get(ctx context.Context, id string) (interface{}, error)
}
