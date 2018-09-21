package main

import (
	"fmt"
	"testing"
)

func Test_GildedRose_with_original_set(t *testing.T) {
	// given
	var originalItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	// when
	givenItems := GildedRose(originalItems)

	// then
	if fmt.Sprint(givenItems) != expectedResult {
		t.Errorf("Expect to have %s when calling GildedRose() with no args but got %s", expectedResult, fmt.Sprint(givenItems))
	}
}

func Test_GildedRose_with_backstage_when_concert_in_10_days(t *testing.T) {
	// given
	var originalItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 10, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 9 22} {Conjured Mana Cake 2 5}]"

	// when
	givenItems := GildedRose(originalItems)

	// then
	if fmt.Sprint(givenItems) != expectedResult {
		t.Errorf("Expect to have %s when calling GildedRose() with no args but got %s", expectedResult, fmt.Sprint(givenItems))
	}
}

func Test_GildedRose_with_backstage_when_concert_in_5_days(t *testing.T) {
	// given
	var originalItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 5, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 4 23} {Conjured Mana Cake 2 5}]"

	// when
	givenItems := GildedRose(originalItems)

	// then
	if fmt.Sprint(givenItems) != expectedResult {
		t.Errorf("Expect to have %s when calling GildedRose() with no args but got %s", expectedResult, fmt.Sprint(givenItems))
	}
}

func Test_GildedRose_when_elixir_out_of_date(t *testing.T) {
	// given
	var originalItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", -1, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose -2 5} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	// when
	givenItems := GildedRose(originalItems)

	// then
	if fmt.Sprint(givenItems) != expectedResult {
		t.Errorf("Expect to have %s when calling GildedRose() with no args but got %s", expectedResult, fmt.Sprint(givenItems))
	}
}

func Test_GildedRose_when_brie_out_of_date(t *testing.T) {
	// given
	var originalItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", -1, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie -2 2} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	// when
	givenItems := GildedRose(originalItems)

	// then
	if fmt.Sprint(givenItems) != expectedResult {
		t.Errorf("Expect to have %s when calling GildedRose() with no args but got %s", expectedResult, fmt.Sprint(givenItems))
	}
}

func Test_GildedRose_when_backstage_out_of_date(t *testing.T) {
	// given
	var originalItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", -1, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", -1, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie -2 2} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert -2 0} {Conjured Mana Cake 2 5}]"

	// when
	givenItems := GildedRose(originalItems)

	// then
	if fmt.Sprint(givenItems) != expectedResult {
		t.Errorf("Expect to have %s when calling GildedRose() with no args but got %s", expectedResult, fmt.Sprint(givenItems))
	}
}

func Test_GildedRose_when_brie_has_value_50(t *testing.T) {
	// given
	var originalItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 50},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 50} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	// when
	givenItems := GildedRose(originalItems)

	// then
	if fmt.Sprint(givenItems) != expectedResult {
		t.Errorf("Expect to have %s when calling GildedRose() with no args but got %s", expectedResult, fmt.Sprint(givenItems))
	}
}
