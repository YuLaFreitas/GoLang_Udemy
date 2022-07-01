package controllers

import "net/http"

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criar usuario"))
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
