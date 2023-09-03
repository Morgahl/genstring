package genstring

import (
	"crypto/rand"
	"encoding/binary"
	"math/bits"
	"strings"
)

type CryptoRandomGenerator struct {
	alphabet   Alphabet
	builder    strings.Builder
	randBufOff int
	randBuf    []byte
}

func NewCryptoRandomGenerator(alphabet Alphabet) CryptoRandomGenerator {
	buf := make([]byte, 1024)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}

	return CryptoRandomGenerator{
		alphabet: alphabet,
		randBuf:  buf,
	}
}

func (c CryptoRandomGenerator) Password(length int) Password {
	return Password(c.generate(length))
}

func (c CryptoRandomGenerator) Passwords(length, count int) []Password {
	passwords := make([]Password, count)
	for i := 0; i < count; i++ {
		passwords[i] = c.Password(length)
	}
	return passwords
}

func (c *CryptoRandomGenerator) newSeed() (seed uint64) {
	if c.randBufOff == len(c.randBuf) {
		c.randBufOff = 0
		if _, err := rand.Read(c.randBuf); err != nil {
			panic(err)
		}
	}

	seed = binary.NativeEndian.Uint64(c.randBuf[c.randBufOff : c.randBufOff+8])
	c.randBufOff += 8
	return
}

func (c *CryptoRandomGenerator) generate(length int) (generated string) {
	c.builder.Grow(length)

	alphaLen := uint64(len(c.alphabet))
	seed := uint64(0)
	idx0 := uint64(0)
	idx1 := uint64(0)
	idx2 := uint64(0)
	idx3 := uint64(0)
	for i := 0; i < length; {
		switch i {
		case length - 1:
			if seed < alphaLen {
				seed = c.newSeed()
			}
			seed, idx0 = bits.Div64(0, seed, alphaLen)
			c.builder.WriteByte(c.alphabet[idx0])
			i++

		case length - 2:
			if seed < intPow(alphaLen, 2) {
				seed = c.newSeed()
			}
			seed, idx0 = bits.Div64(0, seed, alphaLen)
			seed, idx1 = bits.Div64(0, seed, alphaLen)
			c.builder.Write([]byte{c.alphabet[idx0], c.alphabet[idx1]})
			i += 2

		case length - 3:
			if seed < intPow(alphaLen, 3) {
				seed = c.newSeed()
			}
			seed, idx0 = bits.Div64(0, seed, alphaLen)
			seed, idx1 = bits.Div64(0, seed, alphaLen)
			seed, idx2 = bits.Div64(0, seed, alphaLen)
			c.builder.Write([]byte{c.alphabet[idx0], c.alphabet[idx1], c.alphabet[idx2]})
			i += 3

		default:
			if seed < intPow(alphaLen, 4) {
				seed = c.newSeed()
			}
			seed, idx0 = bits.Div64(0, seed, alphaLen)
			seed, idx1 = bits.Div64(0, seed, alphaLen)
			seed, idx2 = bits.Div64(0, seed, alphaLen)
			seed, idx3 = bits.Div64(0, seed, alphaLen)
			c.builder.Write([]byte{c.alphabet[idx0], c.alphabet[idx1], c.alphabet[idx2], c.alphabet[idx3]})
			i += 4
		}
	}

	generated = c.builder.String()
	c.builder.Reset()
	return
}
