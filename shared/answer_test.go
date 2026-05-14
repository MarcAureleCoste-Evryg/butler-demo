package shared

import "testing"

func TestGetAnswer(t *testing.T) {
	want := 42
	if got := GetAnswer(); got != want {
		t.Errorf("GetAnswer() = %v, want %v", got, want)
	}
}
