package main

import "fmt"

func main() {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string 
	
	// 위 방식 대신 이런 식으로 short variable 사용 가능 
	empty := "test"
	numCars := 10
	temperature := 0.1 


	fmt.Printf(
		"%v %f %v %q %q %f %v\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
		empty,
		temperature,
		numCars,
	)

}