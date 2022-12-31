package rotas

import (
	"api/src/router/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/ususarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/ususarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/ususarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/ususarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/ususarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
