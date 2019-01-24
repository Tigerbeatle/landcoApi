package middleware

import (
	"net/http"
	"github.com/tigerbeatle/landcoApi/models"
	"encoding/json"
	"reflect"
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	u "github.com/tigerbeatle/landcoApi/utilities"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func RecoverHandler(next http.Handler) http.Handler {
	//Recover is a built-in function that regains control of a panicking goroutine.
	//Recover is only useful inside deferred functions. During normal execution,
	//a call to recover will return nil and have no other effect. If the current
	//goroutine is panicking, a call to recover will capture the value given to
	//panic and resume normal execution.
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				models.WriteError(w, models.ErrInternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func AcceptHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("r.URL.Path = ", r.URL.Path)
		fmt.Println("r.Header Accept = ", r.Header.Get("Accept"))
		fmt.Println("r.Header Authorization = ", r.Header.Get("Authorization"))
		if r.Header.Get("Accept") != "application/json" {
			models.WriteError(w, models.ErrNotAcceptable)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}


func BodyHandler(v interface{}) func(http.Handler) http.Handler {
	t := reflect.TypeOf(v)
	m := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(t).Interface()
			err := json.NewDecoder(r.Body).Decode(val)
			fmt.Println("Middleware BodyHandler - val:",val)
			if err != nil {
				models.WriteError(w, models.ErrBadRequest)
				return
			}

			if next != nil {
				context.Set(r, "body", val)

				params := context.Get(r, "params").(httprouter.Params)
				fmt.Println("middle params:",params)

				next.ServeHTTP(w, r)
			}
		}
		return http.HandlerFunc(fn)
	}
	return m
}

func ContentTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.println("Content-Type:",r.Header.Get("Content-Type"))
		if r.Header.Get("Content-Type") != "application/json" {
			models.WriteError(w, models.ErrUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

/*
func AuthorizationHandler_old(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Token")
		fmt.Println("=======jwtToken:", jwtToken)

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(models.GetSecret()), nil
		})



		fmt.Println("err:",err)
		if err == nil && token.Valid {
		}else{
			if err.Error() == "token is expired"{
				// try to get a new token using the UUID inside this expired token
				var user models.User






				user.UUID = token.Claims["id"].(string)
				b, err := json.Marshal(user)
				if err != nil {
					fmt.Println(err)
					return
				}

				// send message to api to create user profile
				url := "http://127.0.0.1:8003/auth/1.0/accounts/generateJWT"
				client := &http.Client{Timeout: 10 * time.Second}
				req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
				req.Header.Set("Accept", "application/vnd.api+json")
				req.Header.Set("content-type", "application/vnd.api+json")
				res, _ := client.Do(req)
				defer req.Body.Close()

				fmt.Println("res.StatusCode:",res.StatusCode)
				fmt.Println("res.Status:",res.Status)

				if res.StatusCode != 201 {
					if res.StatusCode == 805 {
						models.WriteError(w, models.ErrAccountDisabled)
						return
					}
					if res.StatusCode == 806 {
						models.WriteError(w, models.ErrLexpExpired)
						return
					} else {
						models.WriteError(w, models.ErrInternalServer)
						return
					}
				}

				//  Unpack body which is a json object to get jwtString
				bodyBytes, err := ioutil.ReadAll(res.Body)
				source := (*json.RawMessage)(&bodyBytes)
				var target models.BasicJSONReturn
				err = json.Unmarshal(*source, &target)
				if err != nil {
					// todo log this panic
					panic(err)
				}
				// replace the token in the r header with the new token
				r.Header.Set("Token", target.Payload)
				r.Header.Set("Token-renewed", "true")
			} else {
				models.WriteError(w, models.ErrUserTokenRejected)
				return
			}
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
*/

func AuthorizationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
fmt.Println("tokenHeader:",tokenHeader)
		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		fmt.Println("--Bearer:",splitted[0])
		fmt.Println("{--token-body}",splitted[1])
		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		// parse the token-body using the Secret from the config.json
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(models.GetSecret()), nil
		})


		if err != nil { //Malformed token, returns with http code 403 as usual
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}



		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("UserId %", tk.UserId) //Useful for monitoring
		//ctx := context.WithValue(r.Context(), "user", tk.UserId)


		context.Set(r, "user", tk.UserId)


		//r = r.WithContext(ctx)

		next.ServeHTTP(w, r) //proceed in the middleware chain!
	},
	)

}