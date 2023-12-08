package datasource

import (
	"fmt"
	"testing"

	"github.com/RNekoCloud/deep-dive-http/helper"
)

func TestModify(t *testing.T) {
	data1 := City{
		Name:     "Solo",
		Province: "Jawa Tengah",
	}

	id := helper.UUIDGen()

	Store[id] = data1

	for _, val := range Store {
		fmt.Println("City:", val.Name)
	}
}
