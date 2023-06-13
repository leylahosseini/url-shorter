package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateCode() string {
	// generate a unique code for the short URL
	h := sha256.New()
	h.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	code := hex.EncodeToString(h.Sum(nil))[:8]
	return code
}
