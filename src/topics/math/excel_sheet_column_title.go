package math

import "bytes"
func convertToTitle(n int) string {
	buffer := new (bytes.Buffer)

	var res []byte
	for n > 0 {
		n --
		res = append(res, byte(n % 26))
		n = n / 26
	}
	for i := len(res)  -1; i >= 0 ;i -- {
		buffer.WriteByte(res[i] + 65)
	}

	return buffer.String()
}