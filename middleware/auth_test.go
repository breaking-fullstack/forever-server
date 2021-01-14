package middleware

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const authTestJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.4TFMMEj2Ejof4za_5H_CVsM2PQX3YDYIMOZ4t0LcWjA"

type testVerifier struct {
	validToken string
}

func (tv *testVerifier) Verify(ctx context.Context, idToken string) (string, error) {
	if idToken != tv.validToken {
		return "", errors.New("invalid token")
	}
	return "testUser", nil
}

func testHandler(c *gin.Context) {
	userID := c.GetString("UID")
	c.String(200, userID)
}

func TestAuth(t *testing.T) {
	r := gin.New()
	authHandler := Auth(&testVerifier{authTestJWT})
	r.GET("/testauth", authHandler, testHandler)

	testCases := []struct {
		token string
		code  int
		uid   string
	}{
		{
			fmt.Sprintf("Bearer %s", authTestJWT),
			200,
			"testUser",
		},
		{
			"",
			401,
			"",
		},
		{
			"Bearer notavalidjwt",
			401,
			"",
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			testReq := httptest.NewRequest(http.MethodGet, "/testauth", nil)
			testReq.Header.Set("Authorization", tc.token)
			testRec := httptest.NewRecorder()

			r.ServeHTTP(testRec, testReq)
			assert.Equal(t, tc.code, testRec.Code)

			uid, err := ioutil.ReadAll(testRec.Body)
			assert.NoError(t, err)
			assert.Equal(t, tc.uid, string(uid))
		})
	}
}
