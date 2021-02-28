package objects

import "encoding/json"

// AuditLog Audit Log object
type AuditLog struct {
	Webhooks []Webhook `json:"webhooks"`
	Users    []User    `json:"user"`
	Entries []AuditLogEntry `json:"audit_log_entries"`
	Integrations []Integration `json:"integrations"`
}

// Webhook represents a webhook
type Webhook struct {
	ID            Snowflake `json:"id"`
	TypeInt       uint8     `json:"type"`
	GuildID       Snowflake `json:"guild_id,omitempty"`
	ChannelID     Snowflake `json:"channel_id"`
	User          User      `json:"user,omitempty"`
	Name          string    `json:"name"`
	Avatar        string    `json:"avatar"`
	Token         string    `json:"token,omitempty"`
	ApplicationID Snowflake `json:"application_id"`
}

// AuditLogEntry an entry from an audit log
type AuditLogEntry struct {
	TargetID     string                    `json:"target_id"`
	Changes      []AuditLogChange          `json:"changes,omitempty"`
	UserID       Snowflake                 `json:"user_id"`
	ID           Snowflake                 `json:"id"`
	ActionType   AuditLogEvent             `json:"action_type"`
	OptionalInfo OptionalAuditLogEntryInfo `json:"options,omitempty"`
	Reason       string                    `json:"reason,omitempty"`
}

// OptionalAuditLogEntryInfo contains optional info
type OptionalAuditLogEntryInfo struct {
	DeleteMemberDays string    `json:"delete_member_days,omitempty"`
	MembersRemoved   string    `json:"members_removed,omitempty"`
	ChannelID        Snowflake `json:"channel_id,omitempty"`
	MessageID        Snowflake `json:"message_id,omitempty"`
	Count            string    `json:"count,omitempty"`
	ID               Snowflake `json:"id,omitempty"`
	// Type "0" for "role" and "1" for "member"
	Type     string `json:"type,omitempty"`
	RoleName string `json:"role_name,omitempty"`
}

// AuditLogEvent specifies the event
type AuditLogEvent uint8

// Audit Log Events
const (
	AuditLogEventGuildUpdate            AuditLogEvent = 1
	AuditLogEventChannelGreate          AuditLogEvent = 10
	AuditLogEventChannelUpdate          AuditLogEvent = 11
	AuditLogEventChannelDelete          AuditLogEvent = 12
	AuditLogEventChannelOverwriteCreate AuditLogEvent = 13
	AuditLogEventChannelOverwriteUpdate AuditLogEvent = 14
	AuditLogEventChannelOverwriteDelete AuditLogEvent = 15
	AuditLogEventMemberKick             AuditLogEvent = 20
	AuditLogEventMemberPrune            AuditLogEvent = 21
	AuditLogEventMemberBanAdd           AuditLogEvent = 22
	AuditLogEventMemberBanRemove        AuditLogEvent = 23
	AuditLogEventMemberUpdate           AuditLogEvent = 24
	AuditLogEventMemberRoleUpdate       AuditLogEvent = 25
	AuditLogEventMemberMove             AuditLogEvent = 26
	AuditLogEventMemberDisconnect       AuditLogEvent = 27
	AuditLogEventBotAdd                 AuditLogEvent = 28
	AuditLogEventRoleCreate             AuditLogEvent = 30
	AuditLogEventRoleUpdate             AuditLogEvent = 31
	AuditLogEventRoleDelete             AuditLogEvent = 32
	AuditLogEventInviteCreate           AuditLogEvent = 40
	AuditLogEventInviteUpdate           AuditLogEvent = 41
	AuditLogEventInviteDelete           AuditLogEvent = 42
	AuditLogEventWebhookCreate          AuditLogEvent = 50
	AuditLogEventWebhookUpdate          AuditLogEvent = 51
	AuditLogEventWebhookDelete          AuditLogEvent = 52
	AuditLogEventEmojiCreate            AuditLogEvent = 60
	AuditLogEventEmojiUpdate            AuditLogEvent = 61
	AuditLogEventEmojiDelete            AuditLogEvent = 62
	AuditLogEventMessageDelete          AuditLogEvent = 72
	AuditLogEventMessageBulkDelete      AuditLogEvent = 73
	AuditLogEventMessagePin             AuditLogEvent = 74
	AuditLogEventMessageUnpin           AuditLogEvent = 75
	AuditLogEventIntegrationCreate      AuditLogEvent = 80
	AuditLogEventIntegrationUpdate      AuditLogEvent = 81
	AuditLogEventIntegrationDelete      AuditLogEvent = 82
)

// AuditLogChange The Change that happened
type AuditLogChange struct {
	New json.Number       `json:"new_value,omitempty"`
	Old json.Number       `json:"old_value,omitempty"`
	Key AuditLogChangeKey `json:"key"`
}

// AuditLogChangeKey used to specify the key in AuditLogChange
type AuditLogChangeKey string

