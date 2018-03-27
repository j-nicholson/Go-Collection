package main

/*
	Program that creates a tic tac toe game
*/

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
	Draws tic tac toe board
*/
func drawBoard(board []string) {
	fmt.Println("\n   |   |")
	fmt.Println(" " + board[1] + " | " + board[2] + " | " + board[3])
	fmt.Println("   |   |")
	fmt.Println("-----------")
	fmt.Println("   |   |")
	fmt.Println(" " + board[4] + " | " + board[5] + " | " + board[6])
	fmt.Println("   |   |")
	fmt.Println("-----------")
	fmt.Println("   |   |")
	fmt.Println(" " + board[7] + " | " + board[8] + " | " + board[9])
	fmt.Println("   |   |")
}

/*
	Chooses first player
*/
func chooseFirstPlayer() string {
	return "computer"
}

/*
	Asks for a rematch upon end of game
*/
func rematch() string {
	return "Would you like to play again? (y/n) "
}

/*
	Places player move at specified board location
*/
func move(board []string, letter string, move int) {
	board[move] = letter
}

/*
	Checks winning combinations to see if player has won the game
*/
func isWinner(board []string, letter string) bool {
	return board[7] == letter && board[8] == letter && board[9] == letter ||
		board[4] == letter && board[5] == letter && board[6] == letter ||
		board[1] == letter && board[2] == letter && board[3] == letter ||
		board[7] == letter && board[4] == letter && board[1] == letter ||
		board[8] == letter && board[5] == letter && board[2] == letter ||
		board[9] == letter && board[6] == letter && board[3] == letter ||
		board[7] == letter && board[5] == letter && board[3] == letter ||
		board[9] == letter && board[5] == letter && board[1] == letter
}

/*
	Returns a copy of the current tic tac toe board
*/
func copyBoard(board []string) []string {
	var boardCopy []string
	for i := range board {
		boardCopy = append(boardCopy, board[i])
	}
	return boardCopy
}

/*
	Checks to see if a specified space is free to make a legal move
*/
func isSpaceFree(board []string, move int) bool {
	if board[move] == " " {
		return true
	}
	return false
}

/*
	Returns legal player move
*/
func getPlayerMove(board []string) int {
	reader := bufio.NewReader(os.Stdin)
	moveInt := 0
	var err error
	for (moveInt != 1 && moveInt != 2 && moveInt != 3 && moveInt != 4 && moveInt != 5 && moveInt != 6 && moveInt != 7 && moveInt != 8 && moveInt != 9) || !isSpaceFree(board, moveInt) {
		fmt.Print("Enter a move (1-9): ")
		move, _ := reader.ReadString('\n')
		moveInt, err = strconv.Atoi(strings.TrimSuffix(move, "\n"))
		if err != nil {
			fmt.Print(err)
		}
	}

	return moveInt
}

/*
	AI to make the computer choose a move if a space is free based on the provided move set
*/
func computerMoveAI(board []string, moveList []int) int {
	var possibleMoves []int
	for i := range moveList {
		if isSpaceFree(board, moveList[i]) {
			possibleMoves = append(possibleMoves, moveList[i])
		}
	}

	if len(possibleMoves) > 0 {
		start := 1
		end := len(possibleMoves)
		if end-start > 0 {
			randomIndex := start + rand.Intn(end-start)
			return possibleMoves[randomIndex]
		}
		return possibleMoves[0]
	}
	return 0
}

