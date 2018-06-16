package main

import (
	"net/http"
	"log"
	"github.com/starboywizzy521/app/Charts/Controllers"
	"github.com/starboywizzy521/app/ModifierProduit/Controller"
	"github.com/starboywizzy521/app/Login"
	"github.com/starboywizzy521/app/Tables"
	"github.com/starboywizzy521/app/Dashboard"
	"github.com/starboywizzy521/app/Approvisionnement"
	"github.com/starboywizzy521/app/HistoriqueApprovisionnement"
	"github.com/starboywizzy521/app/AutoComplete/controller"
	"github.com/starboywizzy521/app/CreateProduct/Control"
	"github.com/starboywizzy521/app/Sortie"
	"github.com/starboywizzy521/app/SortieInstant/ControllerSortieInst"
	"github.com/starboywizzy521/app/SortieArticle/controllerSortieArticle"
	"github.com/starboywizzy521/app/HistoSortRetour"
	"github.com/starboywizzy521/app/Historique"
)



func main() {

	http.HandleFunc("/createcategory",Control.CreateCategoryHandle)

	http.HandleFunc("/createproduct",Control.CreateProductControl)

	http.HandleFunc("/autocomplete",controller.HandleAutoComplete)

	http.HandleFunc("/OperationApprovisionnement",controller.OperationApprovisionnement)

	http.HandleFunc("/ephraujson",Controllers.HandleEphrauJSON)

	http.HandleFunc("/json",Controllers.HandleJSON)

	http.HandleFunc("/misajour", Controller.HandleMisAjour)

	http.HandleFunc("/Détailsduproduit", Controller.HandleModif)

	http.HandleFunc("/receive", Controller.ReceiveAjax)

	http.HandleFunc("/login", Login.HandleLogin)

	http.HandleFunc("/ChangeLogin", Login.HandleChangeLogin)

	http.HandleFunc("/tables", Tables.HandleTables)

	http.HandleFunc("/charts", Controllers.HandleCharts)

	http.HandleFunc("/index", Dashboard.HandleIndex)

	http.HandleFunc("/appro", Approvisionnement.HandleAppro)

	http.HandleFunc("/HistoAppro",HistoriqueApprovisionnement.HandleHistoAppro)

	http.HandleFunc("/sortie", Sortie.HandleSortie)

	http.HandleFunc("/SortieInst", ControllerSortieInst.HandleSortieInst)

	http.HandleFunc("/SortieArticle", controllerSortieArticle.HandleSortieArticle)

	http.HandleFunc("/receiveSortInst", ControllerSortieInst.ReceiveAjax)

	http.HandleFunc("/OperationSortie", controllerSortieArticle.OperationSortie)

	http.HandleFunc("/HistoSortRetour",HistoSortRetour.HandleSortRetour)

	http.HandleFunc("/Histo", Historique.HandleHisto)

	http.Handle("/Login_v1/",http.StripPrefix("/Login_v1",http.FileServer(http.Dir("Login_v1/"))))

	http.Handle("/webPages/",http.StripPrefix("/webPages",http.FileServer(http.Dir("webPages/"))))


		err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
