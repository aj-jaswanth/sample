package todo

type Task struct {
	Identifier     string
	Action         string
	ListIdentifier string
}

type List struct {
	Identifier string
	Name       string
	Tasks      []Task
}
