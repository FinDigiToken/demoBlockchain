package main

import "../core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 10000 FDT to Lance")
	bc.SendData("Send 10000 FDT to Swen")
	bc.Print()
}