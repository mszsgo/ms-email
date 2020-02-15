package schema

type Query struct {
	Email *EmailQuery
}

type EmailQuery struct {
	SendText int64
}
