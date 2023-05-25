package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery"
)

func main() {
	figure := figure.NewFigure("Enigma Camp", "standard", true)
	figure.Print()
	delivery.NewServer().Run()

}
