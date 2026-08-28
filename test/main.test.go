package main

import (
	zermelogo "blacklizardcode/zermelo-go"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	fmt.Println(arg1)
	fmt.Println(arg2)
	code, err := strconv.Atoi(arg2)
	if err != nil {
		return
	}

	c := zermelogo.New(os.Args[1])

	access, err := c.Auth().CodeToBearer(code)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	c.Auth().SetBearerToken(access)
	fmt.Println(access)
}
