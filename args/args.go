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
			values := valuesFromFollowing(a.args[i+1:])

			switch v := result.(type) {
			case *bool:
				if len(values) > 0 {
					return ErrTooManyArgs
				}
				*v = true
			case *int:
				if len(values) < 1 {
					return ErrNoArg
				}
				if len(values) > 1 {
					return ErrTooManyArgs
				}
				value, err := strconv.Atoi(values[0])
				if err != nil {
					continue
				}
				*v = value
			case *string:
				if len(values) < 1 {
					return ErrNoArg
				}
				if len(values) > 1 {
					return ErrTooManyArgs
				}
				*v = values[0]
			case *[]string:
				*v = append(*v, values...)
			}
		}
	}
	return nil
}

func valuesFromFollowing(args []string) []string {
	var values []string
	for _, a := range args {
		if strings.HasPrefix(a, "-") {
			break
		}
		values = append(values, a)
	}
	return values
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

func (a Args) StringList(name string) *[]string {
	result := make([]string, 0)
	a.res["-"+name] = &result
	return &result
}
