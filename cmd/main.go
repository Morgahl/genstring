package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/morgahl/genstring"
)

const (
	PASSWORD_LENGTH = 64
)

func main() {
	var password_length = PASSWORD_LENGTH
	if len(os.Args) > 1 {
		length, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		password_length = length
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', tabwriter.Debug)

	pseudoPass := genstring.NewPseudoRandomGenerator(genstring.PasswordCharAll.PermuteRunes()).Password(password_length)
	fmt.Fprintln(w, "fmt.Println\t", pseudoPass)
	fmt.Fprintf(w, "%%s\t %s\n", pseudoPass)
	fmt.Fprintf(w, "%%v\t %v\n", pseudoPass)
	fmt.Fprintf(w, "%%+v\t %+v\n", pseudoPass)
	fmt.Fprintf(w, "%%#v\t %#v\n", pseudoPass)
	fmt.Fprintf(w, "%%q\t %q\n", pseudoPass)
	fmt.Fprintf(w, "%%x\t %x\n", pseudoPass)
	fmt.Fprintf(w, "%%X\t %X\n", pseudoPass)
	fmt.Fprintf(w, "string(pass)\t %s\n", string(pseudoPass))
	w.Flush()
	println("println(pass)\t|", pseudoPass)

	pseudoPass2 := genstring.NewPseudoRandomGenerator(genstring.PasswordCharAll.PermuteRunes()).Password(password_length)
	fmt.Fprintln(w, "fmt.Println\t", pseudoPass2)
	fmt.Fprintf(w, "%%s\t %s\n", pseudoPass2)
	fmt.Fprintf(w, "%%v\t %v\n", pseudoPass2)
	fmt.Fprintf(w, "%%+v\t %+v\n", pseudoPass2)
	fmt.Fprintf(w, "%%#v\t %#v\n", pseudoPass2)
	fmt.Fprintf(w, "%%q\t %q\n", pseudoPass2)
	fmt.Fprintf(w, "%%x\t %x\n", pseudoPass2)
	fmt.Fprintf(w, "%%X\t %X\n", pseudoPass2)
	fmt.Fprintf(w, "string(pass)\t %s\n", string(pseudoPass2))
	w.Flush()
	println("println(pass)\t|", pseudoPass2)

	cryptoPass := genstring.NewCryptoRandomGenerator(genstring.PasswordCharAll.PermuteRunes()).Password(password_length)
	fmt.Fprintln(w, "fmt.Println\t", cryptoPass)
	fmt.Fprintf(w, "%%s\t %s\n", cryptoPass)
	fmt.Fprintf(w, "%%v\t %v\n", cryptoPass)
	fmt.Fprintf(w, "%%+v\t %+v\n", cryptoPass)
	fmt.Fprintf(w, "%%#v\t %#v\n", cryptoPass)
	fmt.Fprintf(w, "%%q\t %q\n", cryptoPass)
	fmt.Fprintf(w, "%%x\t %x\n", cryptoPass)
	fmt.Fprintf(w, "%%X\t %X\n", cryptoPass)
	fmt.Fprintf(w, "string(pass)\t %s\n", string(cryptoPass))
	w.Flush()
	println("println(pass)\t|", cryptoPass)
}
