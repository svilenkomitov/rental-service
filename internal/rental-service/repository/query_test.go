package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_buildQuery(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		expected := "SELECT rentals.*, users.first_name, users.last_name, rentals.price_per_day * rentals.sleeps as price FROM rentals JOIN users ON rentals.user_id=users.id WHERE rentals.price_per_day * rentals.sleeps >= 1 AND rentals.price_per_day * rentals.sleeps <= 2000000 AND rentals.id IN (5, 7, 15) AND st_distance(geography(st_makepoint(rentals.lng,rentals.lat)), geography(st_makepoint(-117.279999, 32.830002))) * 0.000621371192 < 100 ORDER BY price LIMIT 5 OFFSET 1"

		actual := NewQueryBuilder().
			Select("SELECT rentals.*, users.first_name, users.last_name, rentals.price_per_day * rentals.sleeps as price FROM rentals JOIN users ON rentals.user_id=users.id").
			Where("rentals.price_per_day * rentals.sleeps >= 1").
			Where("rentals.price_per_day * rentals.sleeps <= 2000000").
			Where("rentals.id IN (5, 7, 15)").
			Where("st_distance(geography(st_makepoint(rentals.lng,rentals.lat)), geography(st_makepoint(-117.279999, 32.830002))) * 0.000621371192 < 100").
			OrderBy("price").
			Limit(5).
			Offset(1).
			Build()

		assert.Equal(t, expected, actual)
	})

	// TODO: add more tests
}
