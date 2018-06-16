package controller

import (
	"net/http"
	"html/template"
	"github.com/starboywizzy521/app/AutoComplete/Model"
)

func HandleAutoComplete(w http.ResponseWriter,r *http.Request){

	r.ParseForm()

	if r.Method=="GET"{
		t, _ := template.ParseFiles("AutoComplete/Vue/appro.html","webPages/include-header.html", "webPages/include-body.html", "webPages/include-footer.html")
		t.Execute(w,Model.Autocomplete())
	} else {

	}

}