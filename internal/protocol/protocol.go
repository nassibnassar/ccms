package protocol

// client request to CCMS server
type Request struct {
	Commands string `json:"commands"` // one or more commands
}
