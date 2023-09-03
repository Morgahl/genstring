package genstring_test

import (
	"testing"

	"github.com/morgahl/genstring"
)

func BenchmarkCryptoRandomGenerator_Password_1(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 1)
}

func BenchmarkCryptoRandomGenerator_Password_2(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 2)
}

func BenchmarkCryptoRandomGenerator_Password_4(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 4)
}

func BenchmarkCryptoRandomGenerator_Password_8(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 8)
}

func BenchmarkCryptoRandomGenerator_Password_16(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 16)
}

func BenchmarkCryptoRandomGenerator_Password_32(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 32)
}

func BenchmarkCryptoRandomGenerator_Password_64(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 64)
}

func BenchmarkCryptoRandomGenerator_Password_128(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 128)
}

func BenchmarkCryptoRandomGenerator_Password_256(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 256)
}

func BenchmarkCryptoRandomGenerator_Password_512(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 512)
}

func BenchmarkCryptoRandomGenerator_Password_1024(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 1024)
}

func BenchmarkCryptoRandomGenerator_Password_2048(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 2048)
}

func BenchmarkCryptoRandomGenerator_Password_4096(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 4096)
}

func BenchmarkCryptoRandomGenerator_Password_8192(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 8192)
}

func BenchmarkCryptoRandomGenerator_Password_16384(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 16384)
}

func BenchmarkCryptoRandomGenerator_Password_32768(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 32768)
}

func BenchmarkCryptoRandomGenerator_Password_65536(b *testing.B) {
	benchCryptoRandomGenerator_Password(b, 65536)
}

func benchCryptoRandomGenerator_Password(b *testing.B, length int) {
	g := genstring.NewCryptoRandomGenerator(genstring.PasswordCharAll)
	for i := 0; i < b.N; i++ {
		g.Password(length)
	}
}
