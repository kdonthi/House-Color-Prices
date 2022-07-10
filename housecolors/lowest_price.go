package housecolors

import (
	"fmt"
)

type HouseColorPrices struct {
	red int
	green int
	blue int
}

type Color int

const (
	Red Color = iota
	Green
	Blue
)

// base case -> last element in list => just return
// +1 -> find lowest price of other two colors

func lowestTotalPrice(houseColorPrices []HouseColorPrices) int {
	var greenPrice int
	var bluePrice int
	var redPrice int
	
	if len(houseColorPrices) == 1 {
		greenPrice = houseColorPrices[0].green
		bluePrice = houseColorPrices[0].blue
		redPrice = houseColorPrices[0].red
	} else {
		priceCount := len(houseColorPrices)
		
		greenPrice = recursiveFunc(houseColorPrices, Green, 0, priceCount - 1)
		bluePrice = recursiveFunc(houseColorPrices, Blue, 0, priceCount - 1)
		redPrice = recursiveFunc(houseColorPrices, Red, 0, priceCount - 1)
	}
	
	var result int
	if greenPrice < bluePrice {
		result = greenPrice
	} else {
		result = bluePrice
	}

	if redPrice < result {
		return redPrice
	} else {
		return result
	}
}

func recursiveFunc(prices []HouseColorPrices, houseColor Color, index int, lastIndex int) int {
	currentHousePrices := prices[index]
	if index == lastIndex {
		return houseColorPrice(currentHousePrices, houseColor)
	}

	if houseColor == Red {
		green := recursiveFunc(prices, Green, index + 1, lastIndex)
		blue := recursiveFunc(prices, Blue, index + 1, lastIndex)
		if blue < green {
			return blue + currentHousePrices.red
		} else {
			return green + currentHousePrices.red
		}
	} else if houseColor == Green {
		red := recursiveFunc(prices, Red, index + 1, lastIndex)
		blue := recursiveFunc(prices, Blue, index + 1, lastIndex)
		if blue < red {
			return blue + currentHousePrices.green
		} else {
			return red + currentHousePrices.green
		}
	} else if houseColor == Blue {
		red := recursiveFunc(prices, Red, index + 1, lastIndex)
		green := recursiveFunc(prices, Green, index + 1, lastIndex)
		if red < green {
			return red + currentHousePrices.blue
		} else {
			return green + currentHousePrices.blue
		}
	} else {
		panic(fmt.Sprintf("incorrect color: %v", houseColor))
	}
}

func houseColorPrice (h HouseColorPrices, houseColor Color) int {
	if houseColor == Red {
		return h.red
	} else if houseColor == Blue {
		return h.blue
	} else if houseColor == Green {
		return h.green
	} else {
		panic(fmt.Sprintf("incorrect color: %v", houseColor))
	}
}

func min(i1 int, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}