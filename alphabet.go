package genstring

import (
	"math"
	"math/rand"
)

const (
	Base16Lower     Alphabet = "0123456789abcdef"
	Base16Upper     Alphabet = "0123456789ABCDEF"
	Base32Lower     Alphabet = "0123456789abcdefghijklmnopqrstuv"
	Base32Upper     Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUV"
	Crockford       Alphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	Base36Lower     Alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	Base36Upper     Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Base58          Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	Base62          Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Base64          Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	Base64URL       Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	PasswordCharAll Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-+*/=~!@#$%^&*"
	ASCII85         Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~"
)

type Alphabet string

func (a Alphabet) PermuteRunes() Alphabet {
	ar := []rune(a)
	shuffled := make([]rune, len(ar))
	for i, j := range rand.Perm(len(ar)) {
		shuffled[i] = ar[j]
	}
	return Alphabet(shuffled)
}

func (a Alphabet) PermuteBytes() Alphabet {
	shuffled := make([]byte, len(a))
	for i, j := range rand.Perm(len(a)) {
		shuffled[i] = a[j]
	}
	return Alphabet(shuffled)
}

func (a Alphabet) MinEncBits() int {
	return int(math.Floor(math.Log2(float64(len(a)))))
}

func (a Alphabet) MaxEncBits() int {
	return int(math.Ceil(math.Log2(float64(len(a)))))
}
