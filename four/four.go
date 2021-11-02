package four

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func CalulcateHash(input string) (output int, err error) {
	var genNumber int

	for {
		toHash := input + fmt.Sprint(genNumber)

		hash := md5.Sum([]byte(toHash))
		hashyString := hex.EncodeToString(hash[:])
		fmt.Println("genNumber: ", genNumber)

		if strings.HasPrefix(hashyString, "00000") {
			output = genNumber
			break
		} else {
			genNumber += 1
		}
	}

	return output, nil
}
