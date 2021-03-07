package models

import "github.com/alanrdel/db"

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade, Id  int
}

func BuscaProdutos() []Produto {
	db := db.ConectaBD()

	selectProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarNovosProdutos(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBD()
	//verificar se a query esta corrreta
	insertDadosProduto, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	insertDadosProduto.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaBD()
	delProdudo, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	delProdudo.Exec(id)
	defer db.Close()
}

func EditarProduto(id string) Produto {
	db := db.ConectaBD()
	produtoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoAtualizar := Produto{}
	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoAtualizar.Id = id
		produtoAtualizar.Nome = nome
		produtoAtualizar.Descricao = descricao
		produtoAtualizar.Preco = preco
		produtoAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoAtualizar
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBD()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
