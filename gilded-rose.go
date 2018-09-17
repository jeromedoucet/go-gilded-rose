package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

var currentItems = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

func main() {
	fmt.Println("OMGHAI!")
	fmt.Println(currentItems)
	currentItems = GildedRose(currentItems)
}

func GildedRose(items []Item) []Item {
	for i := 0; i < len(items); i++ {

		// no update at all when sulfuras
		if items[i].name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		// quality over 50 is forbidden
		if items[i].quality >= 50 {
			items[i] = updateSellin(items[i])
			continue
		}

		items[i] = updateSellin(items[i])

		if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
			items[i] = changeBackstageQuality(items[i])
			continue
		}

		if items[i].name == "Aged Brie" {
			items[i] = changeBrieQuality(items[i])
			continue
		}

		items[i] = changeNormalItemQuality(items[i])

	}
	return items
}

func updateSellin(item Item) Item {
	item.sellIn = item.sellIn - 1
	return item
}

func changeBackstageQuality(item Item) Item {
	if item.sellIn > 10 {
		item.quality = item.quality + 1
	} else if item.sellIn > 5 {
		item.quality = item.quality + 2
	} else if item.sellIn > 0 {
		item.quality = item.quality + 3
	} else {
		item.quality = 0
	}
	return item
}

func changeBrieQuality(item Item) Item {
	if item.sellIn >= 0 {
		item.quality = item.quality + 1
	} else {
		item.quality = item.quality + 2
	}
	return item
}

func changeNormalItemQuality(item Item) Item {
	if item.quality <= 0 {
		return item
	}
	if item.sellIn >= 0 {
		item.quality = item.quality - 1
	} else {
		item.quality = item.quality - 2
	}
	return item
}
