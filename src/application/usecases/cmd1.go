package usecases

type Cmd1 struct{}

func NewCmd1() *Cmd1 {
	return &Cmd1{}
}

func (c Cmd1) Execute() string {
	return "Executing cmd1..."
}
