package bobbygo

import (
	"fmt"
	"strings"
)

func DecodeChiper(value string) {
	character := "abcdefghijklmnopqrstuvwxyz"
	kar := strings.Split(character, "")
	vSplit := strings.Split(value, " ")
	for i := 1; i < len(character); i++ {
		fmt.Print(i , ". ")
		for p := 0; p < len(vSplit); p++{
			for _, x := range strings.Split(vSplit[p], ""){
				res := strings.Index(character, x)
				index := res - i 
				if index < 0 {
					index = 26 - (index *  -1)
				}
				fmt.Print( kar[ index % 26])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
func EncodeChiper(value string) {
	character := "abcdefghijklmnopqrstuvwxyz"
	kar := strings.Split(character, "")
	vSplit := strings.Split(value, " ")
	for i := 1; i < len(character); i++ {
		fmt.Print(i , ". ")
		for p := 0; p < len(vSplit); p++{
			for _, x := range strings.Split(vSplit[p], ""){
				res := strings.Index(character, x)
				fmt.Print( kar[ (res - i ) % 26])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}