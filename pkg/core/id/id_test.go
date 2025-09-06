package id_test

import (
	"github.com/chengyu914001/go-common/pkg/core/id"
	"testing"
)

func TestGetPublicIDLen(t *testing.T) {
	if len(id.GetPublicID()) != 24 {
		t.Error("len(id.GetPublicID()) != 24")
	}
}

func TestGetPublicIDUnique(t *testing.T) {
	id1 := id.GetPublicID()
	id2 := id.GetPublicID()
	if id1 == id2 {
		t.Error("id1 == id2")
	}
}

func TestGetTraceID(t *testing.T) {
	if len(id.GetTraceID()) != 24 {
		t.Error("len(id.GetTraceID()) != 24")
	}
}

func TestGetTraceIDUnique(t *testing.T) {
	id1 := id.GetTraceID()
	id2 := id.GetTraceID()
	if id1 == id2 {
		t.Error("id1 == id2")
	}
}
