package mid

import (
	"context"
	"net/http"

	logger "github.com/aforamitdev/pdfsync/zero"
	"github.com/aforamitdev/pdfsync/zero/web"
)

func Error(log *logger.Logger) web.Middleware {
	m := func(before web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := before(ctx, w, r); err != nil {
				log.Errorw(ctx, "ERROR", err.Error())

				switch {
				default:
					er := web.ErrorResponse{
						Error: http.StatusText(),
					}
				}

				if err := web.Response(ctx, w, err); err != nil {
					return err
				}
				if ok := web.IsShutdown(err); ok {
					return err
				}

			}
			return nil
		}
		return h

	}
	return m
}
