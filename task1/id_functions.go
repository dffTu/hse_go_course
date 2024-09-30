package main

func first_function(name string) int {
	var result = 0
	for _, value := range name {
		result = (result*7 + int(value)) % 1013
	}
	return result
}

func second_function(name string) int {
	var result = 0
	for i, value := range name {
		result += i * int(value)
	}
	return result
}
