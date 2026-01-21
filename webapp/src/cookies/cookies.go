package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Config loads envioronment variables used to configure SecureCookie
func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save stores the authentications data in a secure cookie
func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	codedData, error := s.Encode("data", data)
	if error != nil {
		return error
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    codedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Read retrieves the data stored in the cookie
func Read(r *http.Request) (map[string]string, error) {
	cookie, error := r.Cookie("data")
	if error != nil {
		return nil, error
	}

	values := make(map[string]string)
	if error = s.Decode("data", cookie.Value, &values); error != nil {
		return nil, error
	}

	return values, nil
}

// Delete removes the authentications data stored in the cookie
func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
