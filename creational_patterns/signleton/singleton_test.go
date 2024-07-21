package signleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("expected pointer to Signleton after calling GetInstance(), not nil")
	}
	expectedCounter := counter1.AddOne()

	counter2 := GetInstance().GetCount()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but got a different instance")
	}
}
