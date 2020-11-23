package schema

type Follow struct {
	PK string `json:"-"` // FOLLOWER#uuid
	SK string `json:"email"` // PROFILE#email
	ProfileUUID string `json:"uuid"` // target uuid
	FollowerEmail string `json:"-"` // Follower email
}

type RankRecord struct {
	RecordSchema
	UUID string `json:"uuid"`
}