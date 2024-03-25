package main

const (
	bigBrickSize   = 5
	smallBrickSize = 2
)

// 给定一个字符串，返回一个多组字符串组成的新字符串，其中每组字符都是原始字符串中对应字符的长度+字符。
// 如 aabb -> 2a2b
func LengthEncode(input string) string {
	return ""
}

// 使用函数 SlowEncrypt 字符串加密，尽可能快地返回加密后的字符串
func LengthEncrypt(input string) string {
	return ""
}

// 使用大小砖填充空隙，判断给定数量的大小砖是否能填满。
//
// 不能打碎砖块，无需使用全部的砖块。
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
