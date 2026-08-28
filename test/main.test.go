package main

import (
	zermelogo "blacklizardcode/zermelo-go"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg1 := os.Args[1]
	fmt.Println(arg1)
	code, err := strconv.Atoi(arg1)
	if err != nil {
		return
	}

	c := zermelogo.New(os.Args[2])

	access, err := c.Auth().CodeToBearer(code)
	if err == nil {
		c.Auth().SetBearerToken(access)
	} else {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Println(strconv.Itoa(code))
	fmt.Println(access)
}