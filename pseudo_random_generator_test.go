package genstring_test

import (
	"testing"

	"github.com/morgahl/genstring"
)

func BenchmarkPseudoRandomGeneratorPassword_1(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 1)
}

func BenchmarkPseudoRandomGeneratorPassword_2(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 2)
}

func BenchmarkPseudoRandomGeneratorPassword_4(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 4)
}

func BenchmarkPseudoRandomGeneratorPassword_8(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 8)
}

func BenchmarkPseudoRandomGeneratorPassword_16(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 16)
}

func BenchmarkPseudoRandomGeneratorPassword_32(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 32)
}

func BenchmarkPseudoRandomGeneratorPassword_64(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 64)
}

func BenchmarkPseudoRandomGeneratorPassword_128(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 128)
}

func BenchmarkPseudoRandomGeneratorPassword_256(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 256)
}

func BenchmarkPseudoRandomGeneratorPassword_512(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 512)
}

func BenchmarkPseudoRandomGeneratorPassword_1024(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 1024)
}

func BenchmarkPseudoRandomGeneratorPassword_2048(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 2048)
}

func BenchmarkPseudoRandomGeneratorPassword_4096(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 4096)
}

func BenchmarkPseudoRandomGeneratorPassword_8192(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 8192)
}

func BenchmarkPseudoRandomGeneratorPassword_16384(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 16384)
}

func BenchmarkPseudoRandomGeneratorPassword_32768(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 32768)
}

func BenchmarkPseudoRandomGeneratorPassword_65536(b *testing.B) {
	benchPseudoRandomGeneratorPassword(b, 65536)
}

func benchPseudoRandomGeneratorPassword(b *testing.B, length int) {
	g := genstring.NewPseudoRandomGenerator(genstring.PasswordCharAll)
	for i := 0; i < b.N; i++ {
		g.Password(length)
	}
}
