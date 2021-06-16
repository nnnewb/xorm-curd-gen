package parse_test

import (
	"testing"

	"github.com/nnnewb/xorm-curd-gen/pkg/parse"
)

type M struct {
}

type (
	T struct{}
	S struct{}
)

func TestParseFile(t *testing.T) {
	result, err := parse.ParseFile("parse_test.go")
	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Errorf("parse fail but no errors returned")
	}

	if result != nil && len(result.Models) < 3 {
		t.Errorf("parse result contains models less than actual declared")
	}
}
