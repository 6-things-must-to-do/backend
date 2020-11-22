package schema

// Key ...
type Key struct {
	PK string `json:"-"`
	SK string `json:"-"`
}

type Openness struct {
	PK string	`json:"-"`
	SK string	`json:"-"`
}

type Request struct {
	PK string `json:"-"` // REQ#TYPE#~
	SK string `json:"email"`
}