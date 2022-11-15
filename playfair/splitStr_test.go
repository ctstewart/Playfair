package playfair

import (
	"reflect"
	"testing"
)

func Test_splitStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitStr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
