package seqx

import (
	"crypto/rand"
	"encoding/hex"
	"math"
	mrand "math/rand"
	"time"

	"github.com/rs/xid"
)

// 返回一个随机字符串，前12位为 xid ，后为随机数，建议 length >= 12
func GenerateRandomStr(length int) string {
	now := time.Now()
	id := xid.NewWithTime(now)

	// [0..11]: xid
	// [12..23]: rand bytes
	buffLen := length
	if buffLen < 12 {
		buffLen = 12
	}
	randomStr := make([]byte, buffLen)
	copy(randomStr[:], id.Bytes())

	n, _ := rand.Read(randomStr[12:])
	if length >= 12 && n != length-12 {
		rand := mrand.New(mrand.NewSource(now.UnixNano() + int64(n)))
		for i := 12 + n; i < length; i++ {
			randomStr[i] = byte(rand.Int63() % math.MaxUint8)
		}
	} else if length < 12 {
		// xid 前几位为 timestamp 与 machine identifier ，区分度较高
		randomStr = randomStr[:length]
	}

	return hex.EncodeToString(randomStr)
}
