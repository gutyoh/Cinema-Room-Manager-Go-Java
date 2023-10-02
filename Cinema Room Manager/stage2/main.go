package main

/*
[Cinema Room Manager - Stage 2/5: Sold!](https://hyperskill.org/projects/133/stages/710/implement)
-------------------------------------------------------------------------------
[Control statements](https://hyperskill.org/learn/step/16235)
[Arithmetic operations](https://hyperskill.org/learn/step/16679)
[String formatting](https://hyperskill.org/learn/step/16860)
*/

import "fmt"

const (
	smallRoomMaxSeats = 60
	standardPrice     = 10
	reducedPrice      = 8
)

func main() {
	var rows, seatsPerRow int
	fmt.Println("Enter the number of rows:")
	fmt.Scanln(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scanln(&seatsPerRow)

	totalSeats := rows * seatsPerRow
	var totalIncome int

	if totalSeats <= smallRoomMaxSeats {
		totalIncome = totalSeats * standardPrice
	} else {
		frontRows := rows / 2
		totalIncome = frontRows * seatsPerRow * standardPrice

		backRows := rows - frontRows
		totalIncome += backRows * seatsPerRow * reducedPrice
	}

	fmt.Println("\nTotal income:")
	fmt.Printf("$%d", totalIncome)
}
