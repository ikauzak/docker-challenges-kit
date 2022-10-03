---
hide:
  - toc
---

### Diferenças

O comando `ENTRYPOINT` serve para definir a execução de um binário no momento em que o container é iniciado. Geralmente esse parâmetro é bastante utilizado para containers que carregam binários compilados para execuções de comandos. Exemplo `terraform`, `gcloud`, `packer`, `hadolint`, etc... Dessa forma ao executar o container, não é necessário passar o binário no comando de execução:

```shell
docker run -it --rm docker.io/hashicorp/terraform:1.3.0 -help
Usage: terraform [global options] <subcommand> [args]

The available commands for execution are listed below.
The primary workflow commands are given first, followed by
less common or more advanced commands.

Main commands:
  init          Prepare your working directory for other commands
  validate      Check whether the configuration is valid
  plan          Show changes required by the current configuration
  apply         Create or update infrastructure
  destroy       Destroy previously-created infrastructure

Global options (use these before the subcommand, if any):
  -chdir=DIR    Switch to a different working directory before executing the
                given subcommand.
  -help         Show this help output, or the help for a specified subcommand.
  -version      An alias for the "version" subcommand.
```

Já o `CMD` serve como um **argumento** para o `ENTRYPOINT`. No exemplo acima, o `ENTRYPOINT` (comando oculto) foi a chamada do binário `terraform` e o seu **argumento** foi o `-help`.


#### Na prática
Neste exemplo vamos entender melhor a diferença entre o `ENTRYPOINT` e o `CMD`. Criando uma imagem `Docker` que possui 5 arquivos do tipo `txt` dentro do diretório `/myapp`.

O comando `ENTRYPOINT` definido neste exemplo é um `cat`, um simples comando para leitura de arquivo de texto.

Já o `CMD` foi pré definido no exemplo como `test_1.txt`, no entanto é possível sobreescrever esse valor durante a exexecução.

A ideia é utilizar as sintaxes `CMD` e `ENTRYPOINT` e usá-las durante algumas execuções.

1. Crie um Dockerfile com o conteúdo abaixo:
```Dockerfile linenums="1" title="capitulo2/exemplos/2/Dockerfile"
--8<--
docs/apostila/capitulo2/exemplos/2/Dockerfile
--8<--
```
2. Faça o *build* e execute algumas vezes
``` shell
$ docker build . -t teste:leitura
Successfully built d4e2d672eb27
$ docker run -it --rm teste:leitura
Número 1
$ docker run -it --rm teste:leitura test_2.txt
Número 2
$ docker run -it --rm teste:leitura teste_5.txt
Número 5
```
Perceba ao não passar nenhum nome de arquivo (`test_2.txt`, `test_3.txt`, etc) no fim do comando, por padrão o container vai ler o arquivo `test_1.txt` conforme a instrução `CMD` que foi incluída no Dockerfile. É possível sobreescrever o valor padrão do `CMD` ao passar o parâmetro no fim do comando, conforme executados no exemplo.

