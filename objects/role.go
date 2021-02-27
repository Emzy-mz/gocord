package objects

// Role role object
type Role struct {
	ID Snowflake `json:"id"`
	Name string `json:"name"`
	Color uint32 `json:"color"`
	IsHoisted bool `json:"hoisted"`
	Position uint8 `json:"position"`
	Permissions string `json:"permissions"`
	IsManaged bool `json:"managed"`
	IsMentionable bool `json:"mentionable"`
	Tags RoleTags `json:"tags,omitempty"`
}

// RoleTags tags of a role
type RoleTags struct {
	BotID Snowflake `json:"bot_id,omitempty"`
	IntegrationID Snowflake `json:"integration_id,omitempty"`
	// PremiumSubscriber `json:"premium_subscriber,omitempty"`
}
