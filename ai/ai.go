package ai

var MinSolo = []int{0}
var MinPair = []int{1,0}
var MinTrip = []int{2, 1, 0}
var MinTripSingle = []int{8,2,1,0}
var MinTripDouble = []int{9,8,2,1,0}
var MinBomb = []int{3, 2, 1, 0}
var MinConsist = []int{16,12,8,4,0}
var MinConsists [][]int
var MinconPair = []int{9,8,5,4,1,0}
var MinPlane = []int{6,5,4,2,1,0}
var MinPlaneSingle = []int{13,8,6,5,4,2,1,0}
var MinPlaneDouble = []int{18,17,9,8,6,5,4,2,1,0}

var MinCardType = [][]int{
	MinSolo,
	MinPair,
	MinTrip,
	MinTripSingle,
	MinTripDouble,
	MinBomb,
	MinconPair,
	MinPlane,
	MinPlaneSingle,
	MinPlaneDouble,
}

func init() {
	MinConsists = append(MinConsists, MinConsist)
	for i:=0;i<12;i++{
		first := []int{MinConsists[i][0] + 4}
		MinConsists = append(MinConsists, append(first, MinConsists[i]...))
	}

	MinCardType = append(MinCardType, MinConsists...)
}

func GetDiscard(hands []int) []int {
	for i:=len(MinConsists)-1;i>=0;i-- {
		if consist := GetDiscardHint(MinConsists[i], hands); len(consist) != 0 {
			for j := len(consist) - 1;j>=0;j-- {
				if consist[j][0] < 44 {
					return consist[j]
				}
			}
		}
	}

	//todo: 过滤炸弹
	if ret := GetDiscardHint(MinPlaneDouble, hands); len(ret) != 0 {
		return ret[0]
	} else if ret := GetDiscardHint(MinPlaneSingle, hands); len(ret) != 0 {
		return ret[0]
	} else if ret := GetDiscardHint(MinPlane, hands); len(ret) != 0 {
		return ret[0]
	} else if ret := GetDiscardHint(MinconPair, hands); len(ret) != 0 {
		return ret[0]
	} else if ret := GetDiscardHint(MinTripDouble, hands); len(ret) != 0 {
		return ret[0]
	} else if ret := GetDiscardHint(MinTripSingle, hands); len(ret) != 0 {
		return ret[0]
	} else if ret := GetDiscardHint(MinTrip, hands); len(ret) != 0 {
		return ret[0]
	} else if ret := GetDiscardHint(MinPair, hands); len(ret) != 0 {
		return ret[0]
	} else if ret := GetDiscardHint(MinSolo, hands); len(ret) != 0 {
		return ret[0]
	}

	return nil
}