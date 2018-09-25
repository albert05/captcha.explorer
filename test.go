package main

import (
	"github.com/otiai10/gosseract"
	"fmt"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage("./img/1.png")
	t, err := client.Text()
	fmt.Println(t, err)
}
