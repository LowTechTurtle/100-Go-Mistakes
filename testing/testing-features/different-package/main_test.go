package counter_test

// a difference package, counter test instead of counter
// we do this to test on the behavior, not the implementation
// if test heavily relies on the behavior then we can switch the implementation
// without anyone knowing, plus we could reuse the old tests
import (
	counter "_/home/turtle/100-Go-Mistakes/testing/testing-features/different-package"
	"testing"
)

func TestCount(t *testing.T) {
	if counter.Inc() != 1 {
		t.Errorf("expected 1")
	}
}
