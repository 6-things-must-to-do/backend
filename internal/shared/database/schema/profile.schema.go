package schema

// Profile ...
type Profile struct {
	Provider     string // google | apple
	AppID        string `json:"-"` // hashedAppId
	ProfileImage string // Image URL
	Nickname     string
}

// TaskAlertSetting ...
type TaskAlertSetting struct {
	Hour   int `json:"hour" form:"hour" binding:"min=0,max=23"`
	Minute int `json:"minute" form:"minute" binding:"min=0,max=59"`
	Offset int `json:"offset" form:"offset" binding:"required"`
}

// ProfileWithSetting ...
type ProfileWithSetting struct {
	Profile
	TaskAlertSetting TaskAlertSetting `json:"taskAlertSetting,omitempty" dynamo:",set,omitempty"`
}

// ProfileSchema ...
type ProfileSchema struct {
	Key // USER#uuid PROFILE#email
	ProfileWithSetting
}
