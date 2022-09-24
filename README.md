## O que é este repositório?

Este repositório armazena diversos exercícios para a fixação dos conceitos básicos Docker.

## Objetivo

O objetivo é utilizar esse material como um apoio para os alunos que se interessam por Docker, passando alguns desafios aos quais possam tentar resolver por conta própria.

## Testes

Alguns dos desafios possuem o diretório `test`, nele está contido `scripts` para testes de infraestrutura e validar se o exercício foi concluído corretamente.

A ideia é que o próprio aluno(a) consiga entender quais etapas dos desafios não foram concluídas corretamente, ajudando na auto correção dos desafios.

### Como executar os testes

Os testes foram escritos na linguagem de programação [**Golang**](https://go.dev/) boa parte utilizando a biblioteca [Terratest](https://terratest.gruntwork.io/) e [Testify](https://pkg.go.dev/github.com/stretchr/testify).

A execução do teste deve ser feito de dentro do diretório `./test` utilizando o seguinte comando:
```sh
$ go test -v -failfast

=== RUN   TestIfImageExists
    00_test.go:18: 
        	Error Trace:	00_test.go:18
        	Error:      	Should be true
        	Test:       	TestIfImageExists
--- FAIL: TestIfImageExists (0.02s)
FAIL
exit status 1
FAIL	test	0.027s
```
