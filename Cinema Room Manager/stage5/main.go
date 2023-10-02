package main

/*
[Cinema Room Manager - Stage 5/5: Errors!](https://hyperskill.org/projects/133/stages/713/implement)
-------------------------------------------------------------------------------
[Type conversion and overflow](https://hyperskill.org/learn/step/18710)
*/

import "fmt"

const (
	emptySeat         = "S"
	bookedSeat        = "B"
	smallRoomMaxSeats = 60
	frontHalfPrice    = 10
	backHalfPrice     = 8

	percentageFactor = 100.0
)

func initializeCinemaRoom() (int, int, [][]string) {
	var rows, seatsPerRow int
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

	return rows, seatsPerRow, cinema
}

func displayMainMenu(rows int, seatsPerRow int, cinema [][]string) {
	for {
		fmt.Println("\n1. Show the seats")
		fmt.Println("2. Buy a ticket")
		fmt.Println("3. Statistics")
		fmt.Println("0. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			displaySeatingArrangement(seatsPerRow, cinema)
		case 2:
			handleTicketPurchase(rows, seatsPerRow, cinema)
		case 3:
			displayStatistics(rows, seatsPerRow, cinema)
		case 0:
			return
		}
	}
}

func displaySeatingArrangement(seatsPerRow int, cinema [][]string) {
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

func handleTicketPurchase(rows int, seatsPerRow int, cinema [][]string) {
	row, seat := getValidSeat(rows, seatsPerRow, cinema)
	bookSeat(row, seat, rows, cinema)
}

func getValidSeat(rows int, seatsPerRow int, cinema [][]string) (int, int) {
	for {
		fmt.Println("\nEnter a row number:")
		var row int
		fmt.Scanln(&row)

		fmt.Println("Enter a seat number in that row:")
		var seat int
		fmt.Scanln(&seat)

		if isValidSeat(row, seat, rows, seatsPerRow, cinema) {
			return row, seat
		}
	}
}

func isValidSeat(row int, seat int, rows int, seatsPerRow int, cinema [][]string) bool {
	if row <= 0 || row > rows || seat <= 0 || seat > seatsPerRow {
		fmt.Println("\nWrong input!")
		return false
	}

	if cinema[row-1][seat-1] == bookedSeat {
		fmt.Println("\nThat ticket has already been purchased!")
		return false
	}

	return true
}

func bookSeat(row int, seat int, rows int, cinema [][]string) {
	ticketPrice := calculateTicketPrice(rows, row)
	fmt.Printf("\nTicket price: $%d\n", ticketPrice)
	cinema[row-1][seat-1] = bookedSeat
}

func calculateTicketPrice(rows int, row int) int {
	if rows*rows > smallRoomMaxSeats && row > rows/2 {
		return backHalfPrice
	}
	return frontHalfPrice
}

func displayStatistics(rows int, seatsPerRow int, cinema [][]string) {
	purchasedTickets := countPurchasedTickets(cinema)
	totalNumberOfSeats := rows * seatsPerRow

	percentagePurchased := (float64(purchasedTickets) / float64(totalNumberOfSeats)) * percentageFactor

	currentIncome, totalIncome := calculateIncome(rows, cinema)

	fmt.Printf("\nNumber of purchased tickets: %d\n", purchasedTickets)
	fmt.Printf("Percentage: %.2f%%\n", percentagePurchased)
	fmt.Printf("Current income: $%d\n", currentIncome)
	fmt.Printf("Total income: $%d\n", totalIncome)
}

func countPurchasedTickets(cinema [][]string) int {
	count := 0
	for _, rowSeats := range cinema {
		for _, seat := range rowSeats {
			if seat == bookedSeat {
				count++
			}
		}
	}
	return count
}

func calculateIncome(rows int, cinema [][]string) (int, int) {
	currentIncome := 0
	totalIncome := 0
	for rowIdx, rowSeats := range cinema {
		rowPrice := calculateTicketPrice(rows, rowIdx+1)
		for _, seat := range rowSeats {
			totalIncome += rowPrice
			if seat == bookedSeat {
				currentIncome += rowPrice
			}
		}
	}
	return currentIncome, totalIncome
}

func main() {
	rows, seatsPerRow, cinema := initializeCinemaRoom()
	displayMainMenu(rows, seatsPerRow, cinema)
}
