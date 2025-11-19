package seqx

import (
	"time"

	"github.com/rs/xid"
)

func GetTime(id string) (time.Time, error) {
	x, err := xid.FromString(id)
	if err != nil {
		return time.Time{}, err
	}
	return x.Time(), nil
}

func GetPID(id string) (uint16, error) {
	x, err := xid.FromString(id)
	if err != nil {
		return 0, err
	}
	return x.Pid(), nil
}

func GetMachine(id string) ([]byte, error) {
	x, err := xid.FromString(id)
	if err != nil {
		return nil, err
	}
	return x.Machine(), nil
}
