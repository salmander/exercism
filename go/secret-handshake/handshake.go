package secret

import (
	"fmt"
	"strconv"
	"strings"
)

var p = fmt.Println

var secrets = map[string]string{
	"1000": "jump",
	"100":  "close your eyes",
	"10":   "double blink",
	"1":    "wink",
}

func Handshake(number int) []string {
	var returnString []string
	bin := strconv.FormatInt(int64(number), 2)
	p(number, "in binary:", bin)

	// Keep running the loop until the "bin"ary is blank
	for bin != "" {
		// Find the handshake replacement
		for b, str := range secrets {
			if strings.Contains(bin, b) == true {
				bin = strings.Replace(bin, b, "", 1)
				returnString = append(returnString, str)
			}
		}
	}
	return returnString
}
