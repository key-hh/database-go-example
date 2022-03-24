package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleService(t *testing.T) {
	s := NewSQLXRepository()

	err := s.Init()
	assert.NoError(t, err)
	defer s.Close()

	ctx := context.Background()

	err = s.Create(ctx)
	assert.NoError(t, err)

	res, err := s.Get(ctx, "e2")
	assert.NoError(t, err)
	p := res.(Person)
	assert.Equal(t, "f2", p.FirstName)

	res, err = s.List(ctx)
	assert.NoError(t, err)
	ps := res.([]Person)
	assert.Equal(t, 3, len(ps))
}
