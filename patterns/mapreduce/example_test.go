package mapreduce

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapStrToStr(t *testing.T) {
	var names = []string{"Tom", "Terry", "Mary"}
	v := MapStrToStr(names, func(s string) string {
		return strings.ToLower(s)
	})
	assert.Equal(t, []string{"tom", "terry", "mary"}, v)
}

func Test_MapStrToInt(t *testing.T) {
	var names = []string{"Tom", "Terry", "Mary"}
	v := MapStrToInt(names, func(s string) int {
		return len(s)
	})
	assert.Equal(t, []int{3, 5, 4}, v)
}

func Test_Reduce(t *testing.T) {
	var names = []string{"Tom", "Terry", "Mary"}
	v := Reduce(names, func(s string) int {
		return len(s)
	})
	assert.Equal(t, 12, v)
}

func Test_Filter(t *testing.T) {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odds := Filter(numbers, func(n int) bool {
		return n%2 == 1
	})
	assert.Equal(t, []int{1, 3, 5, 7, 9}, odds)
}
