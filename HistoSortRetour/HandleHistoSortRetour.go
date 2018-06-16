package HistoSortRetour

import (
	"net/http"
	"html/template"
	"github.com/starboywizzy521/app/Public/Model_Public"
	"github.com/starboywizzy521/app/Login"
	"github.com/starboywizzy521/app/HistoSortRetour/Model"
)

func HandleSortRetour(w http.ResponseWriter,r *http.Request){
	Model_Public.Test_session(Login.Session)
	PageData:=Model.Extras{Model.RecupProduitsSortRet()}
	t:=template.Must(template.ParseFiles("HistoSortRetour/Vue/HistoSortRetour.html","HistoSortRetour/Vue/gg.html","webPages/include-header.html","webPages/include-body.html", "webPages/include-footer.html"))
	t.Execute(w,PageData)
}

