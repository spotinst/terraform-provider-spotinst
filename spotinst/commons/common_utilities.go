package commons

import (
	"fmt"
	"math/rand"
)

const (
	ResourceCachedIdentifierFormat = "%s-%d"
)

func GenerateCachedResourceId(resourceName string) string {
	id := fmt.Sprintf(string(ResourceCachedIdentifierFormat), resourceName, rand.Intn(100))
	return id
}
