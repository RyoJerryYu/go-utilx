package timerx

import "time"

type Timer interface {
	Next() time.Duration
	Reset()
	Clone() Timer
}
