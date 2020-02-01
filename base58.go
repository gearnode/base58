package base58

import (
	"fmt"
)

type Alphabet struct {
	Decode [128]int8
	Encode [58]byte
}

func NewAlphabet(src string) (*Alphabet, error) {
	if len(src) != 58 {
		return nil, fmt.Errorf("invalid alphabet: base58 alphabets must be 58 bytes long")
	}

	var alphabet Alphabet
	copy(alphabet.Encode[:], src)

	for i := range alphabet.Decode {
		alphabet.Decode[i] = -1
	}

	for i, b := range alphabet.Encode {
		alphabet.Decode[b] = int8(i)
	}

	return &alphabet, nil
}

var (
	// BitcoinAlphabet is the bitcoin alphabet.
	BitcoinAlphabet, _ = NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

	// FlickrAlphabet is the flickr alphabet.
	FlickrAlphabet, _ = NewAlphabet("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
)

func EncodeAlphabet(src []byte, alphabet *Alphabet) string {
	return ""
}

func Encode(src []byte) string {
	return EncodeAlphabet(src, BitcoinAlphabet)
}

func DecodeAlphabet(src string, alphabet *Alphabet) ([]byte, error) {
	return []byte(""), nil
}

func Decode(src string) ([]byte, error) {
	return DecodeAlphabet(src, BitcoinAlphabet)
}
