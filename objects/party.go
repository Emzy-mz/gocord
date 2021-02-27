package objects

// ActivityParty Activity Party
type ActivityParty struct {
	ID string `json:"id,omitempty"`
	// Size first integer is the current size, the second integer is the maximum size
	Size [2]int `json:"size,omitempty"`
}

// ActivityAsset Activity Asset 
type ActivityAsset struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText string `json:"small_text,omitempty"`
}

// ActivitySecret Activity Secret
type ActivitySecret struct {
	Join string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match string `json:"match,omitempty"`
}