// Audit Log Change Keys
const (
	AuditLogChangeKeyName                       AuditLogChangeKey = "name"
	AuditLogChangeKeyDescription                AuditLogChangeKey = "description"
	AuditLogChangeKeyIconHash                   AuditLogChangeKey = "icon_hash"
	AuditLogChangeKeySplashHash                 AuditLogChangeKey = "splash_hash"
	AuditLogChangeKeyDiscoverySplashHash        AuditLogChangeKey = "discovery_splash_hash"
	AuditLogChangeKeyBannerHash                 AuditLogChangeKey = "banner_hash"
	AuditLogChangeKeyOwnerID                    AuditLogChangeKey = "owner_id"
	AuditLogChangeKeyRegion                     AuditLogChangeKey = "region"
	AuditLogChangeKeyPreferredLocale            AuditLogChangeKey = "preferred_locale"
	AuditLogChangeKeyAfkChannelID               AuditLogChangeKey = "afk_channel_id"
	AuditLogChangeKeyAfkTimeout                 AuditLogChangeKey = "afk_timeout"
	AuditLogChangeKeyRulesChannelID             AuditLogChangeKey = "rules_channel_id"
	AuditLogChangeKeyPublicUpdatesChannelID     AuditLogChangeKey = "public_updates_channel_id"
	AuditLogChangeKeyMfaLevel                   AuditLogChangeKey = "mfa_level"
	AuditLogChangeKeyVerificationLevel          AuditLogChangeKey = "verification_level"
	AuditLogChangeKeyExplicitContentFilter      AuditLogChangeKey = "explicit_content_filter"
	AuditLogChangeKeyDefaultMessageNotification AuditLogChangeKey = "default_message_notifications"
	AuditLogChangeKeyVanityURLCode              AuditLogChangeKey = "vanity_url_code"
	AuditLogChangeKeyRoleAdd                    AuditLogChangeKey = "$add"
	AuditLogChangeKeyRoleRemove                 AuditLogChangeKey = "$remove"
	AuditLogChangeKeyPruneDeleteDays            AuditLogChangeKey = "prune_delete_days"
	AuditLogChangeKeyWidgetEnabled              AuditLogChangeKey = "widget_enabled"
	AuditLogChangeKeyWidgetChannelID            AuditLogChangeKey = "widget_channel_id"
	AuditLogChangeKeySystemChannelID            AuditLogChangeKey = "system_channel_id"
	AuditLogChangeKeyPosition                   AuditLogChangeKey = "position"
	AuditLogChangeKeyTopic                      AuditLogChangeKey = "topic"
	AuditLogChangeKeyBitrate                    AuditLogChangeKey = "bitrate"
	AuditLogChangeKeyPermissionOverwrite        AuditLogChangeKey = "permission_overwrites"
	AuditLogChangeKeyNSFW                       AuditLogChangeKey = "nsfw"
	AuditLogChangeKeyApplicationID              AuditLogChangeKey = "application_id"
	AuditLogChangeKeyRateLimitPerUser           AuditLogChangeKey = "rate_limit_per_user"
	AuditLogChangeKeyPermissions                AuditLogChangeKey = "permissions"
	AuditLogChangeKeyColor                      AuditLogChangeKey = "color"
	AuditLogChangeKeyHoist                      AuditLogChangeKey = "hoist"
	AuditLogChangeKeyMentionable                AuditLogChangeKey = "mentionable"
	AuditLogChangeKeyAllow                      AuditLogChangeKey = "allow"
	AuditLogChangeKeyDeny                       AuditLogChangeKey = "deny"
	AuditLogChangeKeyCode                       AuditLogChangeKey = "code"
	AuditLogChangeKeyChannelID                  AuditLogChangeKey = "channel_id"
	AuditLogChangeKeyInviterID                  AuditLogChangeKey = "inviter_id"
	AuditLogChangeKeyMaxUses                    AuditLogChangeKey = "max_uses"
	AuditLogChangeKeyUses                       AuditLogChangeKey = "uses"
	AuditLogChangeKeyMaxAge                     AuditLogChangeKey = "max_age"
	AuditLogChangeKeyTempoary                   AuditLogChangeKey = "temporary"
	AuditLogChangeKeyDeaf                       AuditLogChangeKey = "deaf"
	AuditLogChangeKeyMute                       AuditLogChangeKey = "mute"
	AuditLogChangeKeyNick                       AuditLogChangeKey = "nick"
	AuditLogChangeKeyAvatarHash                 AuditLogChangeKey = "avatar_hash"
	AuditLogChangeKeyID                         AuditLogChangeKey = "id"
	AuditLogChangeKeyType                       AuditLogChangeKey = "type"
	AuditLogChangeKeyEnableEmoticons            AuditLogChangeKey = "enable_emoticons"
	AuditLogChangeKeyExpireBehavior             AuditLogChangeKey = "expire_behavior"
	AuditLogChangeKeyExpireGracePeriod          AuditLogChangeKey = "expire_grace_period"
)

// GetType either returns "incoming" or "follower"
func (c Webhook) GetType() string {
	switch c.TypeInt {
	case 1:
		return "incoming"
	case 2:
		return "follower"
	}
	return ""
}
