package types

import (
	"fmt"
	"strconv"

	"github.com/bookstairs/bookworm/util"
)

type NeedleId uint64

const (
	NeedleIdSize  = 8
	NeedleIdEmpty = 0
)

func NeedleIdToBytes(bytes []byte, needleId NeedleId) {
	util.Uint64toBytes(bytes, uint64(needleId))
}

// NeedleIdToUint64 used to send max needle id to master
func NeedleIdToUint64(needleId NeedleId) uint64 {
	return uint64(needleId)
}

func Uint64ToNeedleId(needleId uint64) NeedleId {
	return NeedleId(needleId)
}

func BytesToNeedleId(bytes []byte) NeedleId {
	return NeedleId(util.BytesToUint64(bytes))
}

func (k NeedleId) String() string {
	return strconv.FormatUint(uint64(k), 16)
}

func ParseNeedleId(idString string) (NeedleId, error) {
	key, err := strconv.ParseUint(idString, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("needle id %s format error: %v", idString, err)
	}
	return NeedleId(key), nil
}
