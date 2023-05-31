//go:build serial
// +build serial

package writer_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testxx_xx(t *testing.T) {
	cases := map[string]struct {
		given interface{}
		want  interface{}
	}{}

	fmt.Println("----------------")
	fmt.Println("serial")
	fmt.Println("----------------")

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := 1
			assert.Equal(t, c.want, got)
		})
	}
}
