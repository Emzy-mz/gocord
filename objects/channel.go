package objects

import "time"

// Snowflake discord ID
type Snowflake string

// ChannelType Specifies the Channel Type
type ChannelType int

// Specifies the Channel Types
const (
	GuildText     ChannelType = 0
	DM            ChannelType = 1
	GuildVoice    ChannelType = 2
	GroupDM       ChannelType = 3
	GuildCategory ChannelType = 4
	GuildNews     ChannelType = 5
	GuildStore    ChannelType = 6
)

// Channel a channel object
type Channel struct {
	// ID channel ID
	ID      Snowflake   `json:"id"`
	TypeInt ChannelType `json:"type"`
	// GuildID server ID
	GuildID Snowflake `json:"guild_id,omitempty"`

	Position         int         `json:"position,omitempty"`
	PermOverwrites   []Overwrite `json:"permission_overwrites,omitempty"`
	Name             string      `json:"name,omitempty"`
	Topic            string      `json:"topic,omitempty"`
	IsNSFW           bool        `json:"nsfw,omitempty"`
	LastMessageID    Snowflake   `json:"last_message_id,omitempty"`
	BitRate          uint32      `json:"bitrate,omitempty"`
	UserLimit        uint16      `json:"user_limit,omitempty"`
	RateLimitPerUser uint16      `json:"rate_limit_per_user,omitempty"`
	Recipients       []User      `json:"recipients,omitempty"`
	Icon             string      `json:"icon,omitempty"`
	OwnerID          Snowflake   `json:"owner_id,omitempty"`
	AppID            Snowflake   `json:"application_id,omitempty"`
	ParentID         Snowflake   `json:"parent_id,omitempty"`
	LastPinTimeStamp string      `json:"last_pin_timestamp,omitempty"`
}

// ToReadableTime turns the ISO8601 timestamp to something that is easier to read
func (t Timestamp) ToReadableTime() string {
	ts, _ := time.Parse("2006-01-02T15:04:05", string(t))
	return ts.String()
}
