package helper

import (
	c_rand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"io"
	"math/rand"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Init() {
	var seed int64
	err := binary.Read(c_rand.Reader, binary.LittleEndian, &seed)
	if err != nil {
		panic(err)
	}
	rand.Seed(seed)
}

// Stretch makes stretched password using salt.
func Stretch(password, salt string) string {
	var b []byte
	s := sha256.New()
	for i := 0; i < 1000; i++ {
		io.WriteString(s, string(b)+password+salt)
		b = s.Sum(nil)
	}
	return base64.StdEncoding.EncodeToString(b)
}

// Salt returns random salt string.
func Salt(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
