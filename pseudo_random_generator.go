package genstring

import (
	"math/bits"
	"math/rand"
	"strings"
)

type PseudoRandomGenerator struct {
	alphabet Alphabet
	source   rand.Rand
	seed     uint64
	seedMax  uint
	seedUsed uint
	builder  strings.Builder
}

func NewPseudoRandomGenerator(alphabet Alphabet) *PseudoRandomGenerator {
	source := *rand.New(rand.NewSource(rand.Int63()))
	return &PseudoRandomGenerator{
		alphabet: alphabet,
		source:   source,
		seed:     source.Uint64(),
		seedMax:  64 / intLog2[int, uint](len(alphabet)),
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
			p.seed = p.source.Uint64()
			p.seedUsed = 0
		}
	}

	generated = p.builder.String()
	p.builder.Reset()
	return
}
