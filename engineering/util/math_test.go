package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundFloat(t *testing.T) {
	type args struct {
		val       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test-1",
			args{0, 0},
			0,
		},
		{
			"test-2",
			args{0, 1},
			0,
		},
		{
			"test-3",
			args{1, 1},
			1,
		},
		{
			"test-4",
			args{12.3456789, 0},
			12,
		},
		{
			"test-5",
			args{12.3456789, 1},
			12.3,
		},
		{
			"test-6",
			args{12.3456789, 2},
			12.35,
		},
		{
			"test-7",
			args{12.3456789, 5},
			12.34568,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RoundFloat(tt.args.val, tt.args.precision), "RoundFloat(%v, %v)", tt.args.val, tt.args.precision)
		})
	}
}
