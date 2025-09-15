package id_test

import (
	"testing"

	"github.com/chengyu914001/go-common/pkg/core/id"
)

func TestGenerateInnerIDLen(t *testing.T) {
	innerId := id.GenerateInnerID()
	if len(innerId) != 26 {
		t.Error(len(innerId), 26)
	}
}

func TestGenerateInnerIDUnique(t *testing.T) {
	innerIDSet := make(map[string]struct{})
	n := 1000
	for range n {
		innerIDSet[id.GenerateInnerID()] = struct{}{}
	}
	if len(innerIDSet) != n {
		t.Error(len(innerIDSet), n)
	}
}

func TestGeneratePublicID(t *testing.T) {
	publicID1 := id.GeneratePublicID()
	if len(publicID1) != 24 {
		t.Error(len(publicID1), 24)
	}
}

func TestGeneratePublicIDUnique(t *testing.T) {
	publicIDSet := make(map[string]struct{})
	n := 1000
	for range n {
		publicIDSet[id.GeneratePublicID()] = struct{}{}
	}
	if len(publicIDSet) != n {
		t.Error(len(publicIDSet), n)
	}
}
