package httpApi

import (
	"az-appservice/log"
	"go.uber.org/zap"
	"net/http"
)

func GetPingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		workloads := q["workload"]
		var workload string
		if len(workloads) == 0 {
			log.Log.Debug("Ping without workload")
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			workload = workloads[0]
			log.Log.Debug("Ping", zap.String("workload", workload))
			_, err := w.Write([]byte("Pinged with " + workload + "\n"))
			if err != nil {
				log.Log.Error("Ping error", zap.Error(err))
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}
