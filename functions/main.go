package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYc":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	fmt.Println(add(42, 13))
	region, continent := location("Santa Monica")
	fmt.Printf("Matt lvies in %s, %s", region, continent)
}
