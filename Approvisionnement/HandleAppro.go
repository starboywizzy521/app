package Approvisionnement

import (
	"net/http"
	"html/template"
	"github.com/starboywizzy521/app/Public/Model_Public"
	"github.com/starboywizzy521/app/Login"
)

func HandleAppro(w http.ResponseWriter,r *http.Request){
	Model_Public.Test_session(Login.Session)
	t:=template.Must(template.ParseFiles("Approvisionnement/approvisionnement.html","webPages/include-header.html", "webPages/include-body.html", "webPages/include-footer.html"))
	t.Execute(w,nil)
}
