// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

package KeyHash

// KeyHash
// Test

import (
	"testing"
)

func TestMinZeroLenght(t *testing.T) {
	kh := New()
	if kh.Min("") != 0 {
		t.Errorf("it is unacceptable to handle the zero-length string")
	}
}

func TestMinOverLenght(t *testing.T) {
	kh := New()
	if kh.Min("123456789") != 0 {
		t.Errorf("the excess length of the string")
	}
}

func TestMidZeroLenght(t *testing.T) {
	kh := New()
	if kh.Mid("") != 0 {
		t.Errorf("it is unacceptable to handle the zero-length string")
	}
}

func TestMidOverLenght(t *testing.T) {
	kh := New()
	var str string
	for i := 0; i < 1025; i++ {
		str += "a"
	}
	if kh.Mid(str) != 0 {
		t.Errorf("the excess length of the string")
	}
}

func TestMaxZeroLenght(t *testing.T) {
	kh := New()
	if kh.Max("") != 0 {
		t.Errorf("it is unacceptable to handle the zero-length string")
	}
}
