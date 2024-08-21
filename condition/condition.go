package condition

import "fmt"

func Condition(age int) {

	if age < 13 {
		fmt.Println("You are a child.")
	} else if age < 20 {
		fmt.Println("You are a teenager.")
	} else if age < 60 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

}
