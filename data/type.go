package data

const (
	TypeName    = "name"
	TypeDate    = "date"
	TypeAddress = "address"
	TypePhone   = "phone"

	SubTypeStreet = "street"
	SubTypeCity   = "city"
)

var Supported = map[string]bool{
	TypeName:    true,
	TypeDate:    true,
	TypeAddress: true,
	TypePhone:   true,
}
