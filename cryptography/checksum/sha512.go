// Compute and check SHA256/SHA512 message digest
package checksum

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
)

type SHA512 struct {
	b []byte
}

// Data512 returns the SHA512 checksum of data.
func Data512(data []byte) SHA512 {

	s := sha512.Sum512(data)
	return SHA512{b: s[:]}
}

// Path512 returns the SHA512 checksum of the file in path.
func Path512(path string) (SHA512, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return SHA512{}, err
	}
	defer file.Close()

	h := sha512.New()

	_, err = io.Copy(h, file)
	if err != nil {
		return SHA512{}, err
	}

	return SHA512{b: h.Sum(nil)}, nil
}

func (s SHA512) Bytes() []byte {
	return s.b
}

func (s SHA512) String() string {
	return hex.EncodeToString(s.b)
}
