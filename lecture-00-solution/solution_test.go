package solution

import "testing"

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	want := "Hello :world_map:"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
