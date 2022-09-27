## Objetivos
O principal objetivo é entender como funcionam as `tags`de imagens Docker e o conceito de infraestrutura imutável. :tada:

## Atividades

### Atualize o código
1. Acesse o arquivo `src/static/js/app.js`, altere o conteúdo da linha #56:
```
-                <p className="text-center">No items yet! Add one above!</p>
+                <p className="text-center">Add one todo above!</p>
```
2. Faça uma nova imagem com a tag `helloworld:v2`.

3. Remova o container com a tag `:v1` e inicie um novo container com a tag `:v2` recém criada.

4. Abra a aplicação no navegador e verique se o seu código foi atualizado com sucesso.
