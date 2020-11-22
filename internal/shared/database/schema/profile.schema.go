package schema

// Profile ...
type Profile struct {
	Provider     string `json:"-"`// google | apple
	AppID        string `json:"-"` // hashedAppId
	ProfileImage string `json:"profileImage"`// Image URL
	Nickname     string	`json:"nickname"`
}

// TaskAlertSetting ...
type TaskAlertSetting struct {
	Hour   int `json:"hour" form:"hour" binding:"min=0,max=23"`
	Minute int `json:"minute" form:"minute" binding:"min=0,max=59"`
	Offset int `json:"offset" form:"offset" binding:"required"`
}

type OpennessCollection struct {
	Account int `json:"account"`
	Task int `json:"task"`
	Record int `json:"record"`
}

// ProfileWithSetting ...
type ProfileWithSetting struct {
	Profile
	TaskAlertSetting TaskAlertSetting `json:"taskAlertSetting,omitempty" dynamo:",set,omitempty"`
}

// ProfileSchema ...
type ProfileSchemaWithSetting struct {
	Key // USER#uuid PROFILE#email
	ProfileWithSetting
}

type ProfileSchema struct {
	Key
	Profile
}
