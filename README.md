## O que é este repositório?

Este repositório armazena diversos exercícios para a fixação dos conceitos básicos Docker, juntamente com um laboratório pré criado para ser executado em máquina virtual.

## Objetivo

O objetivo é utilizar esse material como um apoio para os alunos que se interessam por Docker, passando alguns desafios aos quais possam tentar resolver por conta própria.

## Sobre o diretório `app`
O diretório `app` é onde os exercícios serão feitos e executados. Trata-se de uma aplicação escrita em `node` para ter como referência durante as execuções e testes.

Todos os arquivos deverão ser criados dentro do diretório `app`

## Sobre docs
Um site com as instruções dos exercícios será criado durante o processo do `make start` e disponibilizado localmente para uma melhor leitura e compreensão das atividades.

## Sobre o Laboratório

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
  memória ram: 1GB
  vcpu: 1
  espaço em disco: 10GB
```

### `vagrant_config.yaml`
Os recursos de memória ram e cpu podem ser alterados no arquivo [vagrant_config.yaml](vagrant_config.yaml).

### Rede
É importante verificar o bloco de endereço IP que está em uso pelo VirtualBox, por padrão o script utiliza a rede "vboxnet0". Os endereços de cada VM pode ser editado no arquivo `vagrant_config.yaml`.

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

Após o acesso ao ambiente, abra o diretório `/vagrant` para começar os execícios.
```sh
$ cd /vagrant
```
Para remoção completa do ambiente:
```sh
$ vagrant destroy
```
> Esse comando removerá a vm por completa do Virtual Box, no entanto, o conteúdo do diretório `/vagrant` da vm ficará disponível neste diretório de forma intacta.
