package middleware

import (
	"context"
	"errors"
	"net/http"
	"slices"
	"strings"

	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/belobragin/bvv_utils/goutils/token"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

const (
	authCookieName    = "bvv_auth_cookie"
	sessionCookieName = "bvv_session_cookie"
)

type AuthError struct {
	error
	Message string
}

/*
this middleware does NOT distinguish among user, f_user & partner
and codes as user_uuid token data from any field: user_uuid, f_user_uuid & partner_uuid
*/

func AuthenticateRsaBearerMiddleware(
	zapstruct *zap.Logger,
	optionalRoutes []string,
	serviceName string,
	parseKeyFunc func(interface{}, interface{}) error,
	pubKey interface{},
	fail func(w http.ResponseWriter, r *http.Request, e error),
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				ctx                         = r.Context()
				claims                      = new(token.TokenData)
				ac, sc, intendedServiceName string
			)
			route := mux.CurrentRoute(r)
			optional := false
			for _, name := range optionalRoutes {
				if name == route.GetName() {
					zapstruct.Info("Public route, ", zap.String("name: ", name))
					optional = true
					break
				}
			}
			// read cookie:
			authCookie, err := r.Cookie(authCookieName)
			switch err {
			case nil:
				ac = authCookie.Value
				if ac == "" {
					zapstruct.Error(
						"Invalid auth cookie - null value: ", zap.Error(errors.New("")),
					)
					fail(w, r, AuthError{
						error:   err,
						Message: "Invalid auth cookie: null value",
					})
					return
				}
				// ctx = context.WithValue(ctx, pdata.JwtClaimAuthCookieKey, ac)
			case http.ErrNoCookie:
			default:
				zapstruct.Error(
					"Invalid auth cookie - null value: ", zap.Error(err),
				)
				fail(w, r, AuthError{
					error:   err,
					Message: "Invalid auth cookie: no value",
				})
				return
			}
			sessionCookie, err := r.Cookie(sessionCookieName)
			switch err {
			case nil:
				sc = sessionCookie.Value
				if sc == "" {
					zapstruct.Error(
						"Invalid session cookie - null value: ", zap.Error(errors.New("")),
					)
					fail(w, r, AuthError{
						error:   err,
						Message: "Invalid session cookie: null value",
					})
					return
				}
				ctx = context.WithValue(ctx, "auth_cookie", sc)
			case http.ErrNoCookie:
			default:
				zapstruct.Error(
					"Invalid auth cookie - null value: ", zap.Error(err),
				)
				fail(w, r, AuthError{
					error:   err,
					Message: "Invalid session cookie: no value",
				})
				return
			}
			// read headers:
			header := r.Header.Get("Authorization")
			if header == "" || header == "Bearer" {
				if optional {
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				} else {
					err = errors.New("no Authorization header")
					zapstruct.Error(
						"JWT Token not valid: ", zap.Error(err),
					)
					fail(w, r, AuthError{
						error:   err,
						Message: "JWT Token not valid",
					})
					return
				}
			}
			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				err = errors.New("invalid Authorization header")
				zapstruct.Error(
					"JWT Token not valid: ", zap.Error(err),
				)
				fail(w, r, AuthError{
					error:   err,
					Message: "JWT Token not valid",
				})
				return
			}
			var inp interface{} = headerParts[1]
			err = parseKeyFunc(inp, pubKey)
			// err = claims.ParseValidateRsaToken(headerParts[1], pubKey)
			//OTLADKA:
			// fmt.Printf("%+v\n", claims)
			switch err {
			case nil:
			// case jwt.ErrTokenExpired:
			// 	c := claims.ContextID
			// 	if !util.CompareCookieWithHash(ac, c) {
			// 		err = errors.New("Invalid JWT Token context_id")
			// 		zapstruct.Error(
			// 			"JWT Token not valid: ", zap.Error(err),
			// 		)
			// 		fail(w, r, AuthError{
			// 			error:   err,
			// 			Message: "JWT Token not valid",
			// 		})
			// 		return
			// 	}
			// 	// put old token data to context:
			// 	ctx = context.WithValue(ctx, "old_token", claims)
			// 	// do NOT check and record token auth fields:
			// 	goto nexthttp
			default:
				zapstruct.Error(
					"JWT Token not valid: ", zap.Error(err),
				)
				fail(w, r, AuthError{
					error:   err,
					Message: "JWT Token not valid",
				})
				return
			}
			intendedServiceName = strings.Split(mux.CurrentRoute(r).GetName(), "_")[0]
			// if !slices.Contains(claims.RegisteredClaims.Audience, intendedServiceName) &&
			// 	intendedServiceName != pdata.RenewObligatoryName {
			if !slices.Contains(claims.RegisteredClaims.Audience, intendedServiceName) {
				err = mistake.ErrServiceUnIntended
				fail(w, r, AuthError{
					error:   err,
					Message: "Invalid JWT Token",
				})
				return
			}
			// fmt.Printf("%+v\n", claims)
			// if claims.UserUUID != "" {
			// 	ctx = context.WithValue(ctx, "userUUID", claims.UserUUID)
			// } else {
			// 	err = mistake.ErrInvalidToken
			// 	fail(w, r, AuthError{
			// 		error:   err,
			// 		Message: "Invalid JWT Token",
			// 	})
			// 	return
			// }
			ctx = context.WithValue(ctx, "roleID", claims.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
