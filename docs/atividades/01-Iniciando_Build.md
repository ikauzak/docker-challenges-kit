---
hide:
  - toc
---
## Objetivos
O objetivo desta primeira atividade é entender na prática as principais sintaxes de um `Dockerfile`.

## Atividades

1. Crie um arquivo `Dockerfile` dentro do diretório `app` que encontra-se na raíz deste repositório.
2. No Dockerfile, escreva os atributos abaixo seguindo as sintaxes corretamente:
    1. A imagem original deve ser um `node:12-alpine`
    2. Instale os pacotes básicos para execução da aplicação: `apk add --no-cache python2 g++ make`
    3. Altere o ponto de execução para o diretório `/app`
    4. Copie **todos** os arquivos locais para a imagem
    5. Instale pacotes **node** na imagem com `yarn install --production`
    6. O comando de execução deve ser `node src/index.js`
    7. Exponha a porta `3000` na imagem

3. Crie a imagem com a seguinte a tag `helloworld:v1`

4. Execute um container utilizando a imagem criada com as opções:
    1. O container deve ser acessível pelo host na porta `3000`
    2. A execução deve ser em "*background* (O terminal deve estar livre para mais comandos)

5. Acesse a aplicação através do seu navegador do *desktop* utilizando o endereço da vm `lab`.


> Verifique qual é o endereço interno da vm utilizando o comando `ip a`.

> Em caso de dúvida, consulte a documentação [Dockerfile](https://docs.docker.com/engine/reference/builder/) para maiores detalhes.


