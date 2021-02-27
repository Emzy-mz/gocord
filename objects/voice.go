package objects

// VoiceState Connection Status of a user
type VoiceState struct {
	GuildID   Snowflake   `json:"guild_id,omitempty"`
	ChannelID Snowflake   `json:"channel_id"`
	UserID    Snowflake   `json:"user_id"`
	Member    GuildMember `json:"member,omitempty"`
	SessionID string `json:"session_id"`
	IsServerDeafened bool `json:"deaf"`
	IsServerMuted bool `json:"mute"`
	IsDeafened bool `json:"self_deaf"`
	IsMuted bool `json:"self_mute"`
	IsStreaming bool `json:"self_stream,omitempty"`
	HasCameraOn bool `json:"self_video"`
	IsSuppressedByClient bool `json:"suppress"`
}
