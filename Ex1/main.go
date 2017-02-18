package main

import (
	"encoding/json"
	"fmt"
	geo "golangexp/geometry"
	"log"
)

func main() {
	// first example: Marshalling json
	whl := geo.Wheel{geo.Circle{geo.Point{2, 4}, 5}, 20}
	whlJson, err := json.MarshalIndent(whl, "", "	")
	if err != nil {
		log.Fatalf("JSON Marshalling failed: %s", err)
	}
	fmt.Printf("%s\n\n", whlJson)
}
