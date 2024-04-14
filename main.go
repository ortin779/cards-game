package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	cards.saveToFile("test_file.txt")
}
