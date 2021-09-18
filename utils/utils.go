package utils

import "fmt"

func TrimBuffer(buffer []byte) ([]byte, error) {
	eof := 0
	for i := len(buffer) - 1; i >= 0; i-- {
		if buffer[i] != 0 {
			eof = i
			break
		}
	}
	if eof == 0 {
		return nil, fmt.Errorf("could not find end of buffer")
	}
	return buffer[:eof+1], nil
}
