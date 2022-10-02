---
hide:
  - toc
---

## Sobre

A montagem `bind` é uma forma de montar arquivos e diretórios do host `Docker` dentro de um container em execução. Esse arquivo ou diretório é referênciado pelo seu caminho absoluto do host.

Uma das vantagens de uso do `bind` é a possibilidade de usar containers com arquivos que estão sendo desenvolvidos e alterados a todo momento, possibilitando a visibilidade das alterações em tempo real, sem a necessidade de um reinício de container/processo para os efeitos serem aplicados.


### Exemplo de uso para desenvolvimento

Neste exemplo vamos usar a aplicação `nginx` para entender o conceito de `bind` de volumes.

Vamos executar um webserver e acessar o seu front para ver o conteúdo estático, e em seguida vamos alterar o arquivo `index.html` e ver essas alterações serem realizadas diretamente no front-end.

1. Crie um arquivo `index.html`:
```html linenums="1" title="capitulo4/exemplos/2/index.html"
--8<--
docs/apostila/capitulo4/exemplos/2/index.html
--8<--
```

2. Por padrão, o `nginx` lê os arquivos que estão no diretório `/usr/share/nginx/html/`, por este motivo vamos fazer um `bind` do caminho absoluto do arquivo `index.html` para dentro do diretório padrão do `nginx`:
```shell
$ docker run -dit --rm -p 8080:80 -v $(pwd):/usr/share/nginx/html/ nginx:alpine
```

3. Através de um navegador ou comando `curl`, acesse o webserver utilizando o endereço IP da máquina `Docker` na porta `8080` como definido no passo `2.`

4. Abra o arquivo no host, e faça alguma alteração no arquivo `index.html`
```htlm linenums='1'
<body>
-  <h2>Hello from Nginx container</h2>
</body>

<body>
+  <h1>Hello from Nginx!!!!!11111</h1>
</body>
```

5. Faça um "refresh" no navegador e veja se o front foi alterado.

De um modo geral, este exemplo foi apenas uma ilustração de como podemos usar um container como webserver para processar arquivos `html` sem que haja a necessidade de instalação de algo na máquina host.
