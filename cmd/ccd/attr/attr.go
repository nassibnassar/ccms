package attr

type AttrType int

const (
	NoType AttrType = iota
	SmallInt
	BigInt
	Text
	Boolean
	Property
)

type Attr struct {
	Name string
	Type AttrType
	Core bool
}

var AttrSlice = []Attr{
	{Name: "id", Type: BigInt, Core: true},
	{Name: "author", Type: Text, Core: true},
	{Name: "title", Type: Text, Core: true},
	{Name: "full_vendor_name", Type: Text, Core: true},
	{Name: "availability", Type: Text, Core: true},
	{Name: "library_holdings_count", Type: SmallInt, Core: true},
	{Name: "online_database_holdings_count", Type: SmallInt, Core: true},
	{Name: "vendor_holdings_count", Type: SmallInt, Core: true},
	{Name: "holdings_count", Type: SmallInt, Core: true},
	{Name: "decision", Type: Boolean, Core: false},
	{Name: "fund", Type: Property, Core: false},
	{Name: "track", Type: Property, Core: false},
}

var AttrMap map[string]Attr

func init() {
	AttrMap = make(map[string]Attr)
	for i := range AttrSlice {
		AttrMap[AttrSlice[i].Name] = AttrSlice[i]
	}
}

func AttrName(name string) Attr {
	return AttrMap[name]
}

func IsAttr(name string) bool {
	_, ok := AttrMap[name]
	return ok
}
