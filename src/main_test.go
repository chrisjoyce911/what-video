package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNumber(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				a: 1,
				b: 2,
			},
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := AddNumber(tt.args.a, tt.args.b)

			assert.Equal(t, got, tt.want, "they should be equal")

			if got = AddNumber(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
