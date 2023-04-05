//go:build parallel
// +build parallel

package model_test

func TestUser_HasPass(t *testing.T) {
	cases := map[string]struct {
		given interface{}
		want  interface{}
	}{}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := 1
			assert.Equal(t, c.want, got)
		})
	}
}
