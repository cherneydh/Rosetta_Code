/*
 * compiler.c
 * 
 * Copyright 2017 Daniel github: @cherneydh
 * 
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 * 
 * 
 */

//15 Puzzle implemented in Go
//Author: Daniel Cherney (github: @cherneydh)
//If you aren't using linux to run this, comment out the clear func



package main

import ("fmt"
		"math/rand"
		"os"
		"os/exec"
)

var board = [16]string{"0","0","0","0","0","0","0","0","0","0","0","0",
	"0","0","0","0"}
var correct = [16]string{" 1"," 2"," 3"," 4"," 5"," 6"," 7"," 8"," 9",
	"10","11","12","13","14","15","  "}
var read bool
var err bool
var command string
var temp string
var blank_location int
var row_location int

func buildboard(){
	var random int
	
	for i:=0;i<16;i++{
		random = rand.Intn(16)
		
		if board[random] == "0" && i == 15{
			board[random] = correct[i]
			blank_location = random
		}else if board[random] == "0"{
			board[random] = correct[i]
		}else{
			i--
		}		
	}	
}

func movepiece(){
	if read == false{
		fmt.Println("Enter command to give to the blank space.")
		fmt.Println("Commands: up = u, left = l, right = r, down = d")
		fmt.Println("I.E. typing \"u\" will swap the blank space with the "+
			"one above it")
		read = true
	}
	fmt.Println("Insert command: ")
	fmt.Scanln(&command)
	switch command {
	case "u":
			if ((blank_location-4)<0){
				err = true
				clear()
				printboard()
				fmt.Println("ERROR: OUT OF BOUNDS")
				movepiece()
				break;
			}
			temp = board[blank_location]
			board[blank_location] = board[blank_location-4]
			board[blank_location-4] = temp	
			blank_location = blank_location-4
	case "d":
			if ((blank_location+4)>15){
				err = true
				clear()
				printboard()
				fmt.Println("ERROR: OUT OF BOUNDS")
				movepiece()
				break;
			}
			temp = board[blank_location]
			board[blank_location] = board[blank_location+4]
			board[blank_location+4] = temp	
			blank_location = blank_location+4	
	case "l":
			if ((blank_location-1)<0){
				err = true
				clear()
				printboard()
				fmt.Println("ERROR: OUT OF BOUNDS")
				movepiece()
				break;
			}
			row_location = blank_location % 4
			if ((row_location-1)<0){
				err = true
				clear()
				printboard()
				fmt.Println("ERROR: OUT OF BOUNDS")
				movepiece()
				break;
			}
			temp = board[blank_location]
			board[blank_location] = board[blank_location-1]
			board[blank_location-1] = temp
			blank_location = blank_location-1	
	case "r":
			if ((blank_location+1)>15){
				err = true
				clear()
				printboard()
				fmt.Println("ERROR: OUT OF BOUNDS")
				movepiece()
				break;
			}
			row_location = blank_location % 4
			if ((row_location+1)>3){
				err = true
				clear()
				printboard()
				fmt.Println("ERROR: OUT OF BOUNDS")
				movepiece()
				break;
			}
			temp = board[blank_location]
			board[blank_location] = board[blank_location+1]
			board[blank_location+1] = temp
			blank_location = blank_location+1
	default:
			err = true
			clear()
			printboard()
			fmt.Println("ERROR: UNRECOGNIZED COMMAND")
			movepiece()
	}	
}

func checkboard() bool{
	for i:=0;i<16;i++{
		if board[i] != correct[i]{
			return false
		}
	}
	return true	
}

func printboard(){
	for i:=0; i<16; i=i+4{
		fmt.Println("______________________________________")
		fmt.Println("")
		fmt.Println("|   " + board[i] + "   |   " + board[i+1] + "   |   " + 
			board[i+2] + "   |   " + board[i+3] + "   |")
		if i == 12{
			fmt.Println("______________________________________")
			fmt.Println("")
		}	
	}	
			
	if (read == true && err == false){
		fmt.Println("\n\n\n\n\n\n\n")
	}else if (read == true && err == true){
		fmt.Println("\n\n\n\n\n\n")
		err = false
	}	
}

func clear(){
		c:=exec.Command("clear") // This function resets terminal to keep it clean
		c.Stdout=os.Stdout		 // If you aren't on linux, comment this out
		c.Run()					 // It will print newlines instead of clearing terminal
}		

func main(){
	var game_over = false
	buildboard()
	
	for game_over == false {
		clear()
		printboard()
		movepiece()
		game_over = checkboard()
	}	
	fmt.Println("Congrats! You won!")
}	
