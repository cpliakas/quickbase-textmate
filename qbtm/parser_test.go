package qbtm_test

import (
	"testing"

	"github.com/cpliakas/quickbase-textmate/qbtm"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		s  string
		ex string
	}{
		{"()", "()"},
		{"(Text t, Number p, Text d)", "($1, $2, $3)"},
		{"(<any> x, ...)", "($1)"},
		{"(Boolean condition1, <any> result1, ..., <any> else-result)", "($1, $2)"},
		{"(UserList ul, UserList ul1,UserList ul2 ..)", "($1, $2)"},
		{"(User u ..)", "()"},
	}

	for _, tt := range tests {
		if actual := qbtm.ParseArgs(tt.s); actual != tt.ex {
			t.Errorf("got %q, expected %q", actual, tt.ex)
		}
	}
}
