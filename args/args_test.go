package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
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

func TestArgsInt(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		value int
		err   error
	}{
		{"int", []string{"-p", "8080"}, 8080, nil},
		{"int default 0", []string{}, 0, nil},
		{"int no arg", []string{"-p"}, 0, ErrNoArg},
		{"int too many arg", []string{"-p", "8080", "8081"}, 0, ErrTooManyArgs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := New(tt.args...)
			port := args.Int("p")

			err := args.Parse()

			assert.Equal(t, tt.err, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.value, *port)
		})
	}
}

func TestArgsString(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		value string
		err   error
	}{
		{"string", []string{"-d", "/usr/logs"}, "/usr/logs", nil},
		{"string default ''", []string{}, "", nil},
		{"string no arg", []string{"-d"}, "", ErrNoArg},
		{"string too many arg", []string{"-d", "/usr/logs", "/usr/vars"}, "", ErrTooManyArgs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := New(tt.args...)
			dir := args.String("d")

			err := args.Parse()

			assert.Equal(t, tt.err, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.value, *dir)
		})
	}
}

func TestArgsStringList(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		value []string
		err   error
	}{
		{"string list", []string{"-g", "this", "is"}, []string{"this", "is"}, nil},
		{"string list default", []string{}, []string{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := New(tt.args...)
			li := args.StringList("g")

			err := args.Parse()

			assert.Equal(t, tt.err, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.value, *li)
		})
	}
}

func TestArgsIntList(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		value []int
		err   error
	}{
		{"int list", []string{"-i", "1", "2", "-3", "5"}, []int{1, 2, -3, 5}, nil},
		{"int list default", []string{}, []int{}, nil},
		{"int list invalid", []string{"-i", "a"}, []int{}, ErrInvalidArg},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := New(tt.args...)
			// use -i instead of the confusing -d flag
			li := args.IntList("i")

			err := args.Parse()

			assert.Equal(t, tt.err, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.value, *li)
		})
	}
}
