package args

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

func (a Args) Parse() {
	for _, arg := range a.args {
		if result, ok := a.res[arg]; ok {
			*(result.(*bool)) = true
		}
	}
}

func (a Args) Bool(name string) *bool {
	var result bool
	a.res["-"+name] = &result
	return &result
}

func (a Args) Int(name string) *int {
	return nil
}

func (a Args) String(name string) *string {
	return nil
}
