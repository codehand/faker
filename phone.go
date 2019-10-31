package faker

import (
	"regexp"
	"strings"
)

// PhoneNumer ...
func PhoneNumer() string {
	s := "###########"
	first := true
	for _, sm := range regexp.MustCompile(`#`).FindAllStringSubmatch(s, -1) {
		if first {
			s = strings.Replace(s, sm[0], string(randIntRange(0, 9)), 1)
			first = false
		} else {
			s = strings.Replace(s, sm[0], string(randIntRange(0, 9)), 1)
		}
	}
	return s
}
