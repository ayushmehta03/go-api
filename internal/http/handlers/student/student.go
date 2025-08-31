package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/ayushmehta03/go-api/internal/types"
	"github.com/ayushmehta03/go-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
				slog.Info("creating student", slog.Any("student", student))


		err := json.NewDecoder(r.Body).Decode(&student)
	
			if errors.Is(err, io.EOF) {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
				return
			
			
		}

		if err!=nil{
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err));
			return 

		}



		// validate

if	err:=validator.New().Struct(student);err!=nil{
	validateErrs:=err.(validator.ValidationErrors)
	response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrs))
}







		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
