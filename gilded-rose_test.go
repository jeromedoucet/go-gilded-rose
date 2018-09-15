package main

import (
	"fmt"
	"testing"
)

func Test_GildedRose(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}

	// when
	result := fmt.Sprintf("%v", GildedRose(givenItems))

	// then
	if result != expectedResult {
		t.Errorf("expect the result to be %s but got, %s", expectedResult, result)
	}
}

func Test_GildedRose_backstage_sellin_10(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 9 22} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 10, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}

	// when
	result := fmt.Sprintf("%v", GildedRose(givenItems))

	// then
	if result != expectedResult {
		t.Errorf("expect the result to be %s but got, %s", expectedResult, result)
	}
}
