package main

import "fmt"

func main() {
	var name string
	var projects int

	fmt.Scanln(&name)
	fmt.Scanln(&projects)

	var totalProjects = projects * 3

	fmt.Printf("The architect %s will need %d hours to complete %d project/s.", name, totalProjects, projects)
}