package bt

import (
	"testing"
)

func TestStatusValues(t *testing.T) {
	if Success != 0 {
		t.Errorf("Success should be 0, got %d", Success)
	}

	if Failure != 1 {
		t.Errorf("Failure should be 1, got %d", Failure)
	}

	if Running != 2 {
		t.Errorf("Running should be 2, got %d", Running)
	}
}
