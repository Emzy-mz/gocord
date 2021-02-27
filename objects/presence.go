package objects

// PresenceUpdate Current State on a Guild
type PresenceUpdate struct {
	User         User         `json:"user"`
	GuildID      Snowflake    `json:"guild_id"`
	Status       string       `json:"status"`
	Activities   []Activity   `json:"activities"`
	ClientStatus ClientStatus `json:"client_status"`
}

// ClientStatus The Status of a User
type ClientStatus struct {
	Desktop string `json:"desktop,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	Web     string `json:"web,omitempty"`
}
