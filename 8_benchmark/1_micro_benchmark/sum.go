package sum

import (
	"bytes"
	"os"
	"strconv"
)

func Sum(filename string) (ret int64, _ error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	for _, line := range bytes.Split(b, []byte("\n")) {
		num, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			return 0, err
		}

		ret += num
	}

	return ret, nil
}
