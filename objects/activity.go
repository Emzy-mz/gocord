package objects

// ActivityFlags Flags describing what the payload includes
type ActivityFlags int

// User Flags
const (
	Instance    ActivityFlags = 1 << 0
	Join                      = 1 << 1
	Spectate                  = 1 << 2
	JoinRequest               = 1 << 3
	Sync                      = 1 << 4
	Play                      = 1 << 5
)

// Activity Activity Object
type Activity struct {
	Name          string             `json:"name"`
	TypeAsInt     uint8              `json:"type"`
	URL           string             `json:"url,omitempty"`
	AddedAt       uint32             `json:"created_at"`
	Timestamps    ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID Snowflake          `json:"application_id,omitempty"`
	Details       string             `json:"details,omitempty"`
	PartyStatus   string             `json:"state,omitempty"`
	Emoji         Emoji              `json:"emoji,omitempty"`
	Party         ActivityParty      `json:"party,omitempty"`
	Assets        ActivityAsset      `json:"assets,omitempty"`
	Secrets       ActivitySecret     `json:"secrets,omitempty"`
	IsInstanced   bool               `json:"instance,omitempty"`
	Flags         ActivityFlags      `json:"flags,omitempty"`
}

// ActivityType returns a string from TypeAsInt
func (a Activity) ActivityType() string {
	switch a.TypeAsInt {
	case 0:
		return "game"
	case 1:
		return "streaming"
	case 2:
		return "listening"
	case 4:
		return "custom"
	case 5:
		return "competing"
	}
	return ""
}
