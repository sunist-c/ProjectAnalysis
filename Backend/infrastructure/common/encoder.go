package common

import (
	"crypto"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateMd5Len16 Generate a 16-byte md5 summary
func GenerateMd5Len16(str, salt string) (md5 string) {
	return GenerateMd5Len32(str, salt)[8:24]
}

// GenerateMd5Len32 Generate a 32-byte md5 summary
func GenerateMd5Len32(str, salt string) (md5 string) {
	h := crypto.MD5.New()
	h.Write([]byte(str + salt))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateRandomName Generate a random name
func GenerateRandomName() (name string) {
	return GenerateMd5Len16(time.Now().String(), strconv.Itoa(rand.Int()%114514))
}
