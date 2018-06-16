package HistoriqueApprovisionnement

import (
	"net/http"
	"html/template"
	"github.com/starboywizzy521/app/HistoriqueApprovisionnement/Models"
	"github.com/starboywizzy521/app/Public/Model_Public"
	"github.com/starboywizzy521/app/Login"
)

func HandleHistoAppro(w http.ResponseWriter,r *http.Request){
	Model_Public.Test_session(Login.Session)
	ListeProduits,total,Ope:=Models.HistoApr()
	PageData:=Models.Mystruct{ListeProduits,total,Ope}
	t:=template.Must(template.ParseFiles("HistoriqueApprovisionnement/Vue/HistoAppro.html","webPages/include-header.html","HistoriqueApprovisionnement/Vue/HistoApproDataTables.html", "webPages/include-body.html", "webPages/include-footer.html"))
	t.Execute(w,PageData)
}