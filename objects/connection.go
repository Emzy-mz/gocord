package objects

// Connection Connection Object
type Connection struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name"`
	Type                string        `json:"type"`
	IsRevoked           bool          `json:"revoked,omitempty"`
	Integrations        []Integration `json:"integrations,omitempty"`
	IsVerified          bool          `json:"verified"`
	IsFriendSyncEnabled bool          `json:"friend_sync"`
	ShowActivity        bool          `json:"show_activity"`
	Visibility          uint8         `json:"visibility"`
}
