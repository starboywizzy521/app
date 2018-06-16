package Controllers

import (
	"net/http"
	"github.com/starboywizzy521/app/Charts/Models"
)

func HandleJSON(w http.ResponseWriter, r *http.Request){

	w.Write(Models.SendJsonData())
}
