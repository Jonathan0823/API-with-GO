package middleware
import (
	"errors"
	"net/http"

	"github.com/Jonathan0823/API-with-GO/api"
	"github.com/Jonathan0823/API-with-GO/internal/tools"
	log "github.com/sirupsen/logrus"
)

var unAuthorizedError = errors.New("Unauthorized")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error("Unauthorized request")
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}

		var database = *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.internalErrorHandler(w)
			return
		}

		var LoginDetails *tools.LoginDetails
		LoginDetails = (*database).GetUserLoginDetails(username)

		if(LoginDetails == nil || (token != (*LoginDetails).AuthToken)) {
			log.Error("Unauthorized request")
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}