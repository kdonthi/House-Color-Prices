package housecolors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestTotalPrice(t *testing.T) {
	testTable := []struct{
		name string
		houseColorPrices []HouseColorPrices
		expectedResult int
	}{
		{
			houseColorPrices: []HouseColorPrices{
				{
					red: 1,
					green: 2,
					blue: 3,
				},
				{
					red: 17,
					green: 64,
					blue: 3,
				},
			},
			expectedResult: 4,
		},
		{
			houseColorPrices: []HouseColorPrices{
				{
					red: 1,
					green: 2,
					blue: 3,
				},
				{
					red: 5,
					green: 2,
					blue: 1,
				},
				{
					red: 17,
					green: 18,
					blue:4,
				},
			},
			expectedResult: 7,
		},
		{
			houseColorPrices: []HouseColorPrices{
				{
					red: 1,
					green: 2,
					blue: 3,
				},
				{
					red: 46,
					green: 21,
					blue: 86,
				},
				{
					red: 5,
					green: 2,
					blue: 1,
				},
				{
					red: 17,
					green: 18,
					blue:4,
				},
			},
			expectedResult: 31,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := lowestTotalPrice(testCase.houseColorPrices)
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}