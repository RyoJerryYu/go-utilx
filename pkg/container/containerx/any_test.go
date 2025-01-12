package containerx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToAny(t *testing.T) {
	entities := []Int64IdEntity{{Id: 1}, {Id: 2}, {Id: 3}}
	ids := ToAny(entities)
	assert.Equal(t, []any{Int64IdEntity{Id: 1}, Int64IdEntity{Id: 2}, Int64IdEntity{Id: 3}}, ids)

	entities2 := []StringIdEntity{{Id: "1"}, {Id: "2"}, {Id: "3"}}
	ids2 := ToAny(entities2)
	assert.Equal(t, []any{StringIdEntity{Id: "1"}, StringIdEntity{Id: "2"}, StringIdEntity{Id: "3"}}, ids2)
}
