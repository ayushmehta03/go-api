package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/ayushmehta03/go-api/internal/types"
	"github.com/ayushmehta03/go-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
				slog.Info("creating student", slog.Any("student", student))


		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			if errors.Is(err, io.EOF) {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
				return
			}
			response.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}

		// validate




		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
