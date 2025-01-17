package utils

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WrapResponse(ctx context.Context, c *app.RequestContext, data map[string]interface{}) map[string]interface{} {
	session := sessions.Default(c)
	fmt.Println("session", session.Get("user_id"))
	userID := session.Get("user_id")
	if userID != nil {
		data["user_id"] = userID
	}
	return data
}
