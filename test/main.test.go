package main

import (
	zermelogo "blacklizardcode/zermelo-go"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg1 := os.Args[1] // set url
	arg2 := os.Args[2] // set code
	arg3 := os.Args[3] // set bearer
	c := zermelogo.New(os.Args[1])

	if arg3 != "" {
		c.Auth().SetBearerToken(arg3)
	} else {

		fmt.Println(arg1)
		fmt.Println(arg2)
		code, err := strconv.Atoi(arg2)
		if err != nil {
			return
		}

		access, err := c.Auth().CodeToBearer(code)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
		fmt.Println(access)

		c.Auth().SetBearerToken(access)
	}

	appointments, err := c.GetAppointments(zermelogo.AppointmentsParams{User: "~me", Start: 1789084800, End: 1789167600})
	if err != nil {
		fmt.Println("err")
		fmt.Println(err.Error())
	}
	fmt.Println(appointments[len(appointments) - 1].Locations)

}
