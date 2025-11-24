package ccms

type Response struct {
	Status  string
	Fields  []FieldDescription
	Data    []DataRow
	Message string
}

type FieldDescription struct {
	Name string
	//Type int
}

type DataRow struct {
	Values []string
}
