package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cateiru/access-tracker/core/create"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/access-tracker/routes"
	"github.com/cateiru/access-tracker/tests/tools"
	"github.com/stretchr/testify/require"
)

func server() *http.ServeMux {
	mux := http.NewServeMux()

	routes.Defs(mux)

	return mux
}

func TestCreateAndTrackingByCustomMessage(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	server := httptest.NewServer(server())

	url := server.URL

	redirect := "hogehoge"

	// --- tracking urlを作成する
	resp, err := http.Post(url+"/u", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(fmt.Sprintf("redirect=%s", redirect))))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var createData create.CreatedResponse
	err = json.Unmarshal(tools.ConvertByteResp(resp), &createData)
	require.NoError(t, err)

	require.Equal(t, createData.RedirectUrl, redirect)
	require.NotEmpty(t, createData.AccessKey)
	require.NotEmpty(t, createData.TrackId)

	trackId := createData.TrackId
	accessKey := createData.AccessKey

	// --- trackingする

	resp, err = http.Get(fmt.Sprintf("%s/%s", url, trackId))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// redirectはURLではないため200でそのまま返す
	require.Equal(t, redirect, tools.ConvertResp(resp))

	// --- アクセス履歴を取得する

	resp, err = http.Get(fmt.Sprintf("%s/u?id=%s&key=%s", url, trackId, accessKey))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var histories []models.History
	err = json.Unmarshal(tools.ConvertByteResp(resp), &histories)
	require.NoError(t, err)

	require.Len(t, histories, 1)

	// --- 削除する

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/u?id=%s&key=%s", url, trackId, accessKey), nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// --- track urlにアクセスすると削除済みなので404が帰る
	resp, err = http.Get(fmt.Sprintf("%s/%s", url, trackId))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 404)
}

func TestCreateAndTrackingByURL(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	server := httptest.NewServer(server())

	url := server.URL

	redirect := "https://example.com"

	// --- tracking urlを作成する
	resp, err := http.Post(url+"/u", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(fmt.Sprintf("redirect=%s", redirect))))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var createData create.CreatedResponse
	err = json.Unmarshal(tools.ConvertByteResp(resp), &createData)
	require.NoError(t, err)

	require.Equal(t, createData.RedirectUrl, redirect)
	require.NotEmpty(t, createData.AccessKey)
	require.NotEmpty(t, createData.TrackId)

	trackId := createData.TrackId
	accessKey := createData.AccessKey

	// --- trackingする

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	resp, err = client.Get(fmt.Sprintf("%s/%s", url, trackId))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 302)

	// --- アクセス履歴を取得する

	resp, err = http.Get(fmt.Sprintf("%s/u?id=%s&key=%s", url, trackId, accessKey))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var histories []models.History
	err = json.Unmarshal(tools.ConvertByteResp(resp), &histories)
	require.NoError(t, err)

	require.Len(t, histories, 1)

	// --- 削除する

	client = &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/u?id=%s&key=%s", url, trackId, accessKey), nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// --- track urlにアクセスすると削除済みなので404が帰る
	resp, err = http.Get(fmt.Sprintf("%s/%s", url, trackId))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 404)
}
