package faker

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println(Name())
	t.Log(Name())
}
