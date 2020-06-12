package chie

import (
	"bytes"
	"io/ioutil"
	"os"
	"regexp"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

const (
	// Encryption is the conversion factor used for encrypting a plaintext string
	// to TSC format (1)
	Encryption = 1

	// Decryption is the conversion factor used for decrypting a TSC file into
	// human-readable text (-1)
	Decryption = -1
)

// TSC is a struct that defines a TSC input. It contains the contents of the
// input, the size of the input, and the output string.
type TSC struct {
	contents      []byte
	size          int
	output        string
	allowComments bool
}

func (tsc *TSC) convert(conversionFactor int) *TSC {
	var key int

	encodingChar := tsc.getEncodingChar()

	if encodingChar != 0 {
		key = int(encodingChar) % 256
	} else {
		key = 7
	}

	key = key * conversionFactor

	for i := 0; i < len(tsc.contents); i++ {
		char := tsc.contents[i]

		if i != tsc.size/2 {
			char += byte(key)
		}

		tsc.output += string(char)
	}

	return tsc
}

func (tsc *TSC) getEncodingChar() byte {
	return tsc.contents[tsc.size/2]
}

func (tsc *TSC) stripComments(input []byte) []byte {
	pattern := regexp.MustCompile(`( |\r?\n)?\/\/.*`)
	return pattern.ReplaceAll(input, nil)
}

// Decrypt will decrypt a string from encrypted TSC to human-readable format
func (tsc *TSC) Decrypt() *TSC {
	return tsc.convert(Decryption)
}

// Encrypt will encrypt a plaintext string to TSC that's parseable by the game
func (tsc *TSC) Encrypt() *TSC {
	return tsc.convert(Encryption)
}

// FromString will fill the contents property of the TSC struct from the data
// inside of a given string
func (tsc *TSC) FromString(input string) {
	tsc.contents = []byte(input)

	if !tsc.allowComments {
		tsc.contents = tsc.stripComments(tsc.contents)
	}

	tsc.size = len(tsc.contents)
}

// ToString will return the output property of the TSC struct as a string
func (tsc *TSC) ToString() string {
	return tsc.output
}

// FromFile will read the input from a given file path
func (tsc *TSC) FromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if !tsc.allowComments {
		data = tsc.stripComments(data)
	}

	tsc.contents = data
	tsc.size = len(data)

	return nil
}

// ToFile will save the output to a given file path
func (tsc *TSC) ToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	var buffer bytes.Buffer

	enc := transform.NewWriter(&buffer, charmap.ISO8859_1.NewEncoder())
	defer enc.Close()

	_, err = enc.Write([]byte(tsc.output))
	if err != nil {
		return err
	}

	_, err = f.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// AllowComments will ensure that the comments will remain intact in the
// encrypted TSC file. If this method is not called, the TSC parser will
// remove all comments, since the base engine doesn't support them.
func (tsc *TSC) AllowComments() {
	tsc.allowComments = true
}

// NewTSCParser will return a new empty TSC struct
func NewTSCParser() *TSC {
	return &TSC{}
}
