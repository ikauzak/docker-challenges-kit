## Objetivos

Este exercício tem como objetivo demonstrar o de uma aplicação em multi containers e como essas aplicações se comunicam entre si através da rede Docker.

Entender conceitos básicos de DNS.

## Atividades

### Iniciando banco MySQL

1. Crie uma rede docker com o nome `todo-app`

2. Inicie um container com os seguintes parametros:
- network-alias: `mysql`
- network: `todo-app`
- named volume: `todo-mysql-data` mapeado dentro do diretório `/var/lib/mysql`
- var1: `MYSQL_ROOT_PASSWORD` com valor `secret`
- var2: `MYSQL_DATABASE` com o valor `todos`
- imagem: `mysql:5.7`

3. Confirme que o banco de dados está em execução e verifique se há conectividade com o banco. Utilize o comando `mysql -u root -p` para abrir uma sessão no banco.

4. Preencha com a senha `secret` para se conectar ao banco

5. Dentro do banco, execute o seguinte comando:
```shell
mysql> SHOW DATABASES;
```
Verifique se na saída do comando a tabela `todos` aparece.

6. Saía da sessão `mysql`
```shell
mysql> exit
```

### Conectando a aplicação dev ao banco MySQL
1. Inicie a aplicação em modo de desenvolvimento utilizando os seguintes parâmetros: 
- port: `3000` do host com a `3000` do container
- working dir: `/app`
- network: `todo-app`
- var1: `MYSQL_HOST` com o valor `mysql`
- var2: `MYSQL_USER` com o valor `root`
- var3: `MYSQL_PASSWORD` com o valor `secret`
- var4: `MYSQL_DB` com o valor `todos`
- imagem: `node:12-alpine`
- entrypoint: `sh -c "yarn install && yarn run dev"`

2. Acesse a aplicação por um navegador, e crie algumas tarefas.

3. Conecte no banco novamente, e verifique se as tarefas estão sendo gravadas no banco:
```shell
mysql> select * from todo_items;
```
