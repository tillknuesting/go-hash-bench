package main

import (
	x "github.com/cespare/xxhash/v2"
	m "github.com/spaolacci/murmur3"
)

func Sum64StringMurmur3(value string) uint64 {
	return m.Sum64([]byte(value))
}

func Sum64StringXxhash(value string) uint64 {
	return x.Sum64String(value)
}
