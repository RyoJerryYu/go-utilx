package convertx

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// JsonConvert converts between two types using JSON marshaling and unmarshaling.
// It handles special cases for json.Marshaler/Unmarshaler and protobuf messages.
//
// The function will:
// 1. Marshal the input based on its type:
//   - json.Marshaler: uses MarshalJSON()
//   - proto.Message: uses protojson.Marshal()
//   - others: uses json.Marshal()
//
// 2. Unmarshal the JSON bytes into the output based on its type:
//   - json.Unmarshaler: uses UnmarshalJSON()
//   - proto.Message: uses protojson.Unmarshal()
//   - others: uses json.Unmarshal()
//
// Example usage:
//
//	type User struct {
//	    Name string `json:"name"`
//	    Age  int    `json:"age"`
//	}
//
//	in := User{Name: "John", Age: 30}
//	var out map[string]interface{}
//	err := JsonConvert(in, &out)
func JsonConvert(in any, out any) error {
	var (
		b   []byte
		err error
	)
	switch in := in.(type) {
	case json.Marshaler:
		b, err = in.MarshalJSON()
	case proto.Message:
		b, err = protojson.Marshal(in)
	default:
		b, err = json.Marshal(in)
	}
	if err != nil {
		return err
	}

	switch out := out.(type) {
	case json.Unmarshaler:
		return out.UnmarshalJSON(b)
	case proto.Message:
		return protojson.Unmarshal(b, out)
	default:
		return json.Unmarshal(b, out)
	}
}
