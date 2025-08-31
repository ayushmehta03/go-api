package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/ayushmehta03/go-api/internal/storage"
	"github.com/ayushmehta03/go-api/internal/types"
	"github.com/ayushmehta03/go-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
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
	return 
}

	
	lastid,err:=storage.CreateStudent(
		student.Name,
		student.Email,
		student.Age,
	)
	slog.Info("user created succesfull",slog.String("userId",fmt.Sprint(lastid)))
	if err!=nil{
		response.WriteJson(w,http.StatusInternalServerError,err)
		return
	}






		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastid})
	}
}
