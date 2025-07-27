package lsfDraw

import (
	"fmt"
	"math"
	"os"

	"golang.org/x/term"
)

type GridList struct {
	TotalColumns  int
	ColumnWidths  []int
	TotalElements int
	ElementWidths []int
	MaxWidth      int
	ExtraWidth    int
	MultiRow      bool
}

func InitializeGrid(fes *FileEntries) GridList {

	//Make new lengths with padding
	paddedSizes := make([]int, 0)
	for _, e := range *fes {
		paddedSizes = append(paddedSizes, e.RawLen+2)
	}

	//Cell budget & cost -- if Multirow is needed
	var cellCost int
	var multirow bool = false
	var extraWidth int = 0
	for _, v := range paddedSizes {
		cellCost += v
	}
	cellBudget := GetCliWidth()
	if cellCost > cellBudget {
		multirow = true
	} else {
		extraWidth = cellBudget - cellCost
	}

	return GridList{
		TotalColumns:  1,
		ColumnWidths:  []int{100},
		TotalElements: len(*fes),
		ElementWidths: paddedSizes,
		MaxWidth:      cellBudget,
		ExtraWidth:    extraWidth,
		MultiRow:      multirow,
	}
}

func GetCliWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n\n", err)
		os.Exit(1)
	}
	return width
}

// 125, 100, 75, 125, 80, 90, 100, 200 | 400
func CalcColumnSize(elements []int, widthMax int) []int {
	const widthPadding = 2
	columnSizes := make([]int, 0)
	widthBudget := int(widthMax)

	for _, itemLength := range elements {
		widthCost := int(itemLength) + widthPadding
		if widthBudget >= widthCost {
			columnSizes = append(columnSizes, widthCost)
			widthBudget -= widthCost
		} else {
			qtyColumns := len(columnSizes)
			qtyElements := len(elements)
			iterations := int(math.Ceil(float64(qtyElements) / float64(qtyColumns)))

			newColumnSizes := make([]int, qtyColumns)
			for col := range qtyColumns {
				maxVal := elements[col]
				for i := 1; i < iterations; i++ {
					offset := col + i*qtyColumns
					if offset >= qtyElements {
						break
					}
					if elements[offset] > maxVal {
						maxVal = elements[offset]
					}
				}
				newColumnSizes[col] = maxVal + widthPadding
			}
			return newColumnSizes
		}
	}
	return columnSizes
}
