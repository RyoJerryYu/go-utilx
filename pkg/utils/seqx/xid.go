package seqx

import (
	"time"

	"github.com/rs/xid"
)

func NewID() string {
	return xid.New().String()
}

// Compare returns an integer comparing two IDs. It behaves just like `bytes.Compare`.
// The result will be 0 if two IDs are identical, -1 if current id is less than the other one,
// and 1 if current id is greater than the other.
func CompareID(id1, id2 string) (int, error) {
	xid1, err := xid.FromString(id1)
	if err != nil {
		return 0, err
	}

	xid2, err := xid.FromString(id2)
	if err != nil {
		return 0, err
	}

	return xid1.Compare(xid2), nil
}

func CheckIDAscending(id1, id2 string) (bool, error) {
	i, err := CompareID(id1, id2)
	if err != nil {
		return false, err
	}

	return i == -1, nil
}

func TimeFromID(id string) (time.Time, error) {
	x, err := xid.FromString(id)
	if err != nil {
		return time.Time{}, err
	}
	return x.Time(), nil
}

func TimestampFromID(id string) (uint64, error) {
	t, err := TimeFromID(id)
	if err != nil {
		return 0, err
	}

	return uint64(t.UnixNano() / int64(time.Millisecond)), nil
}

func PrevID(id string) (string, error) {
	x, err := xid.FromString(id)
	if err != nil {
		return "", err
	}

	return xid.NewWithTime(x.Time().Add(-time.Millisecond)).String(), nil
}

func NextID(id string) (string, error) {
	x, err := xid.FromString(id)
	if err != nil {
		return "", err
	}

	return xid.NewWithTime(x.Time().Add(time.Millisecond)).String(), nil
}
