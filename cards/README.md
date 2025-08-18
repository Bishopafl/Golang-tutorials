# Deck of Cards using Go

Creating a simple card shuffler in Go to learn more about it. So far, I have identified methods that will be needed and what those methods will do.

## Go Packages used
	- "fmt"
	- "math/rand"
	- "os"
	- "strings"
	- "time"

## Custom Methods

### 1. newDeck
This will create and return a list of playing cards. Essentially an array of strings.

### 2. print
Log out the contents of a deck of cards.

### 3. shuffle
Shuffles all the cards in a deck.

### 4. deal
Create a 'hand' of cards.

### 5. saveToFile
Save a list of cards to a file on the local machine.

### 6. newDeckFromFile
Load a list of cards from the local machine.

## Deck creation
for the creation of the deck, I am initilizing an empty deck. Then creating two different slices: 
    - One slice for `cardSuits` ("Spades", "Hearts", etc) 
    - One slice of `cardValues` ("Ace", "Two", etc)

Then I will have two foreach loops that will add a new card of 'value of suit' to the 'cards' deck.

## Testing

When initially starting this test, I had to run `go mod init cards` in the terminal to create the `go.mod` test file.

Once the file named `go.mod` is created in the project that will contain all of the functionality testing the deck methods.

Go does not handle tests well, and so you have to create a way to test that will create the test file, then delete that test file after the tests run successfully. 