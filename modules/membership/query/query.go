package query

type QueryResult struct {
	Result interface{}
	Error  error
}

type MemberQuery interface {
	FindByEmail(email string) <-chan QueryResult
}
