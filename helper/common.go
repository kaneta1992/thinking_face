package helper

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func DividedByteSlice(slice []byte, chunkSize int) [][]byte {
	ret := [][]byte{}
	sliceSize := len(slice)
	for i := 0; i < sliceSize; i += chunkSize {
		offsetIndex := Min(i+chunkSize, sliceSize)
		ret = append(ret, slice[i:offsetIndex])
	}
	return ret
}
