package crypt

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

const (
	HalfSize = 64
	Size     = 2 * HalfSize

	Time = 16

	MemSize     = 128 * 1024
	ThreadCount = 4
)

type SaltedHash [Size]byte
type Hash [HalfSize]byte
type Salt [HalfSize]byte

func (sh SaltedHash) Salt() (s Salt) {
	copy(s[:], sh[:HalfSize])
	return s
}

func (sh SaltedHash) Hash() (h Hash) {
	copy(h[:], sh[HalfSize:])
	return h
}

func SaltAndHash(pass []byte) (sh SaltedHash) {
	_, _ = rand.Read(sh[:HalfSize])
	hash := argon2.IDKey(pass, sh[:HalfSize], Time, MemSize, ThreadCount, HalfSize)
	copy(sh[HalfSize:], hash)
	return sh
}

func Verify(pass []byte, sh SaltedHash) bool {
	s, h := sh.Salt(), sh.Hash()
	hash := argon2.IDKey(pass, s[:], Time, MemSize, ThreadCount, HalfSize)
	return bytes.Equal(h[:], hash)
}
