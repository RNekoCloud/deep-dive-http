package helper

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := UUIDGen()
	fmt.Println("Generated ID:", uuid)
}
