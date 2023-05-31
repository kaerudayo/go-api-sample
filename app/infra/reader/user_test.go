//go:build parallel
// +build parallel

package reader_test

import (
	"testing"
	"time"

	"github.com/api-sample/app/domain/model"
	"github.com/api-sample/app/infra"
	"github.com/api-sample/app/infra/reader"
	"github.com/stretchr/testify/assert"
)

func TestUser_FindByID(t *testing.T) {
	r := reader.NewUserQueryImpl(infra.DB)
	loadFixtures(t)
	cases := map[string]struct {
		ID   string
		want model.User
	}{
		"success": {
			ID: "user_1",
			want: model.User{
				ID:        "user_1",
				Name:      "kaeru",
				Password:  "01H1N1R83HJ7AM3H5GM71T22M6",
				Email:     "user_1@example.com",
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := r.FindByID(c.ID)
			assert.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestUser_FindByEmail(t *testing.T) {
	r := reader.NewUserQueryImpl(infra.DB)
	loadFixtures(t)
	cases := map[string]struct {
		email string
		want  model.User
	}{
		"success": {
			email: "user_1@example.com",
			want: model.User{
				ID:        "user_1",
				Name:      "kaeru",
				Password:  "01H1N1R83HJ7AM3H5GM71T22M6",
				Email:     "user_1@example.com",
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := r.FindByEmail(c.email)
			assert.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}
