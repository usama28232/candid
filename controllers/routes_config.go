package controllers

import (
	"authexample/logging"
	"authexample/shared"
	"authexample/users"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var NOAUTH = []string{"/hello"}

func RegisterRoutes() *mux.Router {

	mux := mux.NewRouter()
	mux.Use(appMiddleware)
	helloCont := &HelloController{}
	userCont := &UserController{}

	mux = register(userCont, mux)
	mux = register(helloCont, mux)

	mux.StrictSlash(false)
	return mux
}

func appMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessLogger := logging.GetAccessLogger()
		startTime := time.Now()
		meta := AppRequestLog{}
		meta.Method = r.Method
		meta.Url = r.URL.Path
		meta.Agent = r.UserAgent()

		if shared.CollectionContainsOrStartsWith(NOAUTH, r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" && strings.HasPrefix(authHeader, "Basic ") {
				// Decode the base64-encoded username and password
				encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
				decoderStr := shared.DecodeBase64(encodedCredentials)

				if len(decoderStr) == 0 {
					http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
					return
				}

				cred := shared.StringSplit(decoderStr, ":")

				if len(cred) > 1 {
					authErr := users.AuthenticateUser(cred[0], cred[1])
					if authErr == nil {
						next.ServeHTTP(w, r)
					} else {
						http.Error(w, authErr.Error(), http.StatusUnauthorized)
						return
					}
				} else {
					http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
					return
				}
			} else {
				// Authentication header is missing or invalid
				http.Error(w, "unable to parse Auth Header", http.StatusUnauthorized)
			}
		}
		meta.Duration = time.Since(startTime).Milliseconds()
		accessLogger.Infow("HttpRequest", zap.Any("v", meta))

	})
}