/*
	AI to make the computer choose a move based on lists of previous losses
*/
func getComputerMoveBasedOnLosingBoard(board []string, losingBoard []interface{}, computerLetter string) int {
	var playerLetter string
	if computerLetter == "X" {
		playerLetter = "O"
	} else {
		playerLetter = "X"
	}
	losingMoves := []interface{}{[]string{}}
	for i := range losingBoard {
		losingMoves = append(losingMoves, losingBoard[i])
	}

	s := reflect.ValueOf(losingMoves)
	for i := 0; i < s.Len(); i++ {
		for j := 1; j < s.Index(i).Elem().Len(); j++ {
			if s.Index(i).Elem().Index(j).Interface() != computerLetter {
				copy := copyBoard(board)
				if isSpaceFree(copy, j) {
					move(copy, playerLetter, j)
					if isWinner(copy, playerLetter) {
						return j
					}
				}
			}
		}
	}
	return computerMoveAI(board, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

/*
	AI to make the computer choose a random move if available
*/
func getComputerMove(board []string) int {
	return computerMoveAI(board, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

/*
	Checks to see if the game board's spaces are full
*/
func isBoardFull(board []string) bool {
	isFull := true
	for i := 1; i < 10; i++ {
		if isSpaceFree(board, i) {
			isFull = false
		}
	}
	return isFull
}

/*
	Reads loss boards from computerlosses.txt file if it exists
*/
func readLossesFromFile() []interface{} {
	losses := []interface{}{[]byte{}}
	fileLosses, err := ioutil.ReadFile("computerlosses.txt")
	for i := range fileLosses {
		losses = append(losses, fileLosses[i])
	}
	jsonData := json.Unmarshal(fileLosses, &losses)
	if jsonData != nil {
		fmt.Print(jsonData)
	}
	if err != nil {
		fmt.Print(err)
	}
	return losses
}

/*
	Writes loss boards to computerlosses.txt file if the player wins
*/
func writeLossesToFile(badMoves []interface{}) {
	losses := []interface{}{[]byte{}}
	for i := range badMoves {
		losses = append(losses, badMoves[i])
	}
	jsonData, _ := json.Marshal(losses)
	err := ioutil.WriteFile("computerlosses.txt", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

/*
	Checks to see if a specified file exists
*/
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

/*
	Main routine to run tic tac toe program
*/
func main() {
	fmt.Println("* * * Welcome to Tic Tac Toe * * *")
	computerWinStatus := true
	badMoves := []interface{}{[]string{}}

	for true {
		gameBoard := []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "}
		playerLetter := "O"
		computerLetter := "X"

		turn := chooseFirstPlayer()
		fmt.Print("The " + turn + " will go first.")
		gameIsPlaying := true

		// Checks if a losses.txt file exists and reads bad move boards from it
		exists, existsErr := fileExists("computerlosses.txt")
		if existsErr != nil {
			fmt.Print(existsErr)
		}
		if exists {
			fileLosses := readLossesFromFile()
			for i := range fileLosses {
				badMoves = append(badMoves, fileLosses[i])
			}
			computerWinStatus = false
		}

		for gameIsPlaying {
			if turn == "player" {
				drawBoard(gameBoard)
				playerMove := getPlayerMove(gameBoard)
				move(gameBoard, playerLetter, playerMove)

				if isWinner(gameBoard, playerLetter) {
					drawBoard(gameBoard)

					// Writes bad moves to losses.txt file if the player wins
					badMoves = append(badMoves, copyBoard(gameBoard))
					writeLossesToFile(badMoves)
					computerWinStatus = false

					fmt.Print("You have won!\n")
					gameIsPlaying = false
				} else {
					if isBoardFull(gameBoard) {
						drawBoard(gameBoard)
						fmt.Print("This game ends in a tie!\n")
						gameIsPlaying = false
					} else {
						turn = "computer"
					}
				}
			} else {
				drawBoard(gameBoard)
				// Determines computer move based on its current win status
				if !computerWinStatus {
					computerMove := getComputerMoveBasedOnLosingBoard(gameBoard, badMoves, computerLetter)
					move(gameBoard, computerLetter, computerMove)
				} else {
					computerMove := getComputerMove(gameBoard)
					move(gameBoard, computerLetter, computerMove)
				}

				if isWinner(gameBoard, computerLetter) {
					drawBoard(gameBoard)
					fmt.Print("The computer has won. You lose.\n")
					gameIsPlaying = false
				} else {
					if isBoardFull(gameBoard) {
						drawBoard(gameBoard)
						fmt.Print("This game ends in a tie!\n")
						gameIsPlaying = false
					} else {
						turn = "player"
					}
				}
			}
		}
		fmt.Print(rematch())
		reader := bufio.NewReader(os.Stdin)
		rematch, _ := reader.ReadString('\n')
		if strings.Contains(rematch, "n") {
			break
		}
	}
}
