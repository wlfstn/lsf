package lsfDraw

import (
	"fmt"
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

	var totalColumns int = 0
	var columnWidths []int
	var cellCost int = 0
	var multirow bool = false
	var extraWidth int = 0

	//Make new lengths with padding
	paddedSizes := make([]int, 0)
	for _, e := range *fes {
		paddedSizes = append(paddedSizes, e.RawLen+2)
	}

	//Cell budget & cost -- if Multirow is needed
	for _, v := range paddedSizes {
		cellCost += v
	}
	extraWidth = GetCliWidth()
	if cellCost > extraWidth {
		multirow = true
		for _, sizeE := range paddedSizes {
			if extraWidth > sizeE {
				totalColumns += 1
				columnWidths = append(columnWidths, sizeE)
				extraWidth -= sizeE
			} else {
				fmt.Printf("Columns: %v | {%v}\n", len(columnWidths), columnWidths)
				fmt.Printf("Remaining Budget: %v \n", extraWidth)
				break
			}
		}
	} else {
		totalColumns = len(paddedSizes)
		columnWidths = paddedSizes
		extraWidth -= cellCost
	}

	return GridList{
		TotalColumns:  totalColumns,
		ColumnWidths:  columnWidths,
		TotalElements: len(*fes),
		ElementWidths: paddedSizes,
		MaxWidth:      GetCliWidth(),
		ExtraWidth:    extraWidth,
		MultiRow:      multirow,
	}
}

func (grid *GridList) RowOverflowBalance() bool {
	var GridColumnReduction bool = false

	//loop through columns & elements to determine resizing [grid.ElementWidths]
	//example 3 columns(A,B,C), 8 entities(100,125,75,125,80,90,100,200), width of 400
	//This would leave the remaining balance of 100, the 4th element would resize column A to 125, subtract 25 from the remaining balance
	var c int = 0
	var resizeCost int = 0
	for i := grid.TotalColumns; i < grid.TotalElements; i++ {

		if grid.ElementWidths[i] > grid.ColumnWidths[c] {
			resizeCost = grid.ElementWidths[i] - grid.ColumnWidths[c]
			fmt.Printf("Resize Cost: %v | Balance: %v\n", resizeCost, grid.ExtraWidth)
			if grid.ExtraWidth >= resizeCost {
				grid.ColumnWidths[c] = grid.ElementWidths[i]
				grid.ExtraWidth -= resizeCost
			} else {
				fmt.Println("Grid Column must reduce")
				GridColumnReduction = true
				break
			}
		}

		if c+1 < grid.TotalColumns {
			c++
		} else {
			c = 0
		}
	}

	if GridColumnReduction && grid.TotalColumns > 1 {
		if grid.ColumnWidths[len(grid.ColumnWidths)-1] > (resizeCost + grid.ExtraWidth) {
			fmt.Println("One Column wasn't enough")
		}
		fmt.Println("Reducing Grid Columns by 1")
		grid.TotalColumns -= 1
		grid.ColumnWidths = grid.ColumnWidths[:len(grid.ColumnWidths)-1]
	}
	return GridColumnReduction
}

func (grid *GridList) CalcRowBudget() {
	fmt.Printf("Last Balance: %v\n", grid.ExtraWidth)
	var newCost int
	for _, e := range grid.ColumnWidths {
		newCost += e
	}
	grid.ExtraWidth = grid.MaxWidth - newCost
}

func (grid *GridList) ResetColumnWidths() {
	for i := range grid.ColumnWidths {
		fmt.Printf("Resetting Column %v from %v to %v\n", i, grid.ColumnWidths[i], grid.ElementWidths[i])
		grid.ColumnWidths[i] = grid.ElementWidths[i]
	}
}

func (grid *GridList) Print(fes *FileEntries) {
	var column int = 0
	var paddingQty int = 2

	for _, e := range *fes {
		if grid.MultiRow {
			if column >= grid.TotalColumns {
				fmt.Println()
				column = 0
			}
			paddingQty = grid.ColumnWidths[column] - e.RawLen
		}

		// in go %-*s means to left align with padding
		if e.IsDir {
			fmt.Printf("\033[34m%s\033[0m%-*s", e.Name, paddingQty, "")
		} else {
			fmt.Printf("%s%-*s", e.Name, paddingQty, "")
		}
		column++
	}
	fmt.Println()
}

func GetCliWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n\n", err)
		os.Exit(1)
	}
	return width
}
