package repository

import (
	"fmt"
	"strings"
)

type DBClause string

const (
	WHERE    DBClause = " WHERE "
	LIMIT    DBClause = " LIMIT "
	OFFSET   DBClause = " OFFSET "
	ORDER_BY DBClause = " ORDER BY "
	AND      DBClause = " AND "
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

type QueryBuilder struct {
	selectStatement  string
	whereConditions  []string
	whereStatement   string
	orderByStatement string
	limitStatement   string
	offsetStatement  string
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (b *QueryBuilder) Select(statement string) *QueryBuilder {
	b.selectStatement = statement
	return b
}

func (b *QueryBuilder) Where(condition string) *QueryBuilder {
	b.whereConditions = append(b.whereConditions, condition)
	return b
}

func (b *QueryBuilder) OrderBy(key string) *QueryBuilder {
	b.orderByStatement = string(ORDER_BY) + fmt.Sprintf("%s", key)
	return b
}

func (b *QueryBuilder) Limit(value int) *QueryBuilder {
	b.limitStatement = string(LIMIT) + fmt.Sprintf("%d", value)
	return b
}

func (b *QueryBuilder) Offset(value int) *QueryBuilder {
	b.offsetStatement = string(OFFSET) + fmt.Sprintf("%d", value)
	return b
}

func (b *QueryBuilder) Build() string {
	if len(b.whereConditions) > 0 {
		b.whereStatement = string(WHERE) + strings.Join(b.whereConditions, string(AND))
	}
	return b.selectStatement + b.whereStatement + b.orderByStatement + b.limitStatement + b.offsetStatement
}
