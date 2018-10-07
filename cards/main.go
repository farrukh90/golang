package main // this means executable file

import "fmt"

/*func main() {
	var card string = "Ace of Spades" // Creates a variable
	fmt.Println(card)
	var card2 string = "someting"
	fmt.Println(card2)
	var card4 bool = true
	fmt.Println(card4)
	var card5 int = 123
	fmt.Println(card5)

	card6 := "few" // This is same as above. This is only for new variables.
	// For changes in the future we can use = sign only.  No need to add a :
	//If : added it gives an error.

	fmt.Println(card6)
}
*/
func newCard() string { // This basically tells that we are trying to return a
	// string which is "Five of Diamonds"
	return "Five of Diamonds"
}

func newCard2() int {
	return 2
}
func newCard3() bool {
	return true
}
func newCard4() int {
	return 5
}
func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	//fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
