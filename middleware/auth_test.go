package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/breaking-fullstack/forever-server/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func testHandler(c *gin.Context) {
	userID := c.GetString("UID")
	c.String(200, userID)
}

func TestAuth(t *testing.T) {
	r := gin.New()
	authHandler := Auth(&testhelper.AuthVerifier{})
	r.GET("/testauth", authHandler, testHandler)

	testCases := []struct {
		token string
		code  int
		uid   string
	}{
		{
			fmt.Sprintf("Bearer %s", testhelper.ValidAuthJWT),
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
