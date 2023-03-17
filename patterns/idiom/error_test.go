package idiom

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	buf := bytes.NewReader(b)
	p, err := parse(buf)

	assert.NotNil(t, p.PI)
	assert.NotNil(t, p.Uate)
	assert.NotNil(t, p.Mine)
	assert.NotNil(t, p.Too)
	assert.Nil(t, err)
}

func Test_Parse2(t *testing.T) {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	buf := bytes.NewReader(b)
	p, err := parse2(buf)

	assert.NotNil(t, p.PI)
	assert.NotNil(t, p.Uate)
	assert.NotNil(t, p.Mine)
	assert.NotNil(t, p.Too)
	assert.Nil(t, err)
}

func Test_Parse3(t *testing.T) {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	buf := bytes.NewReader(b)
	p, err := parse3(buf)

	assert.NotNil(t, p.PI)
	assert.NotNil(t, p.Uate)
	assert.NotNil(t, p.Mine)
	assert.NotNil(t, p.Too)
	assert.Nil(t, err)
}

func Test_Parse4(t *testing.T) {
	b := []byte{0x48, 0x61, 0x6f, 0x20, 0x43, 0x68, 0x65, 0x6e, 0x00, 0x00, 0x2c}
	r := bytes.NewReader(b)

	p := Person{}
	p.ReadName(r).ReadAge(r).ReadWeight(r)

	assert.NotNil(t, p.err)
	assert.Equal(t, io.EOF, p.err)
	assert.NotNil(t, p.Name)
	assert.NotNil(t, p.Age)
	assert.Equal(t, uint8(0), p.Weight)
}
