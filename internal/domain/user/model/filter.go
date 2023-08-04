package model

const (
	SortAsc  = "ASC"
	SortDesc = "DESC"

	OperatorEqual   = "eq"
	OperatorOr      = "or"
	OperatorRange   = "range"
	OperatorIn      = "in"
	OperatorNotIn   = "not in"
	OperatorIsNull  = "is_null"
	OperatorNot     = "not"
	ConditionIfNull = "ifnull"
)

type Filter struct {
	Sorts        []Sort        `json:"sort"`
	Pagination   Pagination    `json:"pagination"`
	FilterFields []FilterField `json:"filter"`
}

type Sort struct {
	Field     string `json:"field"`
	Order     string `json:"order"`
	Condition string `json:"condition"`
}

type Pagination struct {
	Page     int `json:"page" `
	PageSize int `json:"pageSize"`
}

type FilterField struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}
