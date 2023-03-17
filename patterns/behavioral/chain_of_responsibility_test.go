package behavioral

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ChainOfResponsibility(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdWordFilter{})
	assert.Equal(t, false, chain.Filter("ad"))

	chain.AddFilter(&PornographicWordFilter{})
	assert.Equal(t, true, chain.Filter("pornographic words..."))
}
