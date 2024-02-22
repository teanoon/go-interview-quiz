package main

const (
	bigBrickSize   = 5
	smallBrickSize = 2
)

// given the gap and a few bricks, check if they can fill the gap
func CanFillGap(smallQuantity, bigQuantity, totalGap int) bool {
	maxBigRequired := totalGap / bigBrickSize
	if maxBigRequired > bigQuantity {
		return false
	}

	gapLeftMaxBigRequired := totalGap % bigBrickSize
	if smallQuantity*smallBrickSize < gapLeftMaxBigRequired {
		return false
	}

	return gapLeftMaxBigRequired%bigBrickSize == 0
}

// given a string, encode it using length encoding
func LengthEncode(input string) string {
	return ""
}

// given an integer, reverse it
func Reverse(input int) int {
	return 0
}
