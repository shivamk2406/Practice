package constants

type contextKey string

const (
	JWTClaimsKey      contextKey = "JWTClaimsKey"
	AcceptLanguageKey contextKey = "AcceptLanguageKey"
	GrpcMetadataKey   contextKey = "GrpcMetadataKey"
	KafkaHeadersKey   contextKey = "KafkaHeadersKey"
	AuthKey           contextKey = "AuthKey"
)

type Subscription string

const (
	Iron     Subscription = "IRON"
	Silver   Subscription = "SILVER"
	Gold     Subscription = "GOLD"
	Platinum Subscription = "PLATINUM"
)
