package mouses

import "github.com/huyrun/counting-mouse/src/util"

// CountAliveMouse counts the number of alive mouse after n months
// originalAmount: the number of mouse at the beginning
// lifeSpan: the lifespan of a mouse
// n: the number of months
//
// Algorithm:
//   - mouses[x]: the number of mouse which is x months old, len(mouses) = lifeSpan
//   - mouses[0] is the number of mouse from original or are just born
//     mouses[0] = originalAmount
//   - for i = 0 to n-1, at the end of month i:
//     newMouses = sum(mouses[0:lifeSpan]) // mouses are lifeSpan-1 months old will increase 1 age, born new mouses then go away
//     mouses[j] = mouses[j-1] for j = lifeSpan-1 to 1 // increase the age of mouses by 1 month: the mouses are j-1 months old will be j months old
//     mouses[0] = newMouses // assign newMouses to mouses[0]
func CountAliveMouse(originalAmount, lifeSpan, months int) int {
	if lifeSpan == 0 {
		return 0
	}

	mouses := make([]int, lifeSpan)
	mouses[0] = originalAmount

	for i := 0; i < months; i++ {
		newMouses := util.SumArray(mouses)
		for j := lifeSpan - 1; j >= 1; j-- {
			mouses[j] = mouses[j-1]
		}
		mouses[0] = newMouses
	}
	return util.SumArray(mouses)
}
