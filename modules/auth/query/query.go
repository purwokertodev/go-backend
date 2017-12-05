package query

type QueryResult struct {
	Result interface{}
	Error  error
}

type IdentityQuery interface {
	FindByEmail(email string) <-chan QueryResult
}
