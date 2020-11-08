package auth

type loginRequired struct {
	ID       string `form:"id" binding:"required" json:"id"`
	Email    string `form:"email" binding:"required" json:"email"`
	Provider string `form:"provider" binding:"required" json:"provider"`
}

type loginDto struct {
	loginRequired
	Nickname     string `form:"nickname" json:"nickname"`
	ProfileImage string `form:"profileImage" json:"profileImage"`
}
