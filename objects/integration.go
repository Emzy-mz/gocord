package objects

// Integration an integration
type Integration struct {
	ID              Snowflake `json:"id"`
	Name            string    `json:"name"`
	Type            string    `json:"type"`
	IsEnabled       bool      `json:"enabled"`
	IsSyncing       bool      `json:"syncing,omitempty"`
	RoleID          Snowflake `json:"role_id,omitempty"`
	EnableEmoticons bool      `json:"enable_emoticons,omitempty"`
	
}
