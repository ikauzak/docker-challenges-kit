---
hide:
  - toc
---

## Contador web server

Para ilustrar na prática o funcionamento de uma aplicação multi containers, vamos criar uma aplicação em Python para contar a quantidade de requests que foram feitas no navegador e salvar esses dados num banco [Redis](https://redis.io/docs/about/).

A aplicação é simples, o primeiro container (`py-web-counter`) precisa estar conectado na mesma rede `Docker` do banco Redis. No código da aplicação, a comunicação acontece através do nome de serviço `redis` que está pré-escrita no código.


1. O código Python que é responsável por ser executado ao inicializar o container. Destaque na linha #7 que a aplicação está esperando encontrar um serviço com o nome `redis` na porta `6379` (padrão Redis) para se conectar.
```py linenums="1" hl_lines="7" title="capitulo6/exemplos/1/app.py"
--8<--
docs/apostila/capitulo6/exemplos/1/app.py
--8<--
```

2. As dependências necessárias para a execução da aplicação:
```linenums="1" title="capitulo6/exemplos/1/requirements.txt"
--8<--
docs/apostila/capitulo6/exemplos/1/requirements.txt
--8<--
```

3. Definição do `Dockerfile` que criará a imagem:
```Dockerfile linenums="1" hl_lines="9" title="capitulo6/exemplos/1/Dockerfile"
--8<--
docs/apostila/capitulo6/exemplos/1/Dockerfile
--8<--
```
Destaque na linha #9 que é a porta de acesso à aplicação (5000).

### Execução do ambiente

1. Criar a imagem da aplicação:
```shell
$ docker built . -t py-web-counter:v1
```

2. Preparar a rede `Docker` para receber os containers:
```shell
$ docker network create multi-container-lab
```

3. Como descrito acima, a aplicação guarda suas informações de acesso num banco Redis. Precisamos inicializar o Redis antes do nosso container `py-web-counter`:
```shell
$ docker run -dit --network multi-container-lab --name redis redis:alpine
```
Aqui não precisamos expor portas do Redis, pois o mesmo estará em execução na mesma rede da aplicação. Dessa forma, as portas entre ambos estarão abertas para comunicação.

4. Inicializar a aplicação na rede `multi-container-lab`:
```shell
$ docker run -dit --rm --network multi-container-lab -p 5000:5000 py-web-counter:v1
```

5. Acesse a aplicação por um navegador através do IP do host Docker`Docker` na porta `:5000`. Faça alguns *refreshs* e veja o contador incrementando o número de acessos.
