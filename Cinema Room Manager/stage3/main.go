package main

/*
[Cinema Room Manager - Stage 3/5: Tickets](https://hyperskill.org/projects/133/stages/711/implement)
-------------------------------------------------------------------------------
[Loops](https://hyperskill.org/learn/step/14707)
[Working with slices](https://hyperskill.org/learn/step/15935)
[Functional decomposition](https://hyperskill.org/learn/step/17506)
*/

import "fmt"

const (
	emptySeat  = "S"
	bookedSeat = "B"

	smallRoomMaxSeats = 60
	frontHalfPrice    = 10
	backHalfPrice     = 8
)

func initializeCinemaRoom(rows int, seatsPerRow int) [][]string {
	cinema := make([][]string, rows)
	for rowIdx := range cinema {
		cinema[rowIdx] = make([]string, seatsPerRow)
		for seatIdx := range cinema[rowIdx] {
			cinema[rowIdx][seatIdx] = emptySeat
		}
	}
	return cinema
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

func main() {
	var rows, seatsPerRow int
	fmt.Println("Enter the number of rows:")
	fmt.Scanln(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scanln(&seatsPerRow)

	cinema := initializeCinemaRoom(rows, seatsPerRow)
	displaySeatingArrangement(seatsPerRow, cinema)
	handleTicketPurchase(rows, seatsPerRow, cinema)
	displaySeatingArrangement(seatsPerRow, cinema)
}
