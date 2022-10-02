---
hide:
  - toc
---

## Sobre

São volumes gerenciados automaticamente pelo `Docker` e que ficam armazenados no host e podem ser criados através da CLI.

Algumas vantagens de se usar `named volumes` é o seu fácil gerenciamento (CLI), pode ser compartilhado entre vários containers, facilidade para **backups** e entre outras vantagens.

Na prática, um `named volume` é um volume com um rótulo para facilitar seu gerenciamento, com isso podemos utilizá-lo durante a inicialização e mapear o volume em qualquer diretório dentro de um container. Esse mesmo volume também pode ser utilizado em outros containers para outras finalidades.

### Exemplo de escrita em volume

Para ilustrar, neste exemplo vamos criar dois containers ao mesmo tempo com duas finalidades diferentes.

O primeiro vai ser o container que de fato que escreve no arquivo que estará dentro de um volume. E o segundo container será usado para acompanhar a escrita do primeiro e comprovar que os dados estão compartilhados e não serão perdidos em caso de deleção do primeiro container.

1. Vamos criar um novo volume com o nome de 'persistencia':
```shell
$ docker volume create persistencia
persistencia
$ docker volume ls
DRIVER    VOLUME NAME
local     persistencia
```

2. Crie um `Dockerfile` com este conteúdo e uma imagem com a tag `lab-persistencia:v1`
```Dockerfile linenums="1" title="capitulo4/exemplos/1/Dockerfile"
--8<--
docs/apostila/capitulo4/exemplos/1/Dockerfile
--8<--
```
> Esta imagem executa um `script` que salva data, hora e o nome do host num arquivo chamado `arquivo_persistente.txt`.

3. Crie um novo container usando a nova imagem `lab-persistencia:v1` montando o volume no diretório `/outputs`, pois segundo no `Dockerfile` foi definido que a aplicação será executada no diretório `/outputs` dentro do container:
```bash
$ docker run -dit --rm -v persistencia:/outputs lab-persistencia:v1
...
```

4. Enquanto o primeiro container está em execução, inicie um segundo container montando o mesmo volume `persistencia` no diretório `/view`:
```bash
$ docker run -it -w /view --rm -v persistencia:/view ubuntu:20.04
a50e8ca80746
```

5. Leia o arquivo `arquivo_persitente.txt` enquanto o primeiro container continua escrevendo no mesmo:
```bash
$ root@a50e8ca80746: tail -f arquivo_persistente.txt
```

6. Acompanhe a escrita do arquivo e perceba que o arquivo está compartilhado entre os dois containers. Caso o primeiro container pare de executar, volte a executar o comando no passo `3.` e veja que o nome do host será alterado nas novas entradas do arquivo `arquivo_persistente.txt`

### Deleção do volume

O volume permanece criado mesmo após as deleções de imagens e containers, pelo motivo de ser um objeto totalmente "desacoplado" e independente, podendo ser montado novamente em novos containers ou compartilhado entre vários outros em execução.

Para a deleção do volume, use o comando:
```bash
$ docker volume ls
DRIVER    VOLUME NAME
local     persistencia
$ docker volume rm persistencia
```

> Atenção para não deletar o volume incorreto!
