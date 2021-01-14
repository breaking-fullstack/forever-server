package middleware

import (
	"strings"

	"github.com/breaking-fullstack/forever-server/verifier"
	"github.com/gin-gonic/gin"
)

func Auth(v verifier.Verifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		authFields := strings.Fields(authHeader)
		if len(authFields) < 2 {
			c.AbortWithStatus(401)
			return
		}

		idToken := authFields[1]
		userID, err := v.Verify(c.Request.Context(), idToken)
		if err != nil {
			c.AbortWithError(401, err)
		}

		c.Set("UID", userID)
		c.Next()
	}
}
