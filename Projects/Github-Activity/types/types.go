package types

type Repository struct {
	Id      string
	Type    string
	Payload Payload
	Repo    Repo
}

type Repo struct {
	Name string
}

type Payload struct {
	Size int
	Action string
	Ref_Type string
}

