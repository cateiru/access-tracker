// 返すレスポンスを統一します。
//
// Example:
//	err := errors.New("dummy error")
//	ResponseError(w, 500, err) // internal server error
//
//	WriteBody(w, response) // 200 OK
//
package net

import (
	"encoding/json"
	"net/http"

	"github.com/cateiru/access-tracker/logging"
	"github.com/cateiru/access-tracker/utils"
	"github.com/cateiru/go-http-error/httperror"
)

// 独自ステータスコード
const (
	Success        = iota // 成功
	DefaultError          // エラー
	ErrorIntoError        // ResponseError内でのエラー
	BlockedError          // ブロックリストに入っていたエラー
	ExistError            // メールアドレスなどが既に存在しているエラー
	BotError              // Bot判定したためエラー
	TimeOutError          // 時間切れ
	AlreadyDone           // 既に認証済み
	AccountNoExist        // アカウント無い
)

type AbstractResponse struct {
	// 独自ステータスコード
	//
	// 特殊な事情でエラーが起きた場合HTTP ステータスコード以外にこのコードを指定します。
	Code int `json:"code"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	ErrorID    string `json:"error_id"`

	AbstractResponse
}

// エラーをHTTPで返す
// レスポンスではerror idを返し、ログからそのIDを検索することでエラーメッセージを参照できる
//
// http statusはHTTPErrorで定義してください。
// See more: https://github.com/cateiru/go-http-error
func ResponseError(w http.ResponseWriter, err error) {

	var statusCode int
	var customCode int
	if httperr, ok := httperror.CastHTTPError(err); ok {
		statusCode = httperr.StatusCode
		if httperr.Code != 0 {
			customCode = httperr.Code
		} else {
			customCode = 1
		}
	} else {
		statusCode = 500
		customCode = DefaultError
	}

	// 404のときはエラーはトレースしない
	if statusCode == 404 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		id := utils.CreateID(10)
		body := ErrorResponse{
			StatusCode: statusCode,
			ErrorID:    id,
			AbstractResponse: AbstractResponse{
				Code: customCode,
			},
		}
		logging.Sugar.Errorf("HTTP ERROR. id: %v, message: %v", id, err.Error())
		ResponseCustomStatus(w, statusCode, body)
	}
}

// カスタムに独自ステータスコードを決定し、エラーをHTTPで返す
// レスポンスではerror idを返し、ログからそのIDを検索することでエラーメッセージを参照できる
func ResponseErrorCustomCode(w http.ResponseWriter, statusCode int, err error, code int) {
	id := utils.CreateID(10)

	logging.Sugar.Errorf("HTTP ERROR. id: %v, message: %v, code: %v", id, err.Error(), code)

	body := ErrorResponse{
		StatusCode: statusCode,
		ErrorID:    id,
		AbstractResponse: AbstractResponse{
			Code: code,
		},
	}

	ResponseCustomStatus(w, statusCode, body)
}

// ステータスコード200で書き出す
func ResponseOK(w http.ResponseWriter, body interface{}) {
	ResponseCustomStatus(w, http.StatusOK, body)
}

// bodyをHTTP Responceに書き出す
func ResponseCustomStatus(w http.ResponseWriter, statusCode int, body interface{}) {
	bodyByte, err := json.Marshal(body)
	if err != nil {
		id := utils.CreateID(10)
		logging.Sugar.Errorf("WriteError error. id: %v error: %v", id, err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status_code": 500, "code": 2, "error_id": "` + id + `"}`))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(bodyByte)
}
