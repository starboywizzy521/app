package ControllerSortieInst

import (
	"net/http"
	"html/template"
	"github.com/starboywizzy521/app/Public/Model_Public"
	"github.com/starboywizzy521/app/Login"
	"github.com/starboywizzy521/app/SortieInstant/Model"
)

func HandleSortieInst(w http.ResponseWriter,r *http.Request){
	Model_Public.Test_session(Login.Session)
	PageData:=Model.Extras{Model.RecupProduitsSortie()}
	t:=template.Must(template.ParseFiles("SortieInstant/Vue/SortieInst.html","SortieInstant/Vue/gg.html","webPages/include-header.html","webPages/include-body.html", "webPages/include-footer.html"))
	t.Execute(w,PageData)
}

