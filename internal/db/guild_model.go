package db

// GuildModel is the database model for Discord guilds.
type GuildModel struct {
	BanThreshold int    `json:"ban_threshold"`
	Prefix       string `json:"prefix"`
	Gotcha       bool   `json:"gotcha"`
	MemberRole   string `json:"member_role"`
}
