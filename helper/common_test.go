package helper

import "testing"

func TestMin(t *testing.T) {
	cases := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 1, y: 2, expected: 1},
		{x: 3, y: -5, expected: -5},
		{x: 0, y: 0, expected: 0},
	}

	for _, c := range cases {
		if Min(c.x, c.y) != c.expected {
			t.Errorf("This case is unexpected: %v\n", c)
		}
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 1, y: 2, expected: 2},
		{x: 3, y: -5, expected: 3},
		{x: 0, y: 0, expected: 0},
	}

	for _, c := range cases {
		if Max(c.x, c.y) != c.expected {
			t.Errorf("This case is unexpected: %v\n", c)
		}
	}
}

func TestDividedByteSlice(t *testing.T) {
	cases := []struct {
		slice    []byte
		size     int
		expected [][]byte
	}{
		{slice: []byte{1, 2, 3}, size: 1, expected: [][]byte{[]byte{1}, []byte{2}, []byte{3}}},
		{slice: []byte{1, 2, 3, 4, 5, 6}, size: 2, expected: [][]byte{[]byte{1, 2}, []byte{3, 4}, []byte{5, 6}}},
		{slice: []byte{1, 2, 3, 4, 5, 6, 7}, size: 3, expected: [][]byte{[]byte{1, 2, 3}, []byte{4, 5, 6}, []byte{7}}},
		{slice: []byte{1, 2, 3, 4, 5, 6, 7, 8}, size: 3, expected: [][]byte{[]byte{1, 2, 3}, []byte{4, 5, 6}, []byte{7, 8}}},
	}

	for _, c := range cases {
		ret := DividedByteSlice(c.slice, c.size)
		if len(ret) != len(c.expected) {
			t.Errorf("This case is unexpected chunk count: %v\n", c)
			break
		}
		for i, chunk := range ret {
			if len(chunk) != len(c.expected[i]) {
				t.Errorf("This case is unexpected value count: %v\n", c)
				break
			}
			for j, val := range chunk {
				if val != c.expected[i][j] {
					t.Errorf("This case is unexpected: %v\n", c)
					break
				}
			}
		}
	}
}
