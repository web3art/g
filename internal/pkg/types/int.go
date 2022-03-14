package types

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

var ErrOverflow error = errors.New("type overflow")

func ParseInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0, err
	}

	if i > 2147483647 {
		return 0, ErrOverflow
	}

	return int32(i), nil
}

func ParseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func ParseUInt64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func RandomWithRangeInt32(min, max int32) (int32, error) {
	if max < min {
		return 0, errors.New("max must be greater than min")
	}

	rand.Seed(time.Now().UnixNano())
	randval := rand.Float64()

	return int32(math.Floor(float64(min) + randval*float64(max-min) + 0.5)), nil
}

func RandomWithRange(min, max string) (int32, error) {
	minInt, err := ParseInt32(min)

	if err != nil {
		return 0, err
	}

	maxInt, err := ParseInt32(max)

	if err != nil {
		return 0, err
	}

	if maxInt < minInt {
		return 0, errors.New("max must be greater than min")
	}

	return RandomWithRangeInt32(minInt, maxInt)
}

func BigIntFromString(s string) (*big.Int, error) {
	i, b := new(big.Int).SetString(s, 10)

	if !b {
		return nil, fmt.Errorf("big int from string error %s", s)
	}

	return i, nil
}
