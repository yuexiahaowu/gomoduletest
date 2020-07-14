package main

import "fmt"
import _ "time"


func main() {
	fmt.Println(" main ")
}

func init()  {
	fmt.Println("--------main  module init -------")
}