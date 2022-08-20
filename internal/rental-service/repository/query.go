package repository

import (
	"fmt"
	"strconv"
	"strings"
)

type DBClause string

const (
	WHERE    DBClause = "WHERE"
	LIMIT    DBClause = "LIMIT"
	OFFSET   DBClause = "OFFSET"
	ORDER_BY DBClause = "ORDER BY"
	AND      DBClause = "AND"
)

type QueryKey string

const (
	PRICE_MIN_KEY QueryKey = "price_min"
	PRICE_MAX_KEY QueryKey = "price_max"
	LIMIT_KEY     QueryKey = "limit"
	OFFSET_KEY    QueryKey = "offset"
	IDS_KEY       QueryKey = "ids"
	NEAR_KEY      QueryKey = "near"
	SORT_KEY      QueryKey = "sort"
)

type Query []string

var (
	PRICE_MIN Query = []string{"price_min", string(WHERE), " rentals.price_per_day * rentals.sleeps >= %v "}
)

func buildQuery(queries map[QueryKey]interface{}) string {

	// TODO: use price_per_day or calculate rentals.price_per_day * rentals.sleeps ?
	baseQuery := `SELECT rentals.*, users.first_name, users.last_name, rentals.price_per_day * rentals.sleeps as price FROM rentals JOIN users ON rentals.user_id=users.id `
	sortQuery := ""
	limitQuery := ""
	offsetQuery := ""

	whereConditions := []string{}

	for key, value := range queries {
		switch key {
		case PRICE_MIN_KEY:
			{
				whereConditions = append(whereConditions, fmt.Sprintf(" rentals.price_per_day * rentals.sleeps >= %v ", value))
			}
		case PRICE_MAX_KEY:
			{
				whereConditions = append(whereConditions, fmt.Sprintf(" rentals.price_per_day * rentals.sleeps <= %v ", value))
			}
		case IDS_KEY:
			{
				var ids []string
				for _, i := range value.([]int) {
					ids = append(ids, strconv.Itoa(i))
				}

				v := strings.Join(ids, ", ")
				whereConditions = append(whereConditions, fmt.Sprintf(" rentals.id IN (%v) ", v))
			}
		case NEAR_KEY:
			{
				var near []string
				for _, i := range value.([]float64) { // TODO: do not loop (order is important)
					near = append(near, fmt.Sprintf("%f", i))
				}

				v := strings.Join(near, ", ")
				whereConditions = append(whereConditions, fmt.Sprintf(" st_distance(geography(st_makepoint(rentals.lng,rentals.lat)), geography(st_makepoint(%v))) * 0.000621371192 < 100 ", v))
			}
		case SORT_KEY:
			{
				sortQuery += string(ORDER_BY) + fmt.Sprintf(" %v ", value)
			}
		case LIMIT_KEY:
			{
				limitQuery += string(LIMIT) + fmt.Sprintf(" %v ", value)
			}
		case OFFSET_KEY:
			{
				offsetQuery += string(OFFSET) + fmt.Sprintf(" %v ", value)
			}
		}
	}

	var whereQuery string
	if len(whereConditions) > 0 {
		whereQuery = string(WHERE) + strings.Join(whereConditions, string(AND))
	}

	return baseQuery + whereQuery + sortQuery + limitQuery + offsetQuery
}
