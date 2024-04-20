package retrier

import (
	"encoding/hex"
	"strings"
)

func ParseHexString(str string) ([]byte, error) {
	if strings.HasPrefix(str, "0x") {
		str = strings.TrimPrefix(str, "0x")
	}

	return hex.DecodeString(str)
}
