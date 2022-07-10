# KeyHash

KeyHash Library generates a hash of the string as a Uint64.
It may be necessary to use hash tables Golang key is not a string,
but as the numbers to speed up a little bit to work with these tables.

[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/PiHex)](https://goreportcard.com/report/github.com/claygod/KeyHash)
[![API documentation](https://godoc.org/github.com/claygod/KeyHash?status.svg)](https://godoc.org/github.com/claygod/KeyHash)

# Usage

An example of using the KeyHash Library:

```go
package main

import (
	"fmt"
	"github.com/claygod/KeyHash"
)

func main() {
	kh := KeyHash.New()
	fmt.Print("\nTesting KeyHash library:")
	fmt.Print("\n Min: 12345678 -> ", kh.Min("12345678"))
	fmt.Print("\n Mid: long string -> ", kh.Mid("long string"))
	fmt.Print("\n Mid: very long string -> ", kh.Max("very long string"))
}
```

# Perfomance


The first three rows in the table show the speed of the three methods: Min, Mid and Max. The last three lines are shown to clearly see that the speed of the key type uint64 significantly higher than with the key string.

```
BenchmarkKeyHashMin-4           	1000000000	         0.06 ns/op
BenchmarkKeyHashMid-4           	1000000000	         0.06 ns/op
BenchmarkKeyHashMax-4           	2000000000	         0.08 ns/op
BenchmarkKeyHashShortStringKey-4	2000000000	         0.12 ns/op
BenchmarkKeyHashLongStringKey-4 	1000000000	         0.29 ns/op
BenchmarkKeyHashUint64Key-4     	2000000000	         0.09 ns/op
```

# API

Methods:
-  *New* - create a new KeyHash
-  *Min* - receiving a hash of a string up to 8 characters in length. If more than 8 characters, it returns zero. This method has no collisions! The conflicts are eliminated in full, 100 percent.
-  *Mid* - receiving a hash of a string up to ~1024 characters in length. This method can have some conflicts to produce results (probability of collision is small in this method).
-  *Max* - receiving a hash of a string of ANY length. The probability of collisions in this method is equal to twice the sum of the known MD5 algorithm. This method is slightly slower than the previous ones.


# Collisions

An important feature of all the programs that generate the hash, it is the presence and the possibility (probability) of collisions. Conflict arises from the fact that different lines can be generated by the same hashes. If you need to completely eliminate the conflict, in this library, you will approach only method Min. This method has only one drawback, it allows you to use the string is not longer than 8 characters (each character must be 1 byte, note in UTF-8 is often not the case). In other cases, I recommend using the method of Max and Mid method is used only when speed is of paramount importance to you.

# Notice

Attention! Use this library is not difficult. But if your application speed does not play a major and decisive role, use in daily practice string keys, keep it simple the program, remember to KISS.

### Copyright

Copyright © 2022 Eduard Sesigin. All rights reserved. Contacts: claygod@yandex.ru
