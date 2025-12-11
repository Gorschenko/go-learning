package stats

import (
	"fmt"
	"net/http"
	"test/configs"
	"test/packages/response"
	"time"
)

type StatsHandlerDeps struct {
	StatsRepository *StatsRepository
	Config          *configs.Config
}

type StatsHandler struct {
	StatsRepository *StatsRepository
}

func NewStatsHandler(router *http.ServeMux, deps StatsHandlerDeps) {
	handler := &StatsHandler{
		StatsRepository: deps.StatsRepository,
	}

	router.Handle("GET /stats", handler.Get())
}

func (handler *StatsHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fromString := r.URL.Query().Get("from")

		fromDate, err := time.Parse("2006-01-02", fromString)

		if err != nil {
			http.Error(w, "Invalid from", http.StatusBadRequest)
			return
		}

		toString := r.URL.Query().Get("to")

		toDate, err := time.Parse("2006-01-02", toString)

		if err != nil {
			http.Error(w, "Invalid to", http.StatusBadRequest)
			return
		}

		by := r.URL.Query().Get("by")

		isValidBy := by == "day" || by == "month"

		if !isValidBy {
			http.Error(w, "Invalid by", http.StatusBadRequest)
			return
		}

		fmt.Println(fromDate, toDate, by)
		stats := handler.StatsRepository.GetStats(fromDate, toDate, by)

		response.Json(w, stats, http.StatusOK)
	}
}
