package whois

import (
	"context"
	"net/http"

	"github.com/cateiru/access-tracker/core/commom"
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

	history, err := WhoIs(ctx, id, accessKey)
	if err != nil {
		return err
	}

	net.ResponseOK(w, history)

	return nil
}

func WhoIs(ctx context.Context, id string, accessKey string) ([]models.History, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// アクセスキーを検証する
	if err := commom.ValidateKey(ctx, db, id, accessKey); err != nil {
		return nil, status.NewBadRequestError(err).Caller()
	}

	histories, err := models.GetHistoriesByTrackID(ctx, db, id)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return histories, nil
}
