package delete

import (
	"context"
	"net/http"

	"github.com/cateiru/access-tracker/core/common"
	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/access-tracker/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

// Delete tracking url
func DeleteHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := net.GetQuery(r, "id")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}
	accessKey, err := net.GetQuery(r, "key")
	if err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	if err := Delete(ctx, id, accessKey); err != nil {
		return err
	}

	return nil
}

func Delete(ctx context.Context, id string, accessKey string) error {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	// アクセスキーを検証する
	if err := common.ValidateKey(ctx, db, id, accessKey); err != nil {
		return err
	}

	// 削除する
	if err := models.DeleteHistoriesByTrackID(ctx, db, id); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if err := models.DeleteTrackByTrackID(ctx, db, id); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}
