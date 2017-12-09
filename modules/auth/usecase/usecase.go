package usecase

type UseCaseResult struct {
	Result interface{}
	Error  error
}

type AuthUseCase interface {
	GetAccessToken(email, password string) <-chan UseCaseResult
}
