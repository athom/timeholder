package timeholder

import (
	"testing"
	"time"
)

func TestTimeHold(t *testing.T) {
	SetDuration(2 * time.Millisecond)

	key1 := "1234"
	key2 := "5678"
	ok := Hold(key1)
	if ok {
		t.Errorf("first time shold not be held")
	}

	time.Sleep(1 * time.Millisecond)
	ok = Hold(key1)
	if !ok {
		t.Errorf("within time shold be held")
	}
	ok = Hold(key2)
	if ok {
		t.Errorf("first time shold not be held")
	}

	time.Sleep(2 * time.Millisecond)
	ok = Hold(key1)
	if ok {
		t.Errorf("out of time shold not be held")
	}

	ok = Hold(key2)
	if ok {
		t.Errorf("first time shold not be held")
	}

}
