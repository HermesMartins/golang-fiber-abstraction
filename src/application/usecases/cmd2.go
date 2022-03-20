package usecases

type Cmd2 struct{}

func NewCmd2() *Cmd2 {
	return &Cmd2{}
}

func (c Cmd2) Execute() string {
	return "Executing cmd2..."
}
