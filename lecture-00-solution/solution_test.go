package solution

import (
	"testing"

	"github.com/kyokomi/emoji"
)

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	want := emoji.Sprint("Hello :world_map:!")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
