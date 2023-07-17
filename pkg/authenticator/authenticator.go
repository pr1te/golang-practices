package authenticator

type Strategies struct {
	Local   *LocalStrategy
	Session *SessionStrategy
}

type Authenticator struct {
	Strategies *Strategies
}

type AuthUser struct {
	UserID    uint
	ProfileID uint
}

func New(sessionStrategy *SessionStrategy, localStrategy *LocalStrategy) *Authenticator {
	strategies := &Strategies{
		Local:   localStrategy,
		Session: sessionStrategy,
	}

	authenticator := &Authenticator{
		Strategies: strategies,
	}

	return authenticator
}
