package catalog

var Attributes []string = []string{
	"id",
	"author",
	"title",
	"full_vendor_name",
	"availability",
}

var AttributeMap map[string]struct{}

func Init() {
	AttributeMap = make(map[string]struct{})
	for i := range Attributes {
		AttributeMap[Attributes[i]] = struct{}{}
	}
}

func IsAttr(attr string) bool {
	_, ok := AttributeMap[attr]
	return ok
}
