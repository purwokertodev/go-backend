package query

type QueryResult struct {
	Result interface{}
	Error  error
}

type MembershipQuery interface {
	FindByEmail(email string) <-chan QueryResult
}
