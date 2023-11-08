package model

type ReadDBParams struct {
	FROM_TABLE      string
	ID              string   // unique search
	WHERE           []string //
	SEARCH_ARGUMENT string
	ORDER_BY        string
	SORT_DESC       bool //default ASC
}
