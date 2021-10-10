package main

import "listingQuery/Routes"

func main(){

	r := Routes.SetupRoutes()

	err := r.Run()

	if err != nil {
		return
	}
}

