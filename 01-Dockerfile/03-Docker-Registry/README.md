## Objetivos

Este exercício tem como principal objetivo demonstrar como funciona um `registry` para hospedar diferentes tipos de imagens.

Aprender os conceitos dos comandos `pull` e `push` no contexto Docker.

## Atividades

### Subir um registry docker localmente

Nesta etapa criaremos um `registry` local para ser usado em laboratório.

```sh
$ docker run -d \
  -p 5000:5000 \
  --restart=always \
  --name registry \
  -v /mnt/registry:/var/lib/registry \
  registry:2
```

> Este exemplo é para uso experimental somente, não utilize essa solução para ambientes produtivos.

### Suba as imagens no registry local

1. Crie uma tag com base das imagens criadas nos exercícios `01` e `02` como `localhost:5000/helloworld:v1` e `localhost:5000/helloworld:v2`.
2. Faça o `push` das duas imagens para registry local (criado na etapa **Subir um registry docker localmente**)

### Fazendo download das imagens do registry local

1. Faça o `pull` das imagens que foram gravadas no registry local para o seu ambiente
2. Verifique se as imagens estão na sua lista de imagens
