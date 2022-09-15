package delivery

import (
	"encoding/json"
	"net/http"
)

func ReadJson(r *http.Request, dest interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dest); err != nil {
		return err
	}

	return nil
}
