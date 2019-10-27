package test

import (
	fs "speedup/filesystem"
	"speedup/query"
	"speedup/utils"
	"testing"
)

func TestSearchOR(t *testing.T) {

	workFolder := utils.InitializeWorkFolder()

	fileSystem := new(fs.FileSystem).CreateFileSystem("teste", workFolder)
	//IndexWriter := new(idx.IndexWriter).CreateIndex(fileSystem)

	qr := new(query.Query).CreateQuery(fileSystem)

	condiction := new(query.EQ).AddEQ(&query.Map{
		Key:   "nome",
		Value: "thiago luiz",
	}).AddEQ(&query.Map{
		Key:   "idade",
		Value: "32",
	})

	condictionOR := new(query.OR)
	condictionOR.AddOR(condiction)
	rs := qr.FilterOr(condictionOR)

	if len(rs) != 1 {
		t.Error("Resultado deve retornar apenas 1 documento")
	}

	if rs[0] != "2" {
		t.Error("O documento retornado dever ser o código 2")
	}

	///////////////////////////////////////////

	condiction1 := new(query.EQ).AddEQ(&query.Map{
		Key:   "nome",
		Value: "taliba",
	})

	condiction2 := new(query.EQ).AddEQ(&query.Map{
		Key:   "nome",
		Value: "thiago luiz",
	})

	condictionOR = new(query.OR)
	condictionOR.AddOR(condiction1).AddOR(condiction2)
	rs = qr.FilterOr(condictionOR)

	if len(rs) != 2 {
		t.Error("Resultado deve retornar apenas 2 documentos")
	}

	if rs[0] != "1" && rs[1] != "2" {
		t.Error("O documento retornado dever ser o código [1 e 2] ")
	}

	condiction3 := new(query.EQ).AddEQ(&query.Map{
		Key:   "nome",
		Value: "taliba",
	})

	condiction4 := new(query.EQ).AddEQ(&query.Map{
		Key:   "nome",
		Value: "maria saco roxo",
	})

	condictionOR = new(query.OR)
	condictionOR.AddOR(condiction3).AddOR(condiction4)
	rs = qr.FilterOr(condictionOR)

	if len(rs) != 0 {
		t.Error("Não deve retornar nenhum documento")
	}

}
