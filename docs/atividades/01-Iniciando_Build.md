## Objetivos
O objetivo desta primeira atividade é entender na prática as principais sintaxes de um `Dockerfile`.

## Atividades

1. Crie um arquivo `Dockerfile` dentro do diretório `app` que encontra-se na raíz deste repositório.
2. Cole o conteúdo abaixo no arquivo `Dockerfile`

```
FROM node:12-alpine
RUN apk add --no-cache python2 g++ make
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "src/index.js"]
EXPOSE 3000
```

3. Crie a imagem com a seguinte a tag `helloworld:v1`

4. Execute um container utilizando a imagem criada:
```
$ docker run -d -p 3000:3000 helloworld:v1
```

5. Acesse a aplicação através do seu navegador do *desktop* utilizando o endereço da vm `lab` na porta 3000. Exemplo: `http://192.168.56.101:3000`
> Verifique qual é o endereço interno da vm utilizando o comando `ip a`.
