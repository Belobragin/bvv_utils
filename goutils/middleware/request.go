package middleware

import (
	"net/http"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

// this middleware is obsolete (not used anymore anywhere).
func StandardRequestMiddleware(zapstruct *zap.Logger, publicRoutes []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			route := mux.CurrentRoute(r)
			public := false
			for _, name := range publicRoutes {
				if name == route.GetName() {
					public = true
					break
				}
			}
			if public {
				next.ServeHTTP(w, r)
				return
			}
			// jwt payload access validation:
			// fmt.Println("thats me", ctx.Value(pdata."roleID"))
			roleID, ok := ctx.Value("roleID").(string)
			// // OTLADKA:
			// fmt.Println(roleID)
			if !ok {
				// TODO: add metrica
				StandardRespond(zapstruct, w, r, http.StatusForbidden, "The token must contain role ID")
				return
			}
			if roleID == "" {
				// TODO: add metrica
				StandardRespond(zapstruct, w, r, http.StatusForbidden, "app ID must not be empty")
				return
			}
			UserUUID, ok := ctx.Value("userUUID").(string)
			if !ok {
				// TODO: add metrica
				StandardRespond(zapstruct, w, r, http.StatusForbidden, "The token must contain user UUID")
				return
			}
			if UserUUID == "" {
				// TODO: add metrica
				StandardRespond(zapstruct, w, r, http.StatusForbidden, "user ID must not be empty")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
