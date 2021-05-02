package bobbygo

import "fmt"

func calculatePrima(n uint64) {
	var k uint64 = 0
	var on uint64 = 0
	for i := uint64(2); i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i)
			k = n / i
			if n/i == 1 {
				break
			} else if k%i == 0 {
				fmt.Print("^")
				for j := uint64(0); j <= k; j++ {
					if n%i == 0 {
						on++
						n /= i
						if n == 1 {
							fmt.Print(on)
							break
						} else if n%i != 0 {
							fmt.Print(on, " x ")
							break
						}
					}
				}
				k = n
			} else {
				fmt.Print(" x ")
			}
			n = k
		}
		on = 0
	}
}

func PrimeFactorization(value uint64) {
	calculatePrima(value)
}
