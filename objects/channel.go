package objects

import "time"

// Snowflake discord ID
type Snowflake string


// Channel a channel object
type Channel struct {
	// ID channel ID
	ID Snowflake `json:"id"`
	// Type channel type
	TypeInt uint8 `json:"type"`
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

// Type turns TypeInt to string
func (c Channel) Type() string {
	switch c.TypeInt {
	case 0:
		return "text"
	case 1:
		return "dm"
	case 2:
		return "voice"
	case 3:
		return "group"
	case 4:
		return "category"
	case 5:
		return "news"
	case 6:
		return "store"
	}
	return ""
}
