package seqx

import (
	"testing"
	"time"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompareID(t *testing.T) {
	id1 := NewID()
	id2 := NewID()

	var (
		result int
		err    error
	)

	// equal
	result, err = CompareID(id1, id1)
	require.NoError(t, err)
	assert.Equal(t, 0, result)
	result, err = CompareID(id2, id2)
	require.NoError(t, err)
	assert.Equal(t, 0, result)

	// asc
	result, err = CompareID(id1, id2)
	require.NoError(t, err)
	assert.Equal(t, -1, result)

	// desc
	result, err = CompareID(id2, id1)
	require.NoError(t, err)
	assert.Equal(t, 1, result)
}

func TestCheckIDAscending(t *testing.T) {
	preID := NewID()

	for i := 0; i < 100; i++ {
		id := NewID()
		asc, err := CheckIDAscending(preID, id)
		require.NoError(t, err)
		require.True(t, asc)

		preID = id
	}
}

func TestTimeFromID(t *testing.T) {
	ti, err := TimeFromID(NewID())
	assert.NoError(t, err)

	// delta no more than 1s
	assert.Equal(t, 0, int(time.Until(ti).Seconds()))

	_, err = TimeFromID("hello")
	assert.Error(t, err)
}

func TestNewWithTime(t *testing.T) {
	tim, err := time.Parse("2006-01-02 15:04:05", "2021-08-03 02:46:51")
	require.NoError(t, err)
	r, err := TimeFromID("c44as6o6n88m1pv5me7g")
	require.NoError(t, err)
	require.EqualValues(t, tim.UnixNano(), r.UnixNano())

	r, err = TimeFromID("c44as6s0pmuut3fqiq60")
	require.NoError(t, err)
	require.EqualValues(t, tim.UnixNano(), r.UnixNano())

	xidStr := xid.NewWithTime(tim)
	require.EqualValues(t, tim.UnixNano(), xidStr.Time().UnixNano())
}

func TestTimestampFromID(t *testing.T) {
	nowTsSeconds := uint64(time.Now().Unix())
	ts, err := TimestampFromID(NewID())
	assert.NoError(t, err)

	// delta no more than 1s
	assert.Equal(t, nowTsSeconds, ts/1000)

	_, err = TimeFromID("hello")
	assert.Error(t, err)
}

func TestPrevID(t *testing.T) {
	for i := 0; i < 100; i++ {
		id := NewID()
		newId, err := PrevID(id)
		require.NoError(t, err)

		asc, err := CheckIDAscending(newId, id)
		require.NoError(t, err)
		require.True(t, asc)

		del, err := CompareID(newId, id)
		require.NoError(t, err)
		require.Equal(t, -1, del)
	}
}
