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

type Follow struct {
	PK string `json:"follower"`
	SK string `json:"followee"`
}

type Request struct {
	PK string `json:"-"` // REQ#TYPE#~
	SK string `json:"email"`
}