package protocol

type CommandRequest struct {
	Command string `json:"command"`
}

type CommandResponse struct {
	Status  string             `json:"status"`
	Message string             `json:"message,omitempty"`
	Fields  []FieldDescription `json:"fields,omitempty"`
	Data    []DataRow          `json:"data,omitempty"`
}

type FieldDescription struct {
	Name string `json:"name"`
	// DataType int
}

type DataRow struct {
	Values []any `json:"values"`
}
