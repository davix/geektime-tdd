package args

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrTooManyArgs = errors.New("too many arguments")
	ErrNoArg       = errors.New("no argument")
)

type Args struct {
	args []string
	res  map[string]any
}

func New(args ...string) Args {
	return Args{
		args: args,
		res:  make(map[string]any),
	}
}

func (a Args) Parse() error {
	for i, arg := range a.args {
		if result, ok := a.res[arg]; ok {
			switch v := result.(type) {
			case *bool:
				if i != len(a.args)-1 && !strings.HasPrefix(a.args[i+1], "-") {
					return ErrTooManyArgs
				}
				*v = true
			case *int:
				if i == len(a.args)-1 || strings.HasPrefix(a.args[i+1], "-") {
					return ErrNoArg
				}
				if i < len(a.args)-2 && !strings.HasPrefix(a.args[i+2], "-") {
					return ErrTooManyArgs
				}
				value, err := strconv.Atoi(a.args[i+1])
				if err != nil {
					continue
				}
				*v = value
			case *string:
				if i == len(a.args)-1 || strings.HasPrefix(a.args[i+1], "-") {
					return ErrNoArg
				}
				if i < len(a.args)-2 && !strings.HasPrefix(a.args[i+2], "-") {
					return ErrTooManyArgs
				}
				*v = a.args[i+1]
			}
		}
	}
	return nil
}

func (a Args) Bool(name string) *bool {
	var result bool
	a.res["-"+name] = &result
	return &result
}

func (a Args) Int(name string) *int {
	var result int
	a.res["-"+name] = &result
	return &result
}

func (a Args) String(name string) *string {
	var result string
	a.res["-"+name] = &result
	return &result
}
