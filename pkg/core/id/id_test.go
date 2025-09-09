package id_test

import (
	"testing"

	"github.com/chengyu914001/go-common/pkg/core/id"
)

func TestGeneratePrimaryKetIDLen(t *testing.T) {
	if len(id.GeneratePrimaryKetID()) != 24 {
		t.Error("len(id.GetPublicID()) != 24")
	}
}

func TestGeneratePrimaryKetIDUnique(t *testing.T) {
	id1 := id.GeneratePrimaryKetID()
	id2 := id.GeneratePrimaryKetID()
	if id1 == id2 {
		t.Error("id1 == id2")
	}
}

func TestGenerateTraceID(t *testing.T) {
	if len(id.GenerateTraceID()) != 24 {
		t.Error("len(id.GetTraceID()) != 24")
	}
}

func TestGenerateTraceIDUnique(t *testing.T) {
	id1 := id.GenerateTraceID()
	id2 := id.GenerateTraceID()
	if id1 == id2 {
		t.Error("id1 == id2")
	}
}
