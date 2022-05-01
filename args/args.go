package args

type Args []string

func (a Args) Parse() {

}

func (a Args) Bool(name string) *bool {
	return nil
}

func (a Args) Int(name string) *int {
	return nil
}

func (a Args) String(name string) *string {
	return nil
}
