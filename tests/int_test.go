package tests

import (
	"fmt"
	"testing"
)

func TestIntAnd(t *testing.T) {
	var datasetId string = "1"
	var path string = "/Users/sunzhiqiang/Desktop/"
	path = path + string(datasetId)
	fmt.Println(path)
}
