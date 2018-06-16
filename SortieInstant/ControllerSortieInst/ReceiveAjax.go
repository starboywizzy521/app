package ControllerSortieInst

import (
	"net/http"
	"strconv"
	"github.com/starboywizzy521/app/Public/Model_Public"
	"github.com/starboywizzy521/app/Public/Structures_Public"
	"github.com/starboywizzy521/app/SortieInstant/Model"
	"fmt"
)

var prod Structures_Public.Produits

func ReceiveAjax(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if r.Method == "POST"{
	fmt.Println("coucou")
		code,err:=strconv.ParseInt(r.FormValue("code"),10,0)
		Model_Public.CheckErr(err)
		qte,err:=strconv.ParseInt(r.FormValue("quantite"),10,0)
		Model_Public.CheckErr(err)
		benef:=r.FormValue("benef")
		Model_Public.CheckErr(err)
		idOP,err:=strconv.Atoi(r.FormValue("idOP"))
		Model_Public.CheckErr(err)
		description:=r.FormValue("description")
		fmt.Println(code,qte,benef,idOP,description)
		Model.OperatRetour (idOP, code, benef, qte, description)
	}
}
