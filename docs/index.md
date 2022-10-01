## O que é este repositório?

Este repositório armazena diversos exercícios para a fixação dos conceitos básicos Docker, juntamente com um laboratório pré criado para ser executado em máquina virtual.

## Objetivo

O objetivo é utilizar esse material como um apoio para os alunos que se interessam por Docker, passando alguns desafios aos quais possam tentar resolver por conta própria.

## Índice de atividades

- Parte 1 - [Iniciando Build](atividades/01-Iniciando_Build.md)
- Parte 2 - [Atualizando a aplicação](atividades/02-Atualizando_a_aplicacao.md)
- Parte 3 - [Usando repositórios registry](atividades/03-Usando_repositorios_registry.md)
- Parte 4 - [Persistências e pontos de montagem](atividades/04-Persistencias_e_pontos_de_montagem.md)
- Parte 5 - [Multi-container](atividades/05-Multi-container.md)
- Parte 6 - [Usando Docker Compose](atividades/06-Usando_docker-compose.md)

## Testes

Alguns dos desafios possuem o diretório `test`, nele estão contidos `scripts` para testes de infraestrutura com o objetivo de validar se o exercício foi concluído corretamente.

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

## Laboratório

O laboratório é provisionado utilizando as ferramentas [**Virtual Box**](https://www.virtualbox.org/) e [**Vagrant**](https://www.vagrantup.com/).

É necessário a instalação dessas ferramentas para a criação do laboratório, sigam as instruções de instalação na documentação de cada ferramenta. Ambos são compatíveis com sistemas Windows, Linux e macOS.

### Motivação
Este laboratório foi criado com o intuito de ganhar tempo, padronização e praticidade na hora de montar o ambiente para aprendizado.

### O ambiente
O ambiente é executado com duas máquinas virtuais, uma chamada **lab** e outra **client**

1. VM lab: É a máquina aonde 95% das atividades ocorrem.
2. VM client: É utilizada somente durante a atividade Parte 3 - Usando repositório registry para demonstração.

### Requisitos
Para o bom funcionamento do ambiente, os requisistos **minímos** são:

```yaml
lab:
  memória ram: 4GB
  vcpu: 2
  espaço em disco: 10GB

client:
  memória ram: 2GB
  vcpu: 1
  espaço em disco: 10GB
```

> Os recursos de memória ram e cpu podem ser alterados no arquivo [config.yaml](config.yaml).

### Comandos básicos para gerenciar o laboratório:
Para a inicialização do ambiente:
```sh
$ vagrant up
```
> O tempo médio da primeira incialização é de 5 minutos (dependendo da velocidade de conexão com a Internet).

Para acessar o ambiente **lab**:
```sh
$ vagrant ssh lab
```

Para acessar o ambiente **client**
```sh
$ vagrant ssh client
```

Após o acesso ao ambiente, abra o diretório `/vagrant` para começar os execícios.
```sh
$ cd /vagrant
```
Para remoção completa do ambiente:
```sh
$ vagrant destroy
```
> Esse comando removerá a vm por completa do Virtual Box, no entanto, o conteúdo do diretório `/vagrant` da vm ficará disponível neste diretório de forma intacta.
