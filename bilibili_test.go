package bilibili

import (
	"testing"
)

func TestGetArchiveInfo(t *testing.T) {
	_, err := GetArchiveInfo("BV15d4y1175g")
	if err != nil {
		t.Error(err)
		return
	}
}
