package translator

import (
	"context"
	"fmt"
	"reflect"

	"github.com/coding-challenge/api-searching/helpers/util"
	"google.golang.org/grpc/metadata"

	"github.com/coding-challenge/api-searching/config"
)

var dataTranslation map[string]map[string]string
var localeDefault = []string{"en", "my", "id", "th", "ms", "jp"}

func GetLocale(ctx context.Context) string {
	locale, ok := util.GetKeyFromContext(ctx, "locale")
	if !ok {
		headers, exist := metadata.FromIncomingContext(ctx)
		if exist {
			header, existHeader := headers["x-lang"]
			if existHeader {
				if len(header) > 0 {
					return header[0]
				}
			}
		}
		cfg := config.GetConfig()
		locale = cfg.GetString("server.locale")
		//SetLocale(ctx, locale.(string))
	}
	return locale.(string)
}

func SetLocale(ctx context.Context, locale string) context.Context {
	if !IsLocaleSupported(locale) {
		cfg := config.GetConfig()
		locale = cfg.GetString("server.locale")
	}

	ctx = context.WithValue(ctx, "locale", locale)

	return ctx
}

func Trans(ctx context.Context, key string) string {
	locate, ok := util.GetKeyFromContext(ctx, "locale")
	if !ok {
		fmt.Println("can not get locate")
		return key
	}
	dataLocateTranslation, ok := dataTranslation[locate.(string)]
	if !ok {
		fmt.Println("can not get locate array")
		return key
	}
	val, ok := dataLocateTranslation[key]
	if !ok {
		return key
	}
	return val
}

func IsLocaleSupported(locale interface{}) bool {
	switch reflect.TypeOf(localeDefault).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(localeDefault)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(locale, s.Index(i).Interface()) == true {
				return true
			}
		}
	}

	return false
}
