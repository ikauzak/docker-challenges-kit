---
hide:
  - toc
---

## Executando múltiplos containers de uma única vez

A ferramenta `docker compose` serve para definir e compartilhar aplicações multi container.

Utilizando uma sintaxe `YAML`, o arquivo `docker-compose.yaml` define todas as características dos containers e com apenas **um comando**, é possível executar todos os serviços de uma vez só.

## Docker compose contador web server

Para exemplificar a criação de um arquivo `docker-compose`, vamos usar como exemplo a mesma aplicação demonstrada no capítulo anterior.

1. Exemplo de arquivo docker-compose.
```yaml linenums="1" title="capitulo6/exemplos/2/docker-compose.yaml"
--8<--
docs/apostila/capitulo6/exemplos/2/docker-compose.yaml
--8<--
```
Neste arquivo é importante destacar:
    - `version`: É a versão de escrita do docker compose.
    - `services`: São as definições de cada container (`web` e `redis`).
    - `build`: Indica qual o caminho relativo de um arquivo `Dockerfile` para criar uma imagem.
    - `ports`: Expondo a porta `5000` do host na porta `5000` do container (que é onde o webserver python é executado em código).
    - `depends_on`: É uma lista informando uma ordem de dependências que devem estar em funcionamento antes.
    - `image`: A definição da imagem do serviço `redis`, no formato `nome:tag`.

2. Inicie os serviços:
```shell
$ docker compose up -d
...
[+] Running 3/3
 ⠿ Network 2_default    Created                                                                                                   0.0s
 ⠿ Container 2-redis-1  Started                                                                                                   0.9s
 ⠿ Container 2-web-1    Started                                                                                                   0.9s
```
Veja que o docker compose cria automaticamente uma rede para a comunicação entre os dois containers.

3. Verifique os containers em execução:
```shell
$ docker container ls
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS                                       NAMES
d9126c61c7c1   2-web          "flask run"              15 minutes ago   Up 15 minutes   0.0.0.0:5000->5000/tcp, :::5000->5000/tcp   2-web-1
796325276f8c   redis:alpine   "docker-entrypoint.s…"   15 minutes ago   Up 15 minutes   6379/tcp                                    2-redis-1
```

4. Acesse a aplicação por um navegador através da porta `5000`.

5. Após validar os acessos e containers, remova o ambiente:
```shell
$ docker compose down
[+] Running 3/3
 ⠿ Container 2-web-1    Removed                                                                                                  11.0s
 ⠿ Container 2-redis-1  Removed                                                                                                   0.8s
 ⠿ Network 2_default    Removed                                                                                                   0.1s
```

Esse é um exemplo simples de como um `docker compose` se comporta e seu objetivo.

Existem diversas sintaxes para adaptar aplicações em multi containers, verifique a documentação [docker compose](https://docs.docker.com/compose/compose-file/)

## Modificando para desenvolvimento

Da mesma forma que usamos o `bind volumes` para criar um ambiente de desenvolvimento, é possível fazer o mesmo com `docker compose`.

```yaml linenums="1" title="capitulo6/exemplos/3/docker-compose.yaml"
--8<--
docs/apostila/capitulo6/exemplos/3/docker-compose.yaml
--8<--
```
Neste exemplo, foi adicionado dois blocos para criarmos um ambiente de desenvolvimento dessa aplicação:

- `volumes`: Essa sintaxe é a mesma referência ao usar `bind volumes` em `Docker`, ou seja, no caminho relativo do arquivo `docker-compose.yaml` estamos montando o arquivo `app.py` dentro do diretório `/code` do container.
- `environment`: Esse bloco é um dicionário de chaves e valores para incluir variáveis de ambientes dentro do container. Como a aplicação usa a biblioteca `Flask`, quando a variável de ambiente `FLASK_DEBUG` está com valor `True`, a cada alteração nos arquivos o `Flask` refaz a leitura e processa os arquivos alterados.

Vamos testar alterando o arquivo `docs/apostila/capitulo6/exemplos/1/app.py`:

```py linenums="1"

-    return 'Hello World! I have been seen {} times.\n'.format(count)
+    return 'Hello from Docker!! I have been seen {} times.\n'.format(count)
```

Ao atualizar a página, vemos que a mensagem foi alterada na tela sem a necessidade de reiniciarmos a aplicação.
