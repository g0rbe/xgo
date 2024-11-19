package checksum

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

type SHA256 struct {
	b []byte
}

// Data256 returns the SHA256 checksum of data.
func Data256(data []byte) SHA256 {

	s := sha256.Sum256(data)

	return SHA256{b: s[:]}
}

// Path256 returns the SHA256 checksum of the file in path.
// Uses io.Copy() to be usable on large files.
func File256(path string) (SHA256, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return SHA256{}, err
	}
	defer file.Close()

	h := sha256.New()

	_, err = io.Copy(h, file)
	if err != nil {
		return SHA256{}, err
	}

	s := h.Sum(nil)

	return SHA256{b: s[:]}, nil
}

func (s SHA256) Bytes() []byte {
	return s.b
}

func (s SHA256) String() string {
	return hex.EncodeToString(s.b)
}
