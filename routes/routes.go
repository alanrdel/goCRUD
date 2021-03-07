package routes

import (
	"net/http"

	"github.com/alanrdel/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.InsertProduto)
	http.HandleFunc("/delete", controllers.DeleteProduto)
	http.HandleFunc("/edit", controllers.EditaProduto)
	http.HandleFunc("/update", controllers.UpdateProduto)
}
