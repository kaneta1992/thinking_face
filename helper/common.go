package helper

import (
	"io/ioutil"
	"math/rand"
	"path/filepath"
)

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

func DivideByteSlice(slice []byte, chunkSize int) [][]byte {
	ret := [][]byte{}
	sliceSize := len(slice)
	for i := 0; i < sliceSize; i += chunkSize {
		offsetIndex := Min(i+chunkSize, sliceSize)
		ret = append(ret, slice[i:offsetIndex])
	}
	return ret
}

// https://qiita.com/tanksuzuki/items/7866768c36e13f09eedb
// TODO: Cache for each submodule update
func DirWalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, DirWalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		filename := file.Name()
		// Do not list dot files
		if []rune(filename)[0] != rune('.') {
			paths = append(paths, filepath.Join(dir, filename))
		}
	}

	return paths
}

func RandomSelect(data []string) string {
	l := len(data)
	return data[rand.Intn(l)]
}
