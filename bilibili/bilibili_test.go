package bilibili

import (
	"testing"
)

var client Client

func TestGetViewInfo(t *testing.T) {
	view, err := client.GetViewByBvid("BV15d4y1175g")
	if err != nil {
		t.Error(err)
		return
	}
	_, err = client.GetPlayUrlById(view.Arc.Aid, view.Arc.FirstCid)
	if err != nil {
		t.Error(err)
		return
	}
}
