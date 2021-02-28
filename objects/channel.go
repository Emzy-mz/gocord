package objects

import (
	"encoding/json"
	"time"
)

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
	Mentions          []User             `json:"mentions"`
	MentionedRoles    []Role             `json:"mention_roles"`
	MentionedChannels []ChannelMention   `json:"mention_channels,omitempty"`
	Attachments       []Attachment       `json:"attachments"`
	Embeds            []Embed            `json:"embeds"`
	Reactions         []Reaction         `json:"reactions,omitempty"`
	Nonce             json.Number        `json:"nonce,omitempty"`
	IsPinned          bool               `json:"pinned"`
	WebhookID         Snowflake          `json:"webhook_id,omitempty"`
	Type              MessageType        `json:"type"`
	Activity          MessageActivity    `json:"activity,omitempty"`
	Application       MessageApplication `json:"application,omitempty"`
	MessageReference  MessageReference   `json:"message_reference,omitempty"`
	Flags             MessageFlag        `json:"flags,omitempty"`
	Stickers          []Sticker          `json:"stickers,omitempty"`
	ReferencedMessage *Message           `json:"referenced_message,omitempty"`
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

// MessageType specifies the Message Type
type MessageType uint8

// Message Types
const (
	MessageTypeDefault                           MessageType = 0
	MessageTypeRecipientAdd                      MessageType = 1
	MessageTypeRecipientRemove                   MessageType = 2
	MessageTypeCall                              MessageType = 3
	MessageTypeChannelNameChange                 MessageType = 4
	MessageTypeChannelIconChange                 MessageType = 5
	MessageTypeChannelPinnedMessage              MessageType = 6
	MessageTypeGuildMemberJoin                   MessageType = 7
	MessageTypeUserPremiumGuildSubscription      MessageType = 8
	MessageTypeUserPremiumGuildSubscriptionTier1 MessageType = 9
	MessageTypeUserPremiumGuildSubscriptionTier2 MessageType = 10
	MessageTypeUserPremiumGuildSubscriptionTier3 MessageType = 11
	MessageTypeChannelFollowAdd                  MessageType = 12
	MessageTypeGuildDiscoveryDisqualified        MessageType = 14
	MessageTypeGuildDiscoveryRequalified         MessageType = 15
	MessageTypeReply                             MessageType = 19
	MessageTypeApplicationCommand                MessageType = 20
)

// MessageActivity Message Activity Object
type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyID string              `json:"party_id,omitempty"`
}

// MessageActivityType specifies the Message Activity Type
type MessageActivityType int

// Message Activity Types
const (
	MessageActivityJoin        MessageActivityType = 1
	MessageActivitySpectate    MessageActivityType = 2
	MessageActivityListening   MessageActivityType = 3
	MessageActivityJoinRequest MessageActivityType = 5
)

// MessageApplication Message Application Object
type MessageApplication struct {
	ID          Snowflake `json:"id"`
	CoverImage  string    `json:"cover_image,omitempty"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Name        string    `json:"name"`
}

// MessageReference Message Reference Object
type MessageReference struct {
	MessageID       Snowflake `json:"message_id,omitempty"`
	ChannelID       Snowflake `json:"channel_id,omitempty"`
	GuildID         Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool      `json:"fail_if_not_exists,omitempty"`
}

// MessageFlag specifies Message Flags
type MessageFlag int

// Message Flags
const (
	MessageFlagCrossposted          MessageFlag = 1 << 0
	MessageFlagIsCrosspost                      = 1 << 1
	MessageFlagSuppressEmbeds                   = 1 << 2
	MessageFlagSourceMessageDeleted             = 1 << 3
	MessageFlagUrgent                           = 1 << 4
)

// Sticker A Message Sticker
type Sticker struct {
	ID               Snowflake         `json:"id"`
	PackID           Snowflake         `json:"pack_id"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	Tags             string            `json:"tags,omitempty"`
	AssetHash        string            `json:"asset"`
	PreviewAssetHash string            `json:"preview_asset"`
	FormatType       StickerFormatType `json:"format_type"`
}

// StickerFormatType specifies the Sticker Format Type
type StickerFormatType uint8

// Sticker Format Types
const (
	StickerFormatTypePNG    StickerFormatType = 1
	StickerFormatTypeAPNG   StickerFormatType = 2
	StickerFormatTypeLOTTIE StickerFormatType = 3
)

// FollowedChannel Followed Channel Object
type FollowedChannel struct {
	ChannelID Snowflake `jsonn:"channel_id"`
	WebhookID Snowflake `json:"webhook_id"`
}

// AllowedMentions Allowed Mentions Object
type AllowedMentions struct {
	Parse       []AllowedMentionType `json:"parse"`
	Roles       []Snowflake          `json:"roles"`
	Users       []Snowflake          `json:"users"`
	RepliedUser bool                 `json:"replied_user"`
}

// AllowedMentionType specifies Allowed Mention Types
type AllowedMentionType string

// Allowed Mention Types
const (
	AllowedMentionTypeRoleMentions     AllowedMentionType = "roles"
	AllowedMentionTypeUserMentions     AllowedMentionType = "users"
	AllowedMentionTypeEveryoneMentions AllowedMentionType = "everyone"
)

// ToReadableTime turns the ISO8601 timestamp to something that is easier to read
func (t Timestamp) ToReadableTime() string {
	ts, _ := time.Parse("2006-01-02T15:04:05", string(t))
	return ts.String()
}
