package mid

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aforamitdev/pdfsync/internal/sys/validate"
	logger "github.com/aforamitdev/pdfsync/zero"
	"github.com/aforamitdev/pdfsync/zero/web"
)

func Error(log *logger.Logger) web.Middleware {
	m := func(before web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := before(ctx, w, r); err != nil {

				log.Error(ctx, "error prossing")

				var er web.ErrorResponse
				var status int

				switch {
				case validate.IsFieldErrors(err):
					fieldErrors := validate.GetFieldErrors(err)

					er = web.ErrorResponse{
						Error:  "Request not value,or developer feck*ed up contact @amitraidev github",
						Fields: fieldErrors.Fields(),
					}
					status = http.StatusBadRequest

				case web.IsRequestError(err):
					reqErr := web.GetRequestError(err)
					er = web.ErrorResponse{
						Error: reqErr.Error(),
					}
					status = reqErr.Status

				default:
					fmt.Println("DEFAULT")
					er = web.ErrorResponse{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				if err := web.Response(ctx, w, er, status); err != nil {
					return err
				}

				// If we receive the shutdown err we need to return it
				// back to the base handler to shut down the service.

			}
			return nil
		}
		return h

	}
	return m
}
