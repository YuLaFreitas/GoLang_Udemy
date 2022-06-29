package rotas

import (
	"api/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/Usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuarios,
		RequerAutenticacao: false,
	},

	{
		URI:                "/Usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},

	{
		URI:                "/Usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:                "/Usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoUsuarios,
		RequerAutenticacao: false,
	},

	{
		URI:                "/Usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuarios,
		RequerAutenticacao: false,
	},
}
