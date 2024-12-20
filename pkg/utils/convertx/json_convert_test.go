package convertx

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"
)

type TypeRaw struct {
	A int            `json:"a"`
	B string         `json:"b"`
	C float64        `json:"c"`
	D bool           `json:"d"`
	E []int          `json:"e"`
	F map[string]int `json:"f"`
}

func (t TypeRaw) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"a": t.A,
		"b": t.B,
		"c": t.C,
		"d": t.D,
		"e": t.E,
		"f": t.F,
	}

	return json.Marshal(m)
}

type TypeAlias struct {
	Int   int            `json:"a"`
	Str   string         `json:"b"`
	Float float64        `json:"c"`
	Bool  bool           `json:"d"`
	Ints  []int          `json:"e"`
	Map   map[string]int `json:"f"`
}

func (t *TypeAlias) UnmarshalJSON(b []byte) error {
	var raw TypeRaw
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	t.Int = raw.A
	t.Str = raw.B
	t.Float = raw.C
	t.Bool = raw.D
	t.Ints = raw.E
	t.Map = raw.F

	return nil
}

func TestJsonConvert(t *testing.T) {
	t.Run("TypeRaw to TypeAlias", func(t *testing.T) {
		raw := TypeRaw{
			A: 1,
			B: "2",
			C: 3.14,
			D: true,
			E: []int{1, 2, 3},
			F: map[string]int{"a": 1, "b": 2},
		}

		var alias TypeAlias
		err := JsonConvert(raw, &alias)
		require.NoError(t, err)

		require.Equal(t, raw.A, alias.Int)
		require.Equal(t, raw.B, alias.Str)
		require.Equal(t, raw.C, alias.Float)
		require.Equal(t, raw.D, alias.Bool)
		require.Equal(t, raw.E, alias.Ints)
		require.Equal(t, raw.F, alias.Map)
	})
}

func TestJsonConvert_Protobuf(t *testing.T) {
	t.Run("TypeRaw to TypeAlias", func(t *testing.T) {
		raw := TypeRaw{
			A: 1,
			B: "2",
			C: 3.14,
			D: true,
			E: []int{1, 2, 3},
			F: map[string]int{"a": 1, "b": 2},
		}

		var out structpb.Struct
		err := JsonConvert(raw, &out)
		require.NoError(t, err)

		require.Equal(t, raw.A, int(out.Fields["a"].GetNumberValue()))
		require.Equal(t, raw.B, out.Fields["b"].GetStringValue())
		require.Equal(t, raw.C, out.Fields["c"].GetNumberValue())
		require.Equal(t, raw.D, out.Fields["d"].GetBoolValue())
		require.Len(t, out.Fields["e"].GetListValue().AsSlice(), len(raw.E))
		require.Len(t, out.Fields["f"].GetStructValue().Fields, len(raw.F))
	})
}
