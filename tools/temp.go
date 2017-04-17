package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	t, err := GetCompressedTileFromLatLng(43.336665, -80.793457)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	fmt.Printf("%d\n", len(t))

	ioutil.WriteFile("./temp.png", t, 0644)
}
