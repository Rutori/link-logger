package controller

type Auth interface {
	Verify(authHeader string) bool
}
