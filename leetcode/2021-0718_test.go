package leetcode

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func compressString(S string) string {
	var buff bytes.Buffer
	for i := 0; i < len(S); {
		cnt := 1
		tmp := S[i]
		for i < len(S) {
			if i < len(S)-1 && S[i] != S[i+1] {
				i++
				break
			}
			if i < len(S)-1 {
				cnt++
			}
			i++
		}
		buff.WriteString(string(tmp))
		buff.WriteString(strconv.Itoa(cnt))
	}
	r := buff.String()
	if len(r) > len(S) {
		return S
	}

	return r
}

func Test_compressString(t *testing.T) {
	fmt.Println(compressString("rrrrrLLLLLPPPPPPRRRRRgggNNNNNVVVVVVVVVVDDDDDDDDDDIIIIIIIIIIlllllllAAAAqqqqqqqbbbNNNNffffff"))
}
