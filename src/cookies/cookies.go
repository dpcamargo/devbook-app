package cookies

import (
	"app/src/config"
	"net/http"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	dadosCodificados, err := s.Encode("dados", dados)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
