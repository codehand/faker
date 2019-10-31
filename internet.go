package faker

import (
	"fmt"
	"math/rand"
)

var internet = map[string][]string{
	"browser":       {"firefox", "chrome", "internetExplorer", "opera", "safari"},
	"domain_suffix": {"com", "biz", "info", "name", "net", "org", "io"},
	"http_method":   {"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
}

// DomainSuffix will generate a random domain suffix
func DomainSuffix() string {
	return getRandValue([]string{"internet", "domain_suffix"})
}

// HTTPMethod will generate a random http method
func HTTPMethod() string {
	return getRandValue([]string{"internet", "http_method"})
}

// IPv4Address will generate a random version 4 ip address
func IPv4Address() string {
	num := func() int { return 2 + rand.Intn(254) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

// IPv6Address will generate a random version 6 ip address
func IPv6Address() string {
	num := 65536
	return fmt.Sprintf("2001:cafe:%x:%x:%x:%x:%x:%x", rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num))
}

// MacAddress will generate a random mac address
func MacAddress() string {
	num := 255
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num))
}

// Username will genrate a random username based upon picking a random lastname and random numbers at the end
func Username() string {
	return getRandValue([]string{"person", "last"}) + replaceWithNumbers("####")
}
