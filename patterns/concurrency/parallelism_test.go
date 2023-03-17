package concurrency

import (
	"fmt"
	"testing"
)

func Test_Parallelism(t *testing.T) {
	m, err := MD5All("/tmp")
	if err != nil {
		fmt.Println(err)
		t.Skip()
		return
	}

	for k, v := range m {
		fmt.Printf("filename: %s  \nmd5:%x\n\n", k, v)
	}
}
