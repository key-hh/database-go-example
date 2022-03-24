package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-database/ent/entgen"
)

func TestORMRepository(t *testing.T) {
	or := ORMRepository{}

	ctx := context.Background()

	err := or.Init(ctx)
	or.client = or.client.Debug()
	assert.NoError(t, err)

	defer or.Close()

	err = or.Create(ctx)
	assert.NoError(t, err)

	res, err := or.Get(ctx, "testb1")
	assert.NoError(t, err)

	u := res.(*entgen.User)
	assert.Equal(t, u.Age, 11)

	res, err = or.List(ctx)
	assert.NoError(t, err)

	us := res.([]*entgen.User)
	assert.Equal(t, 1, len(us))

	res, err = or.ListX(ctx, 10)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(us))
}
