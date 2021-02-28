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

// Message Text Message
type Message struct {
	ID              Snowflake   `json:"id"`
	ChannelID       Snowflake   `json:"channel_id"`
	GuildID         Snowflake   `json:"guild_id,omitemptyd"`
	Author          User        `json:"author"`
	Member          GuildMember `json:"member,omitempty"`
	Content         string      `json:"content"`
	Timestamp       Timestamp   `json:"timestamp"`
	EditedTimestamp Timestamp   `json:"edited_timestamp"`
	Tts             bool        `json:"tts"`
	MentionEveryone bool        `json:"mention_everyone"`
	//todo check docs
	Mentions          []User           `json:"mentions"`
	MentionedRoles    []Role           `json:"mention_roles"`
	MentionedChannels []ChannelMention `json:"mention_channels,omitempty"`
	Attachments       []Attachment     `json:"attachments"`
	Embeds            []Embed          `json:"embeds"`
	Reactions         []Reaction       `json:"reactions,omitempty"`
}

// ChannelMention Channel Mention Object
type ChannelMention struct {
	ID          Snowflake   `json:"id"`
	GuildID     Snowflake   `json:"guild_id"`
	ChannelType ChannelType `json:"type"`
	Name        string      `json:"name"`
}

// Attachment File
type Attachment struct {
	ID       Snowflake `json:"id"`
	FileName string    `json:"filename"`
	Size     uint32    `json:"size"`
	URL      string    `json:"url"`
	ProxyURL string    `json:"proxy_url"`
	Height   uint16    `json:"height"`
	Width    uint16    `json:"width"`
}

// Embed Message Embed
type Embed struct {
	Title       string         `json:"title,omitempty"`
	Type        EmbedType      `json:"type,omitempty"`
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitempty"`
	Timestamp   Timestamp      `json:"timestamp,omitempty"`
	Color       uint32         `json:"color,omitempty"`
	Footer      EmbedFooter    `json:"footer,omitempty"`
	Image       EmbedImage     `json:"image,omitempty"`
	Thumbnail   EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       EmbedVideo     `json:"video,omitempty"`
	Provider    EmbedProvider  `json:"provider,omitempty"`
	Author      EmbedAuthor    `json:"author,omitempty"`
	Fields      []EmbedField   `json:"fields,omitempty"`
}

// EmbedType specifies the Embed type
type EmbedType string

// Embed Types
const (
	EmbedTypeRich    EmbedType = "rich"
	EmbedTypeImage   EmbedType = "image"
	EmbedTypeVideo   EmbedType = "video"
	EmbedTypeGifV    EmbedType = "gifv"
	EmbedTypeArticle EmbedType = "article"
	EmbedTypeLink    EmbedType = "link"
)

// EmbedFooter Footer of the Embed
type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedImage Image of the Embed
type EmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   uint16 `json:"height,omitempty"`
	Width    uint16 `json:"width,omitempty"`
}

// EmbedThumbnail Thumbnail of the Embed
type EmbedThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   uint16 `json:"height,omitempty"`
	Width    uint16 `json:"width,omitempty"`
}

// EmbedVideo Video of the Embed
type EmbedVideo struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   uint16 `json:"height,omitempty"`
	Width    uint16 `json:"width,omitempty"`
}

// EmbedProvider Provider of the Embed
type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// EmbedAuthor Author of the Embed
type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedField Field of the Embed
type EmbedField struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	IsInline bool   `json:"inline,omitempty"`
}

// Reaction Reaction to a message
type Reaction struct {
	Count    uint32 `json:"count"`
	IReacted bool   `json:"me"`
	Emoji    Emoji  `json:"emoji"`
}

// ToReadableTime turns the ISO8601 timestamp to something that is easier to read
func (t Timestamp) ToReadableTime() string {
	ts, _ := time.Parse("2006-01-02T15:04:05", string(t))
	return ts.String()
}
