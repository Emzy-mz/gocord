package objects

// Integration an integration
type Integration struct {
	ID              Snowflake                 `json:"id"`
	Name            string                    `json:"name"`
	Type            string                    `json:"type"`
	IsEnabled       bool                      `json:"enabled"`
	IsSyncing       bool                      `json:"syncing,omitempty"`
	RoleID          Snowflake                 `json:"role_id,omitempty"`
	EnableEmoticons bool                      `json:"enable_emoticons,omitempty"`
	ExpireBehavior  IntegrationExpireBehavior `json:"expire_behavior,omitempty"`
	// todo: fix int type
	ExpireGracePeriod int                `json:"expire_grace_period,omitempty"`
	User              User               `json:"user,omitempty"`
	Account           IntegrationAccount `json:"account"`
	SyncedAt          Timestamp          `json:"synced_at,omitempty"`
	// todo: fix int type
	SubscriberCount int  `json:"subscriber_count,omitempty"`
	IsRevoked       bool `json:"revoked,omitempty"`
	Application IntegrationApplication `json:"application,omitempty"`
}

// IntegrationExpireBehavior Expire Behavior type
type IntegrationExpireBehavior uint8

// Integration Expire Behavior types
const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = 0
	IntegrationExpireBehaviorKick       IntegrationExpireBehavior = 1
)

// IntegrationAccount account for an integration
type IntegrationAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// IntegrationApplication application for an integration
type IntegrationApplication struct {
	ID          Snowflake `json:"id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Summary     string    `json:"summary"`
	Bot         User      `json:"bot,omitempty"`
}
