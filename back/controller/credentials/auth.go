package credentials

type Auth interface {
	Verify(authHeader string) bool
}
