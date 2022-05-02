package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		args := New("-p", "8080")
		port := args.Int("p")
		args.Parse()
		assert.Equal(t, 8080, *port)
	})
	t.Run("string", func(t *testing.T) {
		args := New("-d", "/usr/logs")
		dir := args.String("d")
		args.Parse()
		assert.Equal(t, "/usr/logs", *dir)
	})
	t.Run("multi-option", func(t *testing.T) {
		args := New("-l", "-p", "8080", "-d", "/usr/logs")
		logging := args.Bool("l")
		port := args.Int("p")
		dir := args.String("d")
		args.Parse()
		assert.True(t, *logging)
		assert.Equal(t, 8080, *port)
		assert.Equal(t, "/usr/logs", *dir)
	})
	// sad path:
	// TODO: - int -p/ -p 8080 8081
	// TODO: - string -d/ -d /usr/logs /usr/vars
	// default value:
	// TODO: -int :0
	// TODO: - string ""

}

func TestArgsBool(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		value bool
		err   error
	}{
		{"bool", []string{"-l"}, true, nil},
		{"bool default false", []string{}, false, nil},
		{"bool no extra arg", []string{"-l", "t"}, false, ErrTooManyArgs},
		{"bool no extra 2 arg", []string{"-l", "t", "f"}, false, ErrTooManyArgs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := New(tt.args...)
			logging := args.Bool("l")

			err := args.Parse()

			assert.Equal(t, tt.err, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.value, *logging)
		})

	}
}
