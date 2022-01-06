package whois

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cateiru/access-tracker/core/common"
	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/access-tracker/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

// Reference access history.
func WhoisHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := net.GetQuery(r, "id")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}
	accessKey, err := net.GetQuery(r, "key")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	// limitを指定する
	// もし、指定がない場合はすべての要素が返る
	limitStr := r.URL.Query().Get("limit")
	limit := -1
	if len(limitStr) != 0 {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller()
		}
	}

	history, err := WhoIs(ctx, id, accessKey, limit)
	if err != nil {
		return err
	}

	net.ResponseOK(w, history)

	return nil
}

func WhoIs(ctx context.Context, id string, accessKey string, limit int) ([]models.History, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// アクセスキーを検証する
	if err := common.ValidateKey(ctx, db, id, accessKey); err != nil {
		return nil, status.NewBadRequestError(err).Caller()
	}

	histories, err := models.GetHistoriesByTrackID(ctx, db, id, limit)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return histories, nil
}
