package main

/*
[Cinema Room Manager - Stage 3/5: Tickets](https://hyperskill.org/projects/133/stages/711/implement)
-------------------------------------------------------------------------------
[Loops](https://hyperskill.org/learn/step/14707)
[Working with slices](https://hyperskill.org/learn/step/15935)
*/

import "fmt"

const (
	emptySeat  = "S"
	bookedSeat = "B"

	smallRoomMaxSeats = 60
	standardPrice     = 10
	reducedPrice      = 8
)

func main() {
	var rows, seatsPerRow, row, seat int

	fmt.Println("Enter the number of rows:")
	fmt.Scanln(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scanln(&seatsPerRow)

	cinema := make([][]string, rows)
	for rowIdx := range cinema {
		cinema[rowIdx] = make([]string, seatsPerRow)
		for seatIdx := range cinema[rowIdx] {
			cinema[rowIdx][seatIdx] = emptySeat
		}
	}

	fmt.Println("\nCinema:")
	fmt.Print("  ")
	for seatIdx := 1; seatIdx <= seatsPerRow; seatIdx++ {
		fmt.Print(seatIdx, " ")
	}
	fmt.Println()
	for rowIdx, rowSeats := range cinema {
		fmt.Printf("%d ", rowIdx+1)
		for _, seatStatus := range rowSeats {
			fmt.Print(seatStatus, " ")
		}
		fmt.Println()
	}

	for {
		fmt.Println("\nEnter a row number:")
		fmt.Scanln(&row)
		fmt.Println("Enter a seat number in that row:")
		fmt.Scanln(&seat)

		if row < 1 || row > rows || seat < 1 || seat > seatsPerRow {
			fmt.Println("\nWrong input!")
			continue
		}

		if cinema[row-1][seat-1] == bookedSeat {
			fmt.Println("\nThat ticket has already been purchased!")
			continue
		}
		break
	}

	isBigRoom := rows*seatsPerRow > smallRoomMaxSeats
	ticketPrice := standardPrice
	if isBigRoom && row > rows/2 {
		ticketPrice = reducedPrice
	}
	fmt.Printf("Ticket price: $%d\n", ticketPrice)
	cinema[row-1][seat-1] = bookedSeat

	fmt.Println("\nCinema:")
	fmt.Print("  ")
	for seatIdx := 1; seatIdx <= seatsPerRow; seatIdx++ {
		fmt.Print(seatIdx, " ")
	}
	fmt.Println()
	for rowIdx, rowSeats := range cinema {
		fmt.Printf("%d ", rowIdx+1)
		for _, seatStatus := range rowSeats {
			fmt.Print(seatStatus, " ")
		}
		fmt.Println()
	}
}
