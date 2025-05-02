package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/HakashiKatake/Students-Rest-Api-Go/internal/types"
	"github.com/HakashiKatake/Students-Rest-Api-Go/internal/utlis/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		slog.Info("creating new student")

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "student created"})
	}
}
