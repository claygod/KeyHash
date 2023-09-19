package KeyHash

// KeyHash
// Bench
// Copyright © 2016-2023 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"strconv"
	"testing"
)

const ADD_PATH_COUNT = 500001

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
	for i := ADD_PATH_COUNT; i > 0; i-- {
		x := strconv.Itoa(i)
		m[x] = i
	}
	b.StartTimer()
	for i := ADD_PATH_COUNT; i > 0; i-- {
		m["2"] = m["1"]
		m["1000"] = m["1001"]
		m["5000"] = m["5001"]
		m["9999"] = m["9991"]
		m["100000"] = m["100001"]
		m["200000"] = m["200001"]

	}
}

// Test claygod/KeyHash Long String key ns/op
func BenchmarkKeyHashLongStringKey(b *testing.B) {
	b.StopTimer()
	m := make(map[string]int)
	for i := ADD_PATH_COUNT; i > 0; i-- {
		x := strconv.Itoa(i) + " this very long strring key"
		m[x] = i
	}
	b.StartTimer()
	for i := ADD_PATH_COUNT; i > 0; i-- {
		m["2 this very long strring key"] = m["1 this very long strring key"]
		m["1000 this very long strring key"] = m["1001 this very long strring key"]
		m["5000 this very long strring key"] = m["5001 this very long strring key"]
		m["9999 this very long strring key"] = m["9991 this very long strring key"]
		m["100000 this very long strring key"] = m["100001 this very long strring key"]
		m["200000 this very long strring key"] = m["200001 this very long strring key"]

	}
}

// Test claygod/KeyHash uint64 key ns/op
func BenchmarkKeyHashUint64Key(b *testing.B) {
	b.StopTimer()
	m := make(map[uint64]int)
	for i := ADD_PATH_COUNT; i > 0; i-- {
		x := uint64(i)
		m[x] = i
	}
	b.StartTimer()
	for i := ADD_PATH_COUNT; i > 0; i-- {
		m[2] = m[1]
		m[1000] = m[1001]
		m[5000] = m[5001]
		m[9999] = m[9991]
		m[100000] = m[100001]
		m[200000] = m[200001]

	}
}
