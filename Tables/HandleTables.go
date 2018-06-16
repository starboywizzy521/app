package Tables

import (
	"net/http"
	"html/template"
	"github.com/starboywizzy521/app/Public/Model_Public"
	"github.com/starboywizzy521/app/Login"
	"github.com/starboywizzy521/app/Tables/Model"
)

func HandleTables(w http.ResponseWriter,r *http.Request){
	Model_Public.Test_session(Login.Session)
	PageData:=Model.Extras{Model.RecupProduits()}
	t:=template.Must(template.ParseFiles("Tables/Vue/tables.html","Tables/Vue/gg.html","webPages/include-header.html","webPages/include-body.html", "webPages/include-footer.html"))
	t.Execute(w,PageData)
}

