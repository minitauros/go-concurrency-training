package test

import (
	"encoding/json"
	"errors"
	"math/rand"
	"sync/atomic"
)

const defaultStringLength = 8
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// String returns a random string of a predefined length.
func String() string {
	return StringN(defaultStringLength)
}

// StringN returns a random string of the given length.
func StringN(length int) string {
	return string(BytesN(length))
}

// Bytes returns a random slice of bytes of a predefined length.
func Bytes() []byte {
	return BytesN(defaultStringLength)
}

// BytesN returns a random slice of bytes of the given length.
func BytesN(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

var uniqueInt64 atomic.Int64

// Int returns an int that has not been returned before, ensuring that always a unique number is returned.
func Int() int {
	// Acknowledged that if the int64 would reach its full capacity, it might not fit in an int, but it is unlikely
	// we'll increment it so many times that it actually reaches that limit.
	return int(uniqueInt64.Add(1))
}

// Uint returns a uint that has not been returned before, ensuring that always a unique number is returned.
func Uint() uint {
	// Acknowledged that if the int64 would reach its full capacity, it might not fit in an int, but it is unlikely
	// we'll increment it so many times that it actually reaches that limit.
	return uint(uniqueInt64.Add(1))
}

// Int64 returns an int64 that has not been returned before, ensuring that always a unique number is returned.
func Int64() int64 {
	return uniqueInt64.Add(1)
}

// Uint64 returns a uint64 that has not been returned before, ensuring that always a unique number is returned.
func Uint64() uint64 {
	return uint64(uniqueInt64.Add(1))
}

// Float32 returns a float32 that has not been returned before, ensuring that always a unique number is returned.
func Float32() float32 {
	return float32(uniqueInt64.Add(1))
}

// Float64 returns a float64 that has not been returned before, ensuring that always a unique number is returned.
func Float64() float64 {
	return float64(uniqueInt64.Add(1))
}

// Error returns a new error with a random string as message.
func Error() error {
	return errors.New(String())
}

// ToJson converts the given value to JSON and panics if that fails.
func ToJson(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

// ToJsonString converts the given value to JSON and panics if that fails. Returns a string instead of a slice of bytes.
func ToJsonString(v interface{}) string {
	return string(ToJson(v))
}
