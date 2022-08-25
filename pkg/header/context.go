package header

type contextKey string

var ForwardHeaders = []string{jwtClaimsHeaderKey, acceptLanguageHeaderKey}
