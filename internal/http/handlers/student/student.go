package student

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/HakashiKatake/Students-Rest-Api-Go/internal/storage"
	"github.com/HakashiKatake/Students-Rest-Api-Go/internal/types"
	"github.com/HakashiKatake/Students-Rest-Api-Go/internal/utlis/response"
	"github.com/go-playground/validator/v10"
)

type Sqlite struct {
	Db *sql.DB
}

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid body: %s", err.Error())))
			return
		}

		// Validate the student data

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		slog.Info("student created", slog.Int64("id", lastId), slog.String("name", student.Name), slog.String("email", student.Email), slog.Int("age", student.Age))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("failed to create student: %s", err.Error())))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		slog.Info("getting student by id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Info("invalid id", slog.String("id", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id: %s", err.Error())))
			return
		}

		student, err := storage.GetStudentById(intId)

		if err != nil {
			slog.Info("student not found", slog.String("id", id))
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("student not found: %s", err.Error())))
			return
		}

		response.WriteJson(w, http.StatusOK, student)

	}
}
