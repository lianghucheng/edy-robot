package ai

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("test里面查看init情况", MinConsists, MinCardType)
	fmt.Println(MinConsists)
	fmt.Println(MinCardType)
	hands := []int{0, 4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48}
	i := 0
	j := len(hands) - 1
	for i < j {
		hands[i], hands[j] = hands[j], hands[i]
		i++
		j--
	}
	fmt.Println("反转hands：", hands)
	for _, v := range MinCardType {
		fmt.Println(GetDiscardHint(v, hands))
	}
}
