package faker

import (
	"fmt"
	"testing"
)

// TestPhoneNumer ...
func TestPhoneNumer(t *testing.T) {
	for i := 0; i < randIntRange(1, 100); i++ {
		t.Run(fmt.Sprintf("TestPhoneNumer%d", i+1), func(t *testing.T) {
			got := PhoneNumer()
			if len(got) != 11 {
				t.Errorf("PhoneNumer() = %v, want not invalid", got)
			}
		})
	}
}
