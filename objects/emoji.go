package objects

// Emoji emoji object
type Emoji struct {
	ID             Snowflake `json:"id"`
	Name           string    `json:"name"`
	Roles          []Role    `json:"roles,omitempty"`
	User           User      `json:"user,omitempty"`
	RequiresColons bool      `json:"require_colons,omitempty"`
	IsManaged      bool      `json:"managed,omitempty"`
	IsAnimated     bool      `json:"animated,omitempty"`
	IsAvailable    bool      `json:"available,omitempty"`
}
