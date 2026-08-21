package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/belobragin/bvv_utils/goutils/token"
	"github.com/belobragin/bvv_utils/goutils/util"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

type AuthError struct {
	error
	Message string
}

type auth interface {
	OutPubKey() (any, error)
	GetLog() *zap.Logger
	GetToken() token.TokenDataI
}

func AuthenticateRsaBearerMiddleware(
	u auth,
	optionalRoutes []string,
	authCookieName, sessionCookieName,
	claimSessionCookieKey string,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				e                         error
				err                       mistake.OuterrorI
				ac, sc, header            string
				headerParts               []string
				authCookie, sessionCookie *http.Cookie
				ctx                       = r.Context()
				route                     = mux.CurrentRoute(r)
				optional                  = false
				zapstruct                 = u.GetLog()
				claims                    = u.GetToken()
				cookie                    = new(string)
				pubKey                    any
			)
			for _, name := range optionalRoutes {
				if name == route.GetName() {
					zapstruct.Info("Public route, ", zap.String("name: ", name))
					optional = true
					break
				}
			}
			// read cookie:
			authCookie, e = r.Cookie(authCookieName)
			switch e {
			case nil:
				ac = authCookie.Value
				if ac == "" {
					err = mistake.NewOutErr(
						mistake.ErrEmptyAuthCookie, http.StatusUnauthorized)
					goto mistakehttp

				}
			case http.ErrNoCookie:
				// intentionally no operand there!
			default:
				err = mistake.NewOutErr(e, http.StatusUnauthorized)
				goto mistakehttp
			}
			sessionCookie, e = r.Cookie(sessionCookieName)
			switch e {
			case nil:
				sc = sessionCookie.Value
				if sc == "" {
					err = mistake.NewOutErr(
						mistake.ErrEmptySessionCookie, http.StatusUnauthorized)
					goto mistakehttp
				}
				ctx = context.WithValue(ctx, claimSessionCookieKey, sc)
			case http.ErrNoCookie:
			default:
				err = mistake.NewOutErr(e, http.StatusUnauthorized)
				goto mistakehttp
			}
			// read headers:
			header = r.Header.Get("Authorization")
			if header == "" || header == "Bearer" {
				if optional {
					goto nexthttp
				} else {
					err = mistake.NewOutErr(mistake.ErrNoAuthHeader, http.StatusUnauthorized)
					goto mistakehttp
				}
			}
			headerParts = strings.Split(header, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				err = mistake.NewOutErr(
					mistake.ErrInvalidJAuthHeader, http.StatusUnauthorized)
				goto mistakehttp
			}
			pubKey, e = u.OutPubKey()
			if e != nil {
				err = mistake.NewOutErr(e, http.StatusUnauthorized)
				goto mistakehttp
			}
			e = claims.ParseValidateRsaToken(headerParts[1], pubKey)
			//OTLADKA:
			// fmt.Printf("%+v\n", claims)
			switch e {
			case nil:
				c, ok := claims.(token.AutoRenewTokenDataI)
				if !ok {
					err = mistake.NewOutErr(
						fmt.Errorf("no renew data"), http.StatusUnauthorized)
					goto mistakehttp
				}
				if cookie = c.GetSessionID(); cookie == nil {
					err = mistake.NewOutErr(
						mistake.ErrAutorenewSessionIDNull, http.StatusUnauthorized)
					goto mistakehttp
				}
				if !util.CompareCookieWithHash(ac, *cookie) {
					err = mistake.NewOutErr(
						mistake.ErrCookieSignatureInvalid, http.StatusUnauthorized)
					goto mistakehttp
				}
				// put token data to context:
				ctx = context.WithValue(ctx, token.JwtOldClaimKey, claims)
				// do NOT check and record token auth fields:
				goto nexthttp
			case jwt.ErrTokenExpired:
				err = mistake.NewOutErr(e, http.StatusRequestTimeout)
				goto mistakehttp
			default:
				err = mistake.NewOutErr(e, http.StatusUnauthorized)
				goto mistakehttp
			}
			// fmt.Printf("%+v\n", claims)
		nexthttp:
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		mistakehttp:
			zapstruct.Error(err.Err().Error())
			StandardRespond(zapstruct, w, r, err.ErrCode(), err.Err().Error())
		})
	}
}
