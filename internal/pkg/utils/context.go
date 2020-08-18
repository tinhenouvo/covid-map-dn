package utils

import (
	"github.com/labstack/echo"

	"crawlerdatacovid/internal/pkg/config"
)

// EchoContextGetCurrentUserID return current userpb id
//func EchoContextGetCurrentUserID(c echo.Context) primitive.ObjectID {
//	userID := c.Get(config.ContextKeyCurrentUserID)
//
//	// Convert to object id
//	objectID, _ := primitive.ObjectIDFromHex(userID.(string))
//	return objectID
//}

// EchoContextGetLang return language set in context
func EchoContextGetLang(c echo.Context) string {
	lang := c.Get(config.ContextKeyLanguage)
	if lang == nil {
		return "vi"
	}
	return lang.(string)
}
