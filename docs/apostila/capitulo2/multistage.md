---
hide:
  - toc
---

### Introdução

Uma boa prática é o tamanho da imagem gerada.

Cada instrução no `Dockerfile` adiciona uma nova camada à imagem. Muitas vezes diversos artefatos e pacotes são instalados e copiados durante as instruções e não serão utilizados durante a execução do container. Aumentando assim o tamanho final da imagem.

Por boas práticas, quanto menor o tamanho da imagem, mais fácil é o transporte (*pull* e *push*), o armazenamento da imagem no host onde o container será executado, velocidade de **deployment**, etc.

O `multi-stage` é uma boa prática para quem quer manter o tamanho das imagens reduzidas e limpas.

#### Usando Multi-stage builds

A técnica de `multi-stage` permite que em apenas um `Dockerfile` é possível existir múltiplos `FROM`. Cada `FROM` faz parte de um estágio da compilação do `Dockerfile` e entre um estágio e outro, é possível copiar artefatos selecionados e mantendo arquivos não desejados fora da imagem "final".Criando uma imagem mais "limpa" e provavelmente menor.

#### Na prática

Usamos como exemplo o `Dockerfile` que está na raíz do projeto:

```Dockerfile linenums="1"
FROM squidfunk/mkdocs-material:8.5.3 AS build
COPY . /docs
WORKDIR /docs
RUN pip install -r requirements.txt && mkdocs build

FROM nginx:alpine
COPY --from=build /docs/site /usr/share/nginx/html
```

Neste exemplo temos dois estágios (`FROM`), aonde o primeiro foi usado para fazer a instalação de dependências Python e compilação de dos arquivos de *front-end*. Note que na linha #1 temos a sintaxe `AS build`, essa é a referência do estágio.

Já no segundo estágio, foi usado uma imagem `nginx:alpine` e copiamos o resultado da compilação do *front-end* para o diretório `/usr/share/nginx/html`. Assim o serviço de `nginx` (web-server) vai inciar somente com os arquivos necesários para disponibilizar o site.
