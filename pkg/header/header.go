package header

const (
	Authorization      string = "authorization"
	AcceptLanguage     string = "accept-language"
	ContentType        string = "content-type"
	XTenantKey         string = "x-tenant-key"
	XJWTClaims         string = "x-jwt-claims"
	XRequestID         string = "x-request-id"
	XForwardedFor      string = "x-forwarded-for"
	XRealIP            string = "x-real-ip"
	XB3TraceID         string = "x-b3-traceid"
	XB3SpanID          string = "x-b3-spanid"
	XB3ParentSpanID    string = "x-b3-parentspanid"
	XB3Sampled         string = "x-b3-sampled"
	XCustomerUserAgent string = "x-customer-user-agent"
)

var RequiredHeaders = []string{
	XJWTClaims,
	AcceptLanguage,
	XTenantKey,
	XRequestID,
	XForwardedFor,
	XRealIP,
	XB3TraceID,
	XB3SpanID,
	XB3ParentSpanID,
	XB3Sampled,
	XCustomerUserAgent,
}
