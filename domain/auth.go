package domain

type AuthError struct {
	Err error
}

func (e *AuthError) Error() string {
	return e.Err.Error()
}
