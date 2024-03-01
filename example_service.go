package main

const (
	bigBrickSize   = 5
	smallBrickSize = 2
)

// 使用大小砖填充空隙，判断给定数量的大小砖是否能刚好填满。
//
// 不能打碎砖块
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

// 给定一个字符串，返回一个新的字符串，其中每个字符都是原始字符串中该字符的长度编码。
func LengthEncode(input string) string {
	return ""
}

// 给定一个整数，返回其数字的反向编码。
func Reverse(input int) int {
	return 0
}
