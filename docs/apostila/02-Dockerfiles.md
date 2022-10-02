---
hide:
  - toc
---

## Dockerfiles

Dockerfiles são arquivos que contém instruções para a criação de uma imagem **Docker**.


Para criar imagens Docker através de um Dockerfile, usamos o seguinte comando:

```bash
docker build -f Dockerfile . -t registry/imagem:ver
```

Onde:
- registry: Nome de onde a imagem será salva
- imagem: O nome da imagem
- ver: É a tag para versionamento da imagem


### Instruções

Algumas instruções básicas utilizadas num `Dockerfile`.

| Instrução      | Descrição                          |
| ----------- | ------------------------------------ |
| `FROM`       | Indica a imagem base para a criação de uma nova imagem. |
| `RUN`       | Executa qualquer comando durante a construção da imagem.  |
| `WORKDIR`    | Define em qual diretórios instruções seguintes serão executados `RUN, CMD, ENRTYPOINT, COPY`  |
| `COPY`    |  Copia arquivos do host atual para a imagem |
| `CMD`    | Providencia comandos de inicialização para um container em execução |

> O `CMD` pode ser utilizado somente uma vez em cada imagem.

Mais detalhes sobre os comandos podem ser encontrados na documentação sobre [Dockerfile](https://docs.docker.com/engine/reference/builder/).

### Exemplo de Dockerfile

```Dockerfile linenums="1"
FROM alpine:3.16.2
WORKDIR /myapp
RUN echo “Hello, World!” > test.txt
CMD ["cat", "test.txt"]
```
> Repare que o arquivo `test.txt` é criado dentro do diretório `/myapp` devido a instrução `WORKDIR` na linha 2

