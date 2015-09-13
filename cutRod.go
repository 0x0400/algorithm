package algorithm

type plan struct {
	Length     int
	MaxRevenue int
	Solution   []int
}

func CutRod(length int, price []int) (int, []int) {
	results := make(map[int]plan)
	result := cutRodAux(length, price, results)
	return result.MaxRevenue, result.Solution
}

func cutRodAux(length int, price []int, results map[int]plan) plan {
	result := plan{length, 0, nil}
	for i := 1; i <= length; i++ {
		remainResult, ok := results[length-i]
		if !ok {
			remainResult = cutRodAux(length-i, price, results)
			results[length-i] = remainResult
		}
		if result.MaxRevenue < price[i]+remainResult.MaxRevenue {
			result.MaxRevenue = price[i] + remainResult.MaxRevenue
			result.Solution = append([]int{i}, remainResult.Solution...)
		}
	}
	return result
}
