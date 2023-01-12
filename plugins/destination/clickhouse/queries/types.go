package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func chType(t schema.ValueType) string {
	// We allow nullable values in arrays.
	switch t {
	case schema.TypeBool:
		return `Nullable(Bool)`

	case schema.TypeFloat:
		return `Nullable(Float64)`

	case schema.TypeInt:
		return `Nullable(Int64)`
	case schema.TypeIntArray:
		return `Array(Nullable(Int64))`

	case schema.TypeString,
		schema.TypeByteArray,
		schema.TypeJSON, // ClickHouse can't handle values like [{"x":{"y":"z"}}] at the moment
		schema.TypeMacAddr,
		schema.TypeCIDR,
		schema.TypeInet:
		return `Nullable(String)`
	case schema.TypeStringArray,
		schema.TypeMacAddrArray,
		schema.TypeCIDRArray,
		schema.TypeInetArray:
		return `Array(Nullable(String))`

	case schema.TypeTimestamp:
		return `Nullable(DateTime64(9))`

	case schema.TypeUUID:
		return `Nullable(UUID)`
	case schema.TypeUUIDArray:
		return `Array(Nullable(UUID))`

	default:
		panic("unsupported type " + t.String())
	}
}
