package api

type CommandRequest struct {
	Commands []string `json:"commands"`
}

type CommandResponse struct {
	Fields []FieldDescription `json:"fields"`
	Data   []DataRow          `json:"data"`
}

type FieldDescription struct {
	Name string `json:"name"`
	// DataType int
}

type DataRow struct {
	Values []string `json:"values"`
}
