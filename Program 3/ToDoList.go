package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* ToDoList... struct containing a list of type string */
type ToDoList struct {
	list []string
}

/* Returns a list of items */
func (l *ToDoList) showList() []string {
	return l.list
}

/* Adds a list item */
func (l *ToDoList) addListItem(item string) {
	l.list = append(l.list, item)
}

/* Moves an item from one location to another */
func (l *ToDoList) moveListItem(location int, destination int) {
	val := l.list[location]
	l.list = append(l.list[:location], l.list[location+1:]...)

	newSlice := make([]string, destination+1)
	copy(newSlice, l.list[:destination])
	newSlice[destination] = val

	l.list = append(newSlice, l.list[destination:]...)
}

/* Completes an item in the list by removing it */
func (l *ToDoList) completeListItem(item int) {
	l.list = append(l.list[:item], l.list[item+1:]...)
}

/* Ends list editing */
func (l *ToDoList) endList() string {
	return "\n* * * To Do List Ended * * *"
}

/* Main routine to run To Do List program */
func main() {
	list := &ToDoList{}
	fmt.Print("* * * To Do List * * *\n")

	for true {
		fmt.Print("\nEnter a command (Show, Add, Move, Complete) or End\n")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')

		// Process Show command
		if strings.Contains(command, "Show") {
			index := 1
			for i := 0; i < len(list.showList()); i++ {
				fmt.Printf("%d. %s", index, list.showList()[i])
				index++
			}

			// Process Add command
		} else if strings.Contains(command, "Add") {
			splitCommand := strings.SplitN(command, " ", 2)
			list.addListItem(splitCommand[1])

			// Process Complete command
		} else if strings.Contains(command, "Complete") {
			splitCommand := strings.SplitN(command, " ", 2)
			itemConv, err := strconv.Atoi(strings.TrimSuffix(splitCommand[1], "\n"))
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			if itemConv > 0 && itemConv <= len(list.showList()) {
				list.completeListItem(itemConv - 1)
			} else {
				fmt.Println("This item is not in the list.")
			}

			// Process Move command
		} else if strings.Contains(command, "Move") {
			splitCommand := strings.SplitN(command, " ", 3)
			location, item1err := strconv.Atoi(splitCommand[1])
			destination, item2err := strconv.Atoi(strings.TrimSuffix(splitCommand[2], "\n"))
			if item1err != nil {
				fmt.Println(item1err)
				os.Exit(0)
			}
			if item2err != nil {
				fmt.Println(item2err)
				os.Exit(0)
			}
			if (location > 0 && location <= len(list.showList())) && (destination > 0 && destination <= len(list.showList())) {
				list.moveListItem(location-1, destination-1)
			} else {
				fmt.Println("This item is not in the list.")
			}

			// Process End command
		} else if strings.Contains(command, "End") {
			fmt.Print(list.endList() + "\n")
			break

			// Process default
		} else {
			fmt.Println(command + " is an unrecognized command.")
		}
	}
}
