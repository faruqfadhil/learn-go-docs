package main

import (
	"fmt"
	"sort"
)

func main() {
	inventory := []int32{2, 5}

	order := int64(100)
	fmt.Println(maximumProfit(inventory, order))
}

func maximumProfit(inventory []int32, order int64) int64 {
	// Author: Faruq
	// 22 may 2022 11:30 WIB.

	suppliersCountMappedByStockProfit := map[int32]int{}
	for _, stock := range inventory {
		for i := int32(1); i <= stock; i++ {
			suppliersCountMappedByStockProfit[i]++
		}
	}

	uniqueStocksProfit := []int32{}
	for key, _ := range suppliersCountMappedByStockProfit {
		uniqueStocksProfit = append(uniqueStocksProfit, key)
	}
	sort.Slice(uniqueStocksProfit, func(i, j int) bool { return uniqueStocksProfit[i] > uniqueStocksProfit[j] })

	profit := int64(0)
	supply := int64(0)

	for _, stockProfit := range uniqueStocksProfit {
		if supply == order {
			break
		}

		if suppliersCount, ok := suppliersCountMappedByStockProfit[stockProfit]; ok {
			diff := order - supply
			if diff >= int64(suppliersCount) {
				supply += int64(suppliersCount)
				profit += (int64(stockProfit) * int64(suppliersCount))
			} else {
				supply += diff
				profit += (int64(stockProfit) * int64(diff))
			}
		}
	}
	return profit
}
