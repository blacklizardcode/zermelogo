package main

import (
	zermelogo "blacklizardcode/zermelo-go"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage: go run ./test <url> <code> <bearer> [schoolInSchoolYear]")
		os.Exit(2)
	}

	baseURL := os.Args[1]
	codeValue := os.Args[2]
	bearer := os.Args[3]
	schoolInSchoolYear := zermelogo.SchoolInSchoolYear("")
	if len(os.Args) > 4 {
		schoolInSchoolYear = zermelogo.SchoolInSchoolYear(os.Args[4])
	}

	c := zermelogo.New(baseURL)

	if bearer != "" {
		c.Auth().SetBearerToken(bearer)
	} else {
		code, err := strconv.Atoi(codeValue)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		access, err := c.Auth().CodeToBearer(code)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		c.Auth().SetBearerToken(access)
	}

	runAllTests(c, schoolInSchoolYear)

	if failedTests > 0 {
		fmt.Printf("%d test(s) failed\n", failedTests)
		os.Exit(1)
	}

	fmt.Println("all tests passed")
}
