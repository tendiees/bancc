package crypt_test

import (
	"testing"

	"github.com/tendiees/bancc/internal/crypt"
)

func BenchmarkHash(b *testing.B) {
	pass := []byte("UHpJxpp88zG/k3D3EvqqiNWdkR6KD4RFfFfX1VFQP6G60lA27rErBp2hGBCcXcU5")
	for i := 0; i < b.N; i++ {
		sh := crypt.SaltAndHash(pass)
		if !crypt.Verify(pass, sh) {
			b.Fatal("verify function is wrong")
		}
	}
}
