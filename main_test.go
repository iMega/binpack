package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/bom-d-van/binpacking"
	"github.com/gedex/bp3d"
)

type goods [4]int

func (g goods) GetWidth() int {
	return g[1]
}

func (g goods) GetHeight() int {
	return g[2]
}

func (g goods) GetDepth() int {
	return g[3]
}

func (g goods) GetWeight() int {
	return 10
}

func TestPack(t *testing.T) {
	items := []binpacking.Item{
		goods{1, 20, 100, 30},
		goods{2, 100, 20, 30},
		goods{3, 20, 100, 30},
		goods{4, 100, 20, 30},
		goods{5, 100, 20, 30},
		goods{6, 100, 100, 30},
		goods{7, 100, 100, 30},
	}
	boxes := []binpacking.Box{
		{Width: 2, Height: 2, Depth: 2, Weight: 100, Name: "Box-name-1"},
	}

	boxes[0].Items = []binpacking.BoxItem{
		{
			Item: goods{1, 1, 1, 10},
		},
	}

	got, err := binpacking.Pack(items)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%#v", got)
}

func Test_bp3d(t *testing.T) {

	items := []*bp3d.Item{
		bp3d.NewItem("Item 1", 1, 1, 1, 1),
		bp3d.NewItem("Item 2", 1, 1, 2, 1),
		bp3d.NewItem("Item 3", 1, 1, 1, 1),
		bp3d.NewItem("Item 4", 1, 1, 1, 1),
		bp3d.NewItem("Item 5", 1, 1, 1, 1),
		bp3d.NewItem("Item 6", 1, 1, 1, 1),
	}

	boxes := []*bp3d.Bin{
		bp3d.NewBin("6", 1, 1, 6, 4),
		bp3d.NewBin("5", 1, 1, 5, 5),
		bp3d.NewBin("4", 1, 1, 4, 4),
		bp3d.NewBin("3", 1, 1, 3, 3),
		bp3d.NewBin("2", 1, 1, 2, 2),
		bp3d.NewBin("1", 1, 1, 1, 1),
	}

	var weightItems float64
	for _, i := range items {
		weightItems += i.GetWeight()
	}

	for _, b := range boxes {
		fmt.Println(b.GetName())
		if b.GetMaxWeight() < weightItems {
			fmt.Printf("skip %s, mw: %f, wi: %f\n", b.GetName(), b.GetMaxWeight(), weightItems)
			continue
		}
		packer := bp3d.NewPacker()
		packer.AddBin(b)
		packer.AddItem(items...)
		if err := packer.Pack(); err != nil {
			log.Fatal(err)
		}

		if len(packer.UnfitItems) > 0 {
			fmt.Printf("UnfitItems: %d\n", len(packer.UnfitItems))
		} else {
			displayPacked(packer.Bins)
			break
		}

	}

}

func displayPacked(bins []*bp3d.Bin) {
	for _, b := range bins {
		fmt.Println(b)
		fmt.Println(" packed items:")
		for _, i := range b.Items {
			fmt.Println("  ", i)
		}
		fmt.Println("")
	}
}
