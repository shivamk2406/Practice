package header

import (
	"context"
	"errors"

	"golang.org/x/text/language"
	"google.golang.org/grpc/metadata"
)

var (
	errIllegalValue = errors.New("illegal value in language map")

	// matcher is the default supported languages. The first language
	// here is treated as a default.
	matcher = language.NewMatcher([]language.Tag{
		language.English,
		language.Arabic,
	})
)

const (
	acceptLanguageHeaderKey string     = "accept-language"
	AcceptLanguageKey       contextKey = "AcceptLanguageKey"
)

func ProcessAcceptLanguage() func(ctx context.Context, md metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		tag, _ := language.MatchStrings(matcher, md[acceptLanguageHeaderKey]...)
		lang, _ := tag.Base()
		return context.WithValue(ctx, AcceptLanguageKey, lang.String())
	}
}

func ParseLanguageMap(lang language.Tag, data map[string]interface{}) (string, error) {
	if val, ok := data[lang.String()].(string); ok {
		return val, nil
	}

	return "", errIllegalValue
}
