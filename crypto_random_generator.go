package genstring

import (
	"crypto/rand"
	"encoding/binary"
	"math/bits"
	"strings"
)

type CryptoRandomGenerator struct {
	alphabet   Alphabet
	randBufOff int
	randBuf    []byte
	seed       uint64
	seedMax    uint
	seedUsed   uint
	builder    strings.Builder
}

func NewCryptoRandomGenerator(alphabet Alphabet) CryptoRandomGenerator {
	buf := make([]byte, 1024)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}

	c := CryptoRandomGenerator{
		alphabet: alphabet,
		randBuf:  buf,
		seedMax:  64 / intLog2[int, uint](len(alphabet)),
	}

	c.seed = c.newSeed()

	return c
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
	if c.randBufOff+8 > len(c.randBuf) {
		c.randBufOff = 0
		if _, err := rand.Read(c.randBuf); err != nil {
			panic(err)
		}
	}

	seed = binary.NativeEndian.Uint64(c.randBuf[c.randBufOff : c.randBufOff+8])
	c.randBufOff += 8
	return
}

func (p *CryptoRandomGenerator) generate(length int) (generated string) {
	var idx0, idx1, idx2, idx3 uint64
	alphaLen := uint64(len(p.alphabet))
	p.builder.Grow(length)

	for i := 0; i < length; {
		switch i {
		case length - 1:
			p.seed, idx0 = bits.Div64(0, p.seed, alphaLen)
			p.builder.WriteByte(p.alphabet[idx0])
			p.seedUsed++
			i++

		case length - 2:
			p.seed, idx0 = bits.Div64(0, p.seed, alphaLen)
			p.seed, idx1 = bits.Div64(0, p.seed, alphaLen)
			p.builder.Write([]byte{p.alphabet[idx0], p.alphabet[idx1]})
			p.seedUsed += 2
			i += 2

		case length - 3:
			p.seed, idx0 = bits.Div64(0, p.seed, alphaLen)
			p.seed, idx1 = bits.Div64(0, p.seed, alphaLen)
			p.seed, idx2 = bits.Div64(0, p.seed, alphaLen)
			p.builder.Write([]byte{p.alphabet[idx0], p.alphabet[idx1], p.alphabet[idx2]})
			p.seedUsed += 3
			i += 3

		default:
			p.seed, idx0 = bits.Div64(0, p.seed, alphaLen)
			p.seed, idx1 = bits.Div64(0, p.seed, alphaLen)
			p.seed, idx2 = bits.Div64(0, p.seed, alphaLen)
			p.seed, idx3 = bits.Div64(0, p.seed, alphaLen)
			p.builder.Write([]byte{p.alphabet[idx0], p.alphabet[idx1], p.alphabet[idx2], p.alphabet[idx3]})
			p.seedUsed += 4
			i += 4
		}

		if p.seedMax-p.seedUsed < 4 {
			p.seed = p.newSeed()
			p.seedUsed = 0
		}
	}

	generated = p.builder.String()
	p.builder.Reset()
	return
}
