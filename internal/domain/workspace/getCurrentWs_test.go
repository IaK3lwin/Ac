package workspace

import (
	"fmt"
	"testing"
)

func TestGetNameCurrentWs_get(t *testing.T) {
	fmt.Println("testando")
	currentWs, err := GetNameWsCurrent()

	if err != nil {
		fmt.Printf("[ERROR] : try get current ws: %s", err)
		t.Fail()
	}

	fmt.Println("currentWs: ", currentWs)
}
