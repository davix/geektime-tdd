package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		args := New("-l")
		logging := args.Bool("l")
		args.Parse()
		assert.True(t, *logging)
	})
	t.Run("no bool", func(t *testing.T) {
		args := New()
		logging := args.Bool("l")
		args.Parse()
		assert.False(t, *logging)
	})
	//t.Run("multi-option", func(t *testing.T) {
	//	args := Args{"-l", "-p", "8080", "-d", "/usr/logs"}
	//	logging := args.Bool("l")
	//	port := args.Int("p")
	//	dir := args.String("d")
	//	args.Parse()
	//	assert.True(t, *logging)
	//	assert.Equal(t, 8080, *port)
	//	assert.Equal(t, "/user/logs", *dir)
	//})
}