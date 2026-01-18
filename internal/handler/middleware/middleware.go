package middleware

type Middleware struct {
	token string
}

func NewMiddleware(token string) *Middleware {
	return &Middleware{
		token: token,
	}
}
