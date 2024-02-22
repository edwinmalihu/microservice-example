package middleware

import (
	"auth-services/utils"
	"os"
	"strconv"
	"time"

	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY_SESSION")))

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})

		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := utils.ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token"})

		} else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Not Valid Token"})

			} else {
				if token.Valid {

					current_time := time.Now()
					var w http.ResponseWriter = ctx.Writer
					var r *http.Request = ctx.Request

					getSession, err := store.Get(r, "session-synapsis")
					fmt.Println("session-synapsis", getSession)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					if getSession.Values["last_hit"] != nil {
						layout := "2006-01-02 15:04:05"
						str := fmt.Sprintf("%s", getSession.Values["last_hit"])
						t, err := time.Parse(layout, str)
						if err != nil {
							fmt.Println(err)
						}

						localTime := time.Now().Add(7 * time.Hour)

						diff := localTime.Sub(t)
						idle_timeout, _ := strconv.Atoi(os.Getenv("IDLE_TIMEOUT"))
						if int(diff.Minutes()) > idle_timeout {
							session, _ := store.Get(r, "session-synapsis")
							session.Values["last_hit"] = nil
							err := session.Save(r, w)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
						} else {
							session, err := store.Get(r, "session-synapsis")
							fmt.Println("session-synapsis", session)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
							session.Values["last_hit"] = current_time.Format("2006-01-02 15:04:05")
							session.Values["username"] = claims["username"]
							err = session.Save(r, w)
							if err != nil {
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
							ctx.Set("username", claims["username"])
						}
					} else {
						session, err := store.Get(r, "session-synapsis")
						fmt.Println("session-synapsis", session)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						session.Values["last_hit"] = time.Now().Format("2006-01-02 15:04:05")
						session.Values["username"] = claims["username"]
						err = session.Save(r, w)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						ctx.Set("username", claims["username"])
					}
				} else {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "Not Valid Token"})
				}

			}
		}

	}

}

func GetUserToken(ctx *gin.Context) string {
	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "No Authorization header found"})

	}
	tokenString := authHeader[len(BearerSchema):]

	token, _ := utils.ValidateToken(tokenString)

	fmt.Println("token", tokenString)

	claims, _ := token.Claims.(jwt.MapClaims)

	username := fmt.Sprintf("%v", claims["username"])

	fmt.Println("user claims:", username)

	return username

}
