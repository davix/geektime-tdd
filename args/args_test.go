package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	args := Args{"-l", "-p", "8080", "-d", "/usr/logs"}
	logging := args.Bool("l")
	port := args.Int("p")
	dir := args.String("d")
	args.Parse()
	assert.True(t, *logging)
	assert.Equal(t, 8080, *port)
	assert.Equal(t, "/user/logs", *dir)
}
