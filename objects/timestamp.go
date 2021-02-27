package objects

// Timestamp a timestamp string
type Timestamp string

// ActivityTimestamps The timestamps at which an activity started and/or ended
type ActivityTimestamps struct {
	Start uint64 `json:"start,omitempty"`
	End uint64 `json:"end,omitempty"`
}
