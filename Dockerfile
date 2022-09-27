## Este arquivo é para a crição da imagem mkdocs para leitura de documentação

FROM squidfunk/mkdocs-material:8.5.3 AS build
COPY . /docs
WORKDIR /docs
RUN mkdocs build

FROM nginx:alpine
COPY --from=build /docs/site /usr/share/nginx/html
