//go:build parallel
// +build parallel

package model_test

import (
	"testing"

	"github.com/api-sample/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_HashPass(t *testing.T) {
	cases := map[string]struct {
		email string
		pass  string
		want  string
	}{
		"success": {
			email: "sample@example.com",
			pass:  "hoge",
			want:  "dc02fb7f8721aae5ffb1a6459605e815",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := model.HashPass(c.email, c.pass)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestUser_ValidPass(t *testing.T) {
	cases := map[string]struct {
		m     model.User
		email string
		pass  string
		want  bool
	}{
		"success": {
			m: model.User{
				Email:    "sample@example.com",
				Password: "dc02fb7f8721aae5ffb1a6459605e815",
			},
			pass: "hoge",
			want: true,
		},
		"failed_model": {
			m: model.User{
				Email:    "sample@example.com",
				Password: "aaaaaaaaaaaaa",
			},
			pass: "hoge",
			want: false,
		},
		"failed_input": {
			m: model.User{
				Email:    "sample@example.com",
				Password: "dc02fb7f8721aae5ffb1a6459605e815",
			},
			pass: "fuga",
			want: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := c.m.ValidPass(c.pass)
			assert.Equal(t, c.want, got)
		})
	}
}
