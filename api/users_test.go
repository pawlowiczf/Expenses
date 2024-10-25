package api

import (
	"bytes"
	"encoding/json"
	"expenses/util"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCreateUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	arg := createUserRequest{
		Username: util.RandomUsername(),
		Email:    util.RandomEmail(),
		Password: "secret",
		FullName: util.RandomUsername(),
	}
	body, err := json.Marshal(arg)
	require.NoError(t, err)
	require.NotEmpty(t, body)

	reqBody := bytes.NewReader(body)
	req := httptest.NewRequest(http.MethodPost, "/v1/create_user", reqBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	serverTest.router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var rsp createUserResponse
	err = json.Unmarshal(w.Body.Bytes(), &rsp)
	require.NoError(t, err)
	require.NotEmpty(t, rsp)

	require.Equal(t, arg.FullName, rsp.FullName)
	require.Equal(t, arg.Email, arg.Email)
	require.Equal(t, arg.Password, arg.Password)
	require.WithinDuration(t, time.Now(), rsp.CreatedAt, time.Second)
}
