package types

const (
	// ModuleName defines the IBC mockapp name
	ModuleName = "mockapp"

	// Version defines the current version the IBC mockapp
	// module supports
	Version = "mockapp-1"

	// PortID is the default port id that mockapp module binds to
	PortID = "mockapp"

	// StoreKey is the store key string for IBC mockapp
	StoreKey = ModuleName

	// RouterKey is the message route for IBC mockapp
	RouterKey = ModuleName

	// QuerierRoute is the querier route for IBC mockapp
	QuerierRoute = ModuleName
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = []byte{0x01}
)
