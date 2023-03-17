package idiom

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestErrorWrap(t *testing.T) {
	readFile := func() error {
		_, err := ioutil.ReadFile("/tmp/pkg_error.go")
		if err != nil {
			return errors.Wrap(err, "read failed")
		}
		return nil
	}

	err := readFile()
	assert.NotNil(t, err)
	fmt.Printf("Error: %+v", err)
}
