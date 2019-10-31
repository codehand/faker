package faker

import (
	"fmt"
	"testing"
)

// TestName ...
func TestName(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestName1",
		},
		{
			name: "TestName2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Name()
			if got == "" {
				t.Errorf("Name() = %v, want not empty", got)
			}
		})
	}
}

// TestNameRandom ...
func TestNameRandom(t *testing.T) {
	for i := 0; i < randIntRange(1, 100); i++ {
		t.Run(fmt.Sprintf("TestNameRandom%d", i+1), func(t *testing.T) {
			got := Name()
			if got == "" {
				t.Errorf("Name() = %v, want not empty", got)
			}
		})
	}
}
