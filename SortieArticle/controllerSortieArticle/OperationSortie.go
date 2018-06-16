package controllerSortieArticle

import (
	"net/http"
	"github.com/starboywizzy521/app/SortieArticle/Models"
)

func OperationSortie(w http.ResponseWriter,r *http.Request){

	benef:= r.FormValue("benef")
	t:= r.FormValue("tableau")
	Models.ModelOperationSortie(t,benef)
}
