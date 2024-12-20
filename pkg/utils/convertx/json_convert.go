package convertx

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

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
