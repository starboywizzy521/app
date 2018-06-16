package Controllers

import (
	"net/http"
	"github.com/starboywizzy521/app/Charts/Models"
)

func HandleEphrauJSON(w http.ResponseWriter, r *http.Request){

	w.Write(Models.SendJsonEphrau())
}