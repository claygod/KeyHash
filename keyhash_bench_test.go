// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

package KeyHash

// KeyHash
// Bench

import (
	"strconv"
	"testing"
)

const ADD_PATH_COUNT = 1000000

// Test claygod/KeyHash Min ns/op
func BenchmarkKeyHashMin(b *testing.B) {
	kh := New()
	for i := ADD_PATH_COUNT; i > 0; i-- {
		kh.Min("a" + strconv.Itoa(i))
	}
}

// Test claygod/KeyHash Mid ns/op
func BenchmarkKeyHashMid(b *testing.B) {
	kh := New()
	for i := ADD_PATH_COUNT; i > 0; i-- {
		kh.Mid("a" + strconv.Itoa(i))
	}
}

// Test claygod/KeyHash Max ns/op
func BenchmarkKeyHashMax(b *testing.B) {
	kh := New()
	for i := ADD_PATH_COUNT; i > 0; i-- {
		kh.Max("a" + strconv.Itoa(i))
	}
}

// Test claygod/KeyHash Short String key ns/op
func BenchmarkKeyHashShortStringKey(b *testing.B) {
	b.StopTimer()
	m := make(map[string]int)
	var a int
	var z int
	for i := ADD_PATH_COUNT; i > 0; i-- {
		x := strconv.Itoa(i)
		m[x] = i
	}
	b.StartTimer()
	for i := ADD_PATH_COUNT; i > 0; i-- {
		z = 0
		a = m["1"]
		a += m["1000"]
		a += m["5000"]
		a += m["9999"]
		a += m["100000"]
		a += m["500000"]
		a += m["999999"]
		z += a
	}
}

// Test claygod/KeyHash Long String key ns/op
func BenchmarkKeyHashLongStringKey(b *testing.B) {
	b.StopTimer()
	m := make(map[string]int)
	var a int
	var z int
	for i := ADD_PATH_COUNT; i > 0; i-- {
		x := strconv.Itoa(i) + " this very long strring key"
		m[x] = i
	}
	b.StartTimer()
	for i := ADD_PATH_COUNT; i > 0; i-- {
		z = 0
		a = m["1 this very long strring key"]
		a += m["1000 this very long strring key"]
		a += m["5000 this very long strring key"]
		a += m["9999 this very long strring key"]
		a += m["100000 this very long strring key"]
		a += m["500000 this very long strring key"]
		a += m["999999 this very long strring key"]
		z += a
	}
}

// Test claygod/KeyHash uint64 key ns/op
func BenchmarkKeyHashUint64Key(b *testing.B) {
	b.StopTimer()
	m := make(map[uint64]int)
	var a int
	var z int
	for i := ADD_PATH_COUNT; i > 0; i-- {
		x := uint64(i)
		m[x] = i
	}
	b.StartTimer()
	for i := ADD_PATH_COUNT; i > 0; i-- {
		z = 0
		a = m[1]
		a += m[1000]
		a += m[5000]
		a += m[9999]
		a += m[100000]
		a += m[500000]
		a += m[999999]
		z += a
	}
}
