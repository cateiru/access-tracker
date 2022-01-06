package track

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/cateiru/access-tracker/core/common"
	"github.com/cateiru/access-tracker/database"
	"github.com/cateiru/access-tracker/models"
	"github.com/cateiru/access-tracker/utils"
	"github.com/cateiru/access-tracker/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

// Tracking and redirect
func TrackHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	// get url path.
	// Example: http://example.com/hoge -> hoge
	path := strings.FieldsFunc(r.URL.Path, func(r rune) bool {
		return r == '/'
	})

	if len(path) != 1 || len(path[0]) == 0 {
		http.Redirect(w, r, "https://cateiru.com", http.StatusFound)
		return nil
	}
	id := path[0]
	ip := net.GetIPAddress(r)
	userAgent := net.GetUserAgent(r)

	redirect, err := Tracking(ctx, id, ip, userAgent)
	if err != nil {
		return err
	}

	if utils.IsUrl(redirect) {
		http.Redirect(w, r, redirect, http.StatusFound)
	} else {
		net.ResponseOK(w, redirect)
	}

	w.Header().Set("Cache-Control", "no-store")

	return nil
}

func Tracking(ctx context.Context, id string, ip string, userAgent string) (string, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	track, err := models.GetTrackByTrackID(ctx, db, id)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	// trackIDが定義されていない場合は404を返す
	if track == nil {
		return "", status.NewNotFoundError(errors.New("")).Caller()
	}

	analyzedUserAgent, err := common.UserAgentToJson(userAgent)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	history := models.History{
		TrackId:   id,
		Ip:        ip,
		UserAgent: string(analyzedUserAgent),
		UniqueId:  utils.CreateID(0),
		Date:      time.Now(),
	}

	if err := history.Add(ctx, db); err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	return track.RedirectUrl, nil
}
