package genstring

import (
	"math/bits"
	"math/rand"
	"strings"
)

type PseudoRandomGenerator struct {
	alphabet Alphabet
	builder  strings.Builder
	source   rand.Rand
}

func NewPseudoRandomGenerator(alphabet Alphabet) *PseudoRandomGenerator {
	return &PseudoRandomGenerator{
		alphabet: alphabet,
		source:   *rand.New(rand.NewSource(rand.Int63())),
	}
}

func (p *PseudoRandomGenerator) Password(length int) Password {
	return Password(p.generate(length))
}

func (p *PseudoRandomGenerator) Passwords(length, count int) []Password {
	passwords := make([]Password, count)
	for i := 0; i < count; i++ {
		passwords[i] = p.Password(length)
	}
	return passwords
}

func (p *PseudoRandomGenerator) generate(length int) (generated string) {
	p.builder.Grow(length)

	alphaLen := uint64(len(p.alphabet))
	seed := uint64(0)
	idx0 := uint64(0)
	idx1 := uint64(0)
	idx2 := uint64(0)
	idx3 := uint64(0)
	for i := 0; i < length; {
		switch i {
		case length - 1:
			if seed < alphaLen {
				seed = p.source.Uint64()
			}
			seed, idx0 = bits.Div64(0, seed, alphaLen)
			p.builder.WriteByte(p.alphabet[idx0])
			i++

		case length - 2:
			if seed < intPow(alphaLen, 2) {
				seed = p.source.Uint64()
			}
			seed, idx0 = bits.Div64(0, seed, alphaLen)
			seed, idx1 = bits.Div64(0, seed, alphaLen)
			p.builder.Write([]byte{p.alphabet[idx0], p.alphabet[idx1]})
			i += 2

		case length - 3:
			if seed < intPow(alphaLen, 3) {
				seed = p.source.Uint64()
			}
			seed, idx0 = bits.Div64(0, seed, alphaLen)
			seed, idx1 = bits.Div64(0, seed, alphaLen)
			seed, idx2 = bits.Div64(0, seed, alphaLen)
			p.builder.Write([]byte{p.alphabet[idx0], p.alphabet[idx1], p.alphabet[idx2]})
			i += 3

		default:
			if seed < intPow(alphaLen, 4) {
				seed = p.source.Uint64()
			}
			seed, idx0 = bits.Div64(0, seed, alphaLen)
			seed, idx1 = bits.Div64(0, seed, alphaLen)
			seed, idx2 = bits.Div64(0, seed, alphaLen)
			seed, idx3 = bits.Div64(0, seed, alphaLen)
			p.builder.Write([]byte{p.alphabet[idx0], p.alphabet[idx1], p.alphabet[idx2], p.alphabet[idx3]})
			i += 4
		}
	}

	generated = p.builder.String()
	p.builder.Reset()
	return
}
