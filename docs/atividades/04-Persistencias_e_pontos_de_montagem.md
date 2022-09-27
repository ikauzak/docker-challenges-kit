## Objetivos

Este exercício tem como principal objetivo demonstrar e aplicar tipos de volumes Docker e suas persistências.

Diferenciar entre o *Named Volumes* e *Bind Mounts*.

## Atividades

### Persistência de dados com *Named Volumes*

A aplicação de demonstração (todo app) armazena os dados criados no caminho `/etc/todos/todo.db` dentro do sistema de arquivos do container.

Neste caso, vamos usar um *Named Volume* para persistir os dados da aplicação para serem reutilizados futuramente.

1. Crie um volume Docker com o nome de `todo-db`.

2. Pare e remova qualquer container de demonstração que esteja em execução.

3. Inicie um novo container adicionando a opção `-v` passando o nome do volume `todo-db` montando no caminho `/etc/todos` do container.

4. Ao iniciar, acesse a aplicação e crie algumas tarefas na interface.

5. Pare e remova o container.

6. Inicie um novo container com os mesmos comandos da instrução número `3`.

7. Acesse a aplicação e verifique se as tarefas ainda estão registradas.

### Utilizando *Bind Mounts*

Com o *Bind Mounts* podemos mapear qualquer diretório do **host** dentro de qualquer caminho do container para persistir os dados no diretório do **host**.

Neste exemplo criaremos um container de desenvolvimento para entendermos o conceito de *Bind Mounts* no Docker.

Para aplicações baseadas em Node, a biblioteca [nodemon](https://www.npmjs.com/package/nodemon) é uma ótima ferramenta para monitorar alterações dos arquivos e reiniciar a aplicação automaticamente.
> Existem outras ferramentas semelhantes para outras linguagens de programação.

Nesta etapa, montaremos um ambiente de desenvolvimento Node para entender o conceito de *Bind Mounts*.

1. Verifique se não há mais nenhum container em execução dentro do host.

2. Execute o seguinte comando para a inicialização do container em modo de desenvolvimento:
```shell
$ docker run -dp 3000:3000 \
     -w /app -v "$(pwd):/app" \
     node:12-alpine \
     sh -c "yarn install && yarn run dev"
```
- `-dp 3000:3000` - executa o container em *background* na porta 3000 do host
- `-w /app` - indica em qual diretório o conteiner será inicializado
- `-v "$(pwd):/app"` - cria um *bind* entre o diretório atual do seu `shell` com o diretório `/app` do container
- `node:12-alpine` - é a mesma imagem utilizada no Dockerfile para a criação da imagem
- `sh -c "yarn install && yarn run dev` - o comando que instala as dependências e inicia a aplicação em modo de desenvolivmento. Veja no arquivo `package.json` que no script `dev` está inicializando `nodemon`

3. Veja os logs do container até que a mensagem `Listening on port 3000` apareça.

4. Vamos alterar a mensagem do botão "Add item" para "Add". Acesse o arquivo `src/static/js/app.js` na linha 109, e faça as alterações:
```
 -                         {submitting ? 'Adding...' : 'Add Item'}
 +                         {submitting ? 'Adding...' : 'Add'}
```

5. Atualize a página e veja que a alteração foi feita.

6. Ao terminar as alterações, execute um novo `build` para criar uma nova versão da aplicação.
