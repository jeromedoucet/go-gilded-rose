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

func Test_GildedRose_backstage_sellin_5(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 4 23} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 5, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}

	// when
	result := fmt.Sprintf("%v", GildedRose(givenItems))

	// then
	if result != expectedResult {
		t.Errorf("expect the result to be %s but got, %s", expectedResult, result)
	}
}

func Test_GildedRose_aged_brie_sellin_minus_1(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie -2 2} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", -1, 0},
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

func Test_GildedRose_backstage_sellin_minus_1(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert -2 0} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", -1, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}

	// when
	result := fmt.Sprintf("%v", GildedRose(givenItems))

	// then
	if result != expectedResult {
		t.Errorf("expect the result to be %s but got, %s", expectedResult, result)
	}
}

func Test_GildedRose_elixir_sellin_minus_1(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose -2 5} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", -1, 7},
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

func Test_GildedRose_vest_sellin_minus_1(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest -2 18} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", -1, 20},
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

func Test_GildedRose_cake_sellin_minus_1(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake -2 4}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		Item{"Conjured Mana Cake", -1, 6},
	}

	// when
	result := fmt.Sprintf("%v", GildedRose(givenItems))

	// then
	if result != expectedResult {
		t.Errorf("expect the result to be %s but got, %s", expectedResult, result)
	}
}

func Test_GildedRose_sulfuras_sellin_minus_1(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros -1 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", -1, 80},
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

func Test_GildedRose_elixir_quality_50(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 50} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 50},
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

func Test_GildedRose_brie_quality_50(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 50} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 50},
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

func Test_GildedRose_vest_quality_50(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 50} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 50},
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

func Test_GildedRose_backstage_quality_50(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 50} {Conjured Mana Cake 2 5}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 50},
		Item{"Conjured Mana Cake", 3, 6},
	}

	// when
	result := fmt.Sprintf("%v", GildedRose(givenItems))

	// then
	if result != expectedResult {
		t.Errorf("expect the result to be %s but got, %s", expectedResult, result)
	}
}

func Test_GildedRose_cake_quality_50(t *testing.T) {
	// given
	expectedResult := "[{+5 Dexterity Vest 9 19} {Aged Brie 1 1} {Elixir of the Mongoose 4 6} {Sulfuras, Hand of Ragnaros 0 80} {Backstage passes to a TAFKAL80ETC concert 14 21} {Conjured Mana Cake 2 50}]"

	var givenItems = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		Item{"Conjured Mana Cake", 3, 50},
	}

	// when
	result := fmt.Sprintf("%v", GildedRose(givenItems))

	// then
	if result != expectedResult {
		t.Errorf("expect the result to be %s but got, %s", expectedResult, result)
	}
}
