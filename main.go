package main

import (
	"log"

	"github.com/gedex/bp3d"
)

func main() {
	p := bp3d.NewPacker()

	// Add bins.
	p.AddBin(bp3d.NewBin("Small Bin", 10, 15, 20, 100))

	// Add items.
	p.AddItem(bp3d.NewItem("Item 1", 20, 2, 1, 2))
	p.AddItem(bp3d.NewItem("Item 2", 3, 3, 2, 3))

	// Pack items to bins.
	if err := p.Pack(); err != nil {
		log.Fatal(err)
	}

}
