package user

type userProfile struct {
	Email        string `json:"email"`
	UUID         string `json:"uuid"`
	ProfileImage string `json:"profileImage"`
	Nickname     string `json:"nickname"`
}
