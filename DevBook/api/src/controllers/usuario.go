package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != erro {
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

	reposotorio := repositorios.NovoReposotioDeUsuarios(db)
	reposotorio.CriarInBD(usuario)

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
