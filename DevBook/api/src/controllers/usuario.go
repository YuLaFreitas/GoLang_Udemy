package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {

		log.Fatal(erro)
	}

	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Concetar()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	reposotorio := repositorios.NovoReposotioDeUsuarios(db)

	usuarioID, erro := reposotorio.CriarInBD(usuario)

	if erro != erro {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioID)))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuario"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BUscar por id"))
}

func AtualizandoUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar"))
}

func DeletarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deltar"))
}
