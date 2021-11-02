package four

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func CalulcateHash(input string) (output int, err error) {
	genNumber := 609043
	toHash := input + fmt.Sprint(genNumber)

	hash := md5.Sum([]byte(toHash))
	hashyString := hex.EncodeToString(hash[:])

	fmt.Println(hashyString)

	if strings.HasPrefix(hashyString, "00000") {
		output = genNumber
	}

	return output, nil
}
