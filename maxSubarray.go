package algorithm

func findMaxCrossingSubarray(array []int, low int, mid int, high int) (lowIdx int, highIdx int, sum int) {
	leftSum := array[mid] - 1
	leftIdx := mid
	sum = 0
	for idx := mid; idx >= low; idx-- {
		sum = sum + array[idx]
		if sum > leftSum {
			leftSum = sum
			leftIdx = idx
		}
	}
	rightSum := array[mid+1] - 1
	rightIdx := mid + 1
	sum = 0
	for idx := mid + 1; idx <= high; idx++ {
		sum = sum + array[idx]
		if sum > rightSum {
			rightSum = sum
			rightIdx = idx
		}
	}
	return leftIdx, rightIdx, leftSum + rightSum
}

func FindMaxSubarray(array []int, low int, high int) (lowIdx int, highIdx int, sum int) {
	if low == high {
		return low, high, array[low]
	}

	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := FindMaxSubarray(array, low, mid)
	rightLow, rightHigh, rightSum := FindMaxSubarray(array, mid+1, high)
	// 使用 分治法 时，有时候 divide 出来的子问题，可能跟 原问题 不太相同
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(array, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	}
	if rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}
	return crossLow, crossHigh, crossSum
}
