package containerx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Int64IdEntity struct {
	Id int64
}

func (e Int64IdEntity) GetId() int64 {
	return e.Id
}

type StringIdEntity struct {
	Id string
}

func (e StringIdEntity) GetId() string {
	return e.Id
}

func (e StringIdEntity) GetName() string {
	return e.Id
}

func TestToId(t *testing.T) {
	entities := []Int64IdEntity{{Id: 1}, {Id: 2}, {Id: 3}}
	ids := ToIds(entities)
	assert.Equal(t, []int64{1, 2, 3}, ids)

	entities2 := []StringIdEntity{{Id: "1"}, {Id: "2"}, {Id: "3"}}
	ids2 := ToIds(entities2)
	assert.Equal(t, []string{"1", "2", "3"}, ids2)
}

func TestMapById(t *testing.T) {
	entities := []Int64IdEntity{{Id: 1}, {Id: 2}, {Id: 3}}
	m := MapByIds(entities)
	assert.Equal(t, map[int64]Int64IdEntity{1: {Id: 1}, 2: {Id: 2}, 3: {Id: 3}}, m)

	entities2 := []StringIdEntity{{Id: "1"}, {Id: "2"}, {Id: "3"}}
	m2 := MapByIds(entities2)
	assert.Equal(t, map[string]StringIdEntity{"1": {Id: "1"}, "2": {Id: "2"}, "3": {Id: "3"}}, m2)
}

func TestFilterByIds(t *testing.T) {
	entities := []Int64IdEntity{{Id: 1}, {Id: 2}, {Id: 3}}
	filtered := FilterByIds(entities, 1, 3)
	assert.Equal(t, []Int64IdEntity{{Id: 1}, {Id: 3}}, filtered)
}

func TestGroupByNames(t *testing.T) {
	entities := []StringIdEntity{{Id: "1"}, {Id: "2"}, {Id: "3"}}
	groups := GroupByNames(entities)
	assert.Equal(t, map[string][]StringIdEntity{"1": {{Id: "1"}}, "2": {{Id: "2"}}, "3": {{Id: "3"}}}, groups)
}

func TestChunkByIds(t *testing.T) {
	entities := []Int64IdEntity{{Id: 1}, {Id: 2}, {Id: 3}}
	chunks := ChunkByIds(entities)
	assert.Equal(t, [][]Int64IdEntity{{{Id: 1}}, {{Id: 2}}, {{Id: 3}}}, chunks)
}
