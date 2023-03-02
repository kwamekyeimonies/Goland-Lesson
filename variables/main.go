package main

import "fmt"

func main() {
	var heat_specific_capacity float32 = 22.798
	var (
		atomic_radius   int = 17
		mass_number     int = 23
		neutron_numbwer int = 5
	)

	fmt.Println("Atomic Number", atomic_radius)
	fmt.Println("Mass Number", mass_number)
	fmt.Println("Neutron Number", neutron_numbwer)
	fmt.Println("Specifi Heat Capacity", heat_specific_capacity)
}
