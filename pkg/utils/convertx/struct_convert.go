package convertx

import (
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/datatypes"
)

// StructPtr2Pb converts a pointer to a go struct to a protobuf struct.
func StructPtr2Pb[T any](in *T) *structpb.Struct {
	if in == nil {
		return nil
	}
	var out structpb.Struct
	err := JsonConvert(in, &out)
	if err != nil {
		return nil
	}
	return &out
}

// StructPb2Ptr converts a protobuf struct to a pointer to a go struct.
func StructPb2Ptr[T any](in *structpb.Struct) *T {
	if in == nil {
		return nil
	}
	var out T
	err := JsonConvert(in, &out)
	if err != nil {
		return nil
	}
	return &out
}

// StructSql2Pb converts a gorm JSON to a protobuf struct.
// It treats the JSON as a object, though it may be an array or a scalar value.
func StructSql2Pb(in datatypes.JSON) *structpb.Struct {
	if len(in) == 0 || string(in) == "null" {
		return nil
	}
	var out structpb.Struct
	err := JsonConvert(in, &out)
	if err != nil {
		return nil
	}
	return &out
}

// StructPb2Sql converts a protobuf struct to a gorm JSON.
func StructPb2Sql(in *structpb.Struct) datatypes.JSON {
	if in == nil {
		return datatypes.JSON("")
	}
	var out datatypes.JSON
	err := JsonConvert(in, &out)
	if err != nil {
		return nil
	}
	return out
}

// StructSql2Ptr converts a gorm JSON to a pointer to a go struct.
// It treats the JSON as a object, though it may be an array or a scalar value.
func StructSql2Ptr[T any](in datatypes.JSON) *T {
	if len(in) == 0 || string(in) == "null" {
		return nil
	}
	var out T
	err := JsonConvert(in, &out)
	if err != nil {
		return nil
	}
	return &out
}

// StructPtr2Sql converts a pointer to a go struct to a gorm JSON.
func StructPtr2Sql[T any](in *T) datatypes.JSON {
	if in == nil {
		return datatypes.JSON("")
	}
	var out datatypes.JSON
	err := JsonConvert(in, &out)
	if err != nil {
		return nil
	}
	return out
}
