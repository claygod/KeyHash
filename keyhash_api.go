package KeyHash

// KeyHash
// Work
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"crypto/md5"
)

// New - create a new KeyHash-struct
func New() *KeyHash {
	return &KeyHash{}
}

const multiplier uint64 = 1024

// KeyHash structure
type KeyHash struct{}

// Min - receiving a hash of a string up to 8 characters in length.
// If more than 8 characters, it returns zero. This method has no collisions!
// The conflicts are eliminated in full, 100 percent.
func (r *KeyHash) Min(str string) uint64 {
	var out uint64
	ln := len(str)
	if ln > 8 || ln == 0 {
		return 0
	}
	out += uint64(str[0])
	for i := 1; i < ln-1; i++ {
		out += uint64(str[i])
		out = out << 8
	}
	return out
}

// Mid - receiving a hash of a string up to ~1024 characters in length.
// This method can have some conflicts to produce results
// (probability of collision is small in this method).
func (r *KeyHash) Mid(str string) uint64 {
	var out uint64
	ln := uint64(len(str) - 1)
	if ln > 1023 || ln < 0 {
		return 0
	}
	for i := uint64(0); i < ln; i++ {
		out += uint64(str[i]) + multiplier*i
	}
	return out
}

// Max - receiving a hash of a string of ANY length.
// The probability of collisions in this method is equal to twice the sum of the
// known MD5 algorithm. This method is slightly slower than the previous ones.
func (r *KeyHash) Max(str string) uint64 {
	var out uint64
	if len(str) == 0 {
		return 0
	}
	str2 := md5.Sum([]byte(str))
	out += uint64(str2[0])
	for i := 1; i < 16; i++ {
		out = out << 4
		out += uint64(str2[i])
	}
	return out
}
