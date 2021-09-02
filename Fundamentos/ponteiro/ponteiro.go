package main

func main() {
	i := 1

	var p *int = nil

	p = &i //pegando o endereco
	// da variavel
	// p = endereco, *p= valor
	//nil = nulo
	*p++
	i++
	// da no msmo

}
