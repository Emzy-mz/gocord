package objects

type sysChannelFlags int

// User Flags
const (
	SuppressJoinNotifications    sysChannelFlags = 1 << 0
	SuppressPremiumSubscriptions                 = 1 << 1
)

// Guild Discord Server
type Guild struct {
	ID                          Snowflake        `json:"id"`
	Name                        string           `json:"name"`
	Icon                        string           `json:"icon"`
	IconHash                    string           `json:"icon_hash,omitempty"`
	Splash                      string           `json:"splash"`
	DiscoverySplash             string           `json:"discovery_splash"`
	IsOwner                     bool             `json:"owner,omitempty"`
	OwnerID                     Snowflake        `json:"owner_id"`
	Permissions                 string           `json:"permissions,omitempty"`
	Region                      string           `json:"region,omitempty"`
	AfkChannelID                Snowflake        `json:"afk_channel_id"`
	AfkTimeOut                  uint16           `json:"afk_timeout"`
	IsWidgetEnabled             bool             `json:"widget_enabled,omitempty"`
	WidgetChannelID             Snowflake        `json:"widget_channel_id,omitempty"`
	VerificationLevel           uint8            `json:"verification_level"`
	DefaultMessageNotifications uint8            `json:"default_message_notifications"`
	ExplicitContentFilterLevel  uint8            `json:"explicit_content_filter"`
	Roles                       []Role           `json:"roles"`
	Emojis                      []Emoji          `json:"emojis"`
	Features                    []string         `json:"features"`
	MFALevel                    uint8            `json:"mfa_level"`
	ApplicationID               Snowflake        `json:"application_id"`
	SystemChannelID             Snowflake        `json:"system_channel_id"`
	SystemChannelFlags          sysChannelFlags  `json:"system_channel_flags"`
	RulesChannelID              Snowflake        `json:"rules_channel_id"`
	CreatedAt                   Timestamp        `json:"joined_at,omitempty"`
	IsLarge                     bool             `json:"large,omitempty"`
	IsUnavailable               bool             `json:"unavailable,omitempty"`
	MemberCount                 uint32           `json:"member_count,oomitempty"`
	VoiceStates                 []VoiceState     `json:"voice_states,omitempty"`
	Members                     []GuildMember    `json:"members,omitempty"`
	Channels                    []Channel        `json:"channels"`
	Presences                   []PresenceUpdate `json:"presences,omitempty"`
	MaxPresences                uint16           `json:"max_presences,omitempty"`
	MaxMembers                  uint32           `json:"max_members,omitempty"`
	VanityURLCode               string           `json:"vanity_url_code"`
	Description                 string           `json:"description"`
	BannerHash                  string           `json:"banner"`
	ServerBoostLevel            uint8            `json:"premium_tier"`
	ServerBoostCount            uint32           `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string           `json:"preferred_locale"`
	PublicUpdatesChannelID      Snowflake        `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        uint32           `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount      uint32           `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    uint32           `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               WelcomeScreen    `json:"welcome_screen,omitempty"`
}

// GuildMember Member of a server
type GuildMember struct {
	User          User        `json:"user,omitempty"`
	Nick          string      `json:"nick,oomitempty"`
	Roles         []Snowflake `json:"roles"`
	JoinedAt      Timestamp   `json:"joined_at"`
	BoostingSince Timestamp   `json:"premium_since,omitempty"`
	IsDeaf        bool        `json:"deaf"`
	IsMute        bool        `json:"mute"`
	IsPending     bool        `json:"pending,omitempty"`
	Permissions   string      `json:"permissions,omitempty"`
}

// VanityURL returns the complete vanity URL for the Guild
func (g Guild) VanityURL() string {
	return "discord.gg/" + g.VanityURLCode
}

// WelcomeScreen Welcome Screen of the Server
type WelcomeScreen struct {
	Description     string                 `json:"description"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

// WelcomeScreenChannel Channel of the Welcome Screen
type WelcomeScreenChannel struct {
	ID          Snowflake `json:"channel_id"`
	Description string    `json:"description"`
	EmojiID     Snowflake `json:"emoji_id"`
	EmojiName   string    `json:"emoji_name"`
}
