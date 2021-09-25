package query

//Query represent the query struct
type Query struct {
	Key       string
	Statement string
}

//Queries array map of query
type Queries map[string]Query

const (
	Topup = "Topup"

	TopupJoin = "TopupJoin"

	QueryNoSuchQuery = "QueryNoSuchQuery"
)

// GetQuery get query by value
func GetQuery(key string, value string) Query {

	q := make(Queries)

	q[Topup] = Query{
		Topup,
		`select sum(amount) From transactions where event_code_id in (?) and date_Added < ? and date_Added >= ?;`,
	}

	q[TopupJoin] = Query{
		TopupJoin,
		`select sum(t.amount) From transactions as t where t.event_code_id in (?) and t.date_Added < ? and t.date_Added >= ?  JOIN transaction_ledgers as tl on t.trasnaction_ledger_id = tl.id join pp_consumer as pp on tl.ref_id = pp.hash_id where pp.pp = 'Bill Payment';`,
	}

	return q[key]
}
