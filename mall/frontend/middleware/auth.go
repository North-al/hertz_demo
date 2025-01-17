package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

const (
	UserIDKey = "user_id"
)

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		session := sessions.Default(c)
		userID := session.Get("user_id")
		ctx = context.WithValue(ctx, UserIDKey, userID)
		c.Next(ctx)
	}
}

func Auth(ctx context.Context, c *app.RequestContext) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	next := c.FullPath()

	if userID == nil {
		c.Redirect(consts.StatusFound, []byte("/sign-in?next="+next))
		c.Abort()
		return
	}

	c.Next(ctx)
}
