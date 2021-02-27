package objects

// UserFlags flags of the User
type UserFlags int

// User Flags
const (
	UserFlagDiscordEmployee     UserFlags = 1 << 0
	UserFlagDiscordPartner                = 1 << 1
	UserFlagHypeSquadEvents               = 1 << 2
	UserFlagBugHunterLVL1                 = 1 << 3
	UserFlagHouseBravery                  = 1 << 6
	UserFlagHouseBrilliance               = 1 << 7
	UserFlagHouseBalance                  = 1 << 8
	UserFlagEarlySupporter                = 1 << 9
	UserFlagTeamUser                      = 1 << 10
	UserFlagSystem                        = 1 << 12
	UserFlagBugHunterLVL2                 = 1 << 14
	UserFlagVerifiedBot                   = 1 << 16
	UserFlagEarlyVerifiedBotDev           = 1 << 17
)

// User Discord User Object
type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Discrim     string    `json:"discriminator"`
	Avatar      string    `json:"avatar"`
	IsBot       bool      `json:"bot,omitempty"`
	IsSystem    bool      `json:"system,omitempty"`
	HasMFA      bool      `json:"mfa_enabled,omitempty"`
	Language    string    `json:"locale,omitempty"`
	IsVerified  bool      `json:"verified,omitempty"`
	EMail       string    `json:"email,omitempty"`
	Flags       UserFlags `json:"flags,omitempty"`
	NitroType   uint8     `json:"premium_type,omitempty"`
	PublicFlags UserFlags `json:"public_flags,omitempty"`
}

// Tag makes a tag from a user object (e.g. Schmenn#1088)
func (u User) Tag() string {
	return u.Username + "#" + u.Discrim
}

// Mention makes a mention string from a user object (e.g. <@489431305444917258>)
func (u User) Mention() string {
	return "<@" + u.ID + ">"
}
