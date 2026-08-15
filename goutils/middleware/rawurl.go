package middleware

// import (
// 	"net/http"

// 	"github.com/belobragin/bvv_utils/goutils/util"
// 	"github.com/gorilla/mux"
// 	"go.uber.org/zap"
// )

// func ProcessUrlMiddleware(
// 	zapstruct *zap.Logger, publicApiName []string,
// ) func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			ctx := r.Context()
// 			values := r.URL.Query()
// 			// check, if the route does not demand authentication
// 			route := mux.CurrentRoute(r)

// 			optional := false
// 			for _, name := range publicApiName {
// 				if name == route.GetName() {
// 					optional = true
// 					break
// 				}
// 			}
// 			if optional {
// 				next.ServeHTTP(w, r.WithContext(ctx))
// 				return
// 			}
// 			ctx, err := util.ParseUrlToCtx(ctx, values)
// 			if err != nil {
// 				// TODO: add metrica
// 				StandardRespond(zapstruct, w, r, http.StatusForbidden, err)
// 				return
// 			}
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		})
// 	}
// }
