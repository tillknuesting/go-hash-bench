package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	x "github.com/cespare/xxhash/v2"
	m "github.com/spaolacci/murmur3"
)

func Sum64StringMurmur3(value string) uint64 {
	return m.Sum64([]byte(value))
}

func Sum64StringXxhash(value string) uint64 {
	return x.Sum64String(value)
}

func Sum64StringSHA1(value string) uint64 {
	h := sha1.New()
	h.Write([]byte(value))
	return uint64(h.Sum(nil)[0])
}

func Sum64StringSHA256(value string) uint64 {
	h := sha256.New()
	h.Write([]byte(value))
	return uint64(h.Sum(nil)[0])
}

func Sum64StringMD5(value string) uint64 {
	h := md5.New()
	h.Write([]byte(value))
	return uint64(h.Sum(nil)[0])
}
