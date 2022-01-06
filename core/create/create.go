package create

import (
	"context"
	"net/http"
	"time"

	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/access-tracker/utils"
	"github.com/cateiru/access-tracker/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type CreatedResponse struct {
	TrackId     string `json:"track_id"`
	AccessKey   string `json:"access_key"`
	RedirectUrl string `json:"redirect_url"`
}

func CreateHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	r.ParseForm()
	redirectUrl := r.PostFormValue("redirect")

	resp, err := Create(ctx, redirectUrl)
	if err != nil {
		return err
	}

	net.ResponseOK(w, resp)
	return nil
}

// Tracking urlを作成します
func Create(ctx context.Context, redirectUrl string) (*CreatedResponse, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	id := utils.CreateID(5)
	accessKey := utils.CreateID(0)

	entity := models.Track{
		TrackId:     id,
		AccessKey:   accessKey,
		RedirectUrl: redirectUrl,
		Create:      time.Now(),
	}

	if err := entity.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return &CreatedResponse{
		TrackId:     id,
		AccessKey:   accessKey,
		RedirectUrl: redirectUrl,
	}, nil
}
