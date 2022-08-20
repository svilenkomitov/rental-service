package repository

import (
	"testing"
)

func TestQuery(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		queries := make(map[QueryKey]interface{})
		queries[PRICE_MIN_KEY] = 1
		queries[PRICE_MAX_KEY] = 2000000
		queries[IDS_KEY] = []int{5, 7, 15}
		queries[NEAR_KEY] = []float64{-117.279999, 32.830002}
		queries[LIMIT_KEY] = 5
		queries[OFFSET_KEY] = 1
		queries[SORT_KEY] = "price"

		t.Log(buildQuery(queries))
	})
}
