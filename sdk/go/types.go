package feast

import "github.com/gojek/feast/sdk/go/protos/feast/types"

type Row map[string]*types.Value

// StrVal is a int64 type feast value
func StrVal(val string) *types.Value {
	return &types.Value{Val: &types.Value_StringVal{StringVal: val}}
}

// Int32Val is a int64 type feast value
func Int32Val(val int32) *types.Value {
	return &types.Value{Val: &types.Value_Int32Val{Int32Val: val}}
}

// Int64Val is a int64 type feast value
func Int64Val(val int64) *types.Value {
	return &types.Value{Val: &types.Value_Int64Val{Int64Val: val}}
}

// DoubleVal is a float64 type feast value
func DoubleVal(val float64) *types.Value {
	return &types.Value{Val: &types.Value_DoubleVal{DoubleVal: val}}
}

// BoolVal is a bool type feast value
func BoolVal(val bool) *types.Value {
	return &types.Value{Val: &types.Value_BoolVal{BoolVal: val}}
}

// BytesVal is a bytes type feast value
func BytesVal(val []byte) *types.Value {
	return &types.Value{Val: &types.Value_BytesVal{BytesVal: val}}
}