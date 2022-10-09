package bobbygo

import (
	"fmt"
	"time"
)

func main() {
	currentTimestamp := time.Now()
	currentYear, currentMonth, _ := currentTimestamp.Date()
	currentLocation := currentTimestamp.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := time.Date(currentYear, currentMonth+1, 0, 23, 59, 59, 999999999, currentLocation)

	fmt.Println("firstOfMonth:", firstOfMonth)
	fmt.Println("lastOfMonth:", lastOfMonth)
}
