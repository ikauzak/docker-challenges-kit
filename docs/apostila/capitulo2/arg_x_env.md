---
hide:
  - toc
---

### Diferenças

O consumo de `ARG` e `ENV` são bem diferentes um do outro, a principal distinção entre eles é o momento de uso de cada um dos comandos e se o desenvolvedor quer ou não que a variável seja persistida no container.

Abaixo segue um exemplo para entender a diferença entre um comando e outro.

Crie outro Dockerfile com o conteúdo abaixo:
```Dockerfile linenums="1"
FROM ubuntu:20.04

ARG MIN=1
ARG MAX=5

ENV ARQUIVO_PADRAO=test_1.txt

WORKDIR /myapp
RUN for n in $(seq $MIN $MAX); do echo "Número $n" > test_$n.txt; done
CMD ["bash", "-c","cat $ARQUIVO_PADRAO"]
```

#### Uso de ENV

Execute a imagem alterando a sua variável de ambiente utilizando a *flag* `-e ARQUIVO_PADRAO=test_5.txt`

```shell
$ docker build . -t teste2:env
Successfully built cc9463f890a2
$ docker run -it --rm teste2:env
Número 1
$ docker run -it --rm -e ARQUIVO_PADRAO=test_5.txt teste2:env
Número 5
```
Aqui estamos alterando o comportamento da execução do `CMD`, pois estamos alterando a variável de ambiente criada dentro do container.

Mais um outro exemplo para ilustrar o comportamento:
```shell
docker run -it --rm -e ARQUIVO_PADRAO=test_3.txt teste2:env bash
root@704c32149151:/myapp# 
root@704c32149151:/myapp# env
HOSTNAME=704c32149151
PWD=/myapp
ARQUIVO_PADRAO=test_3.txt
root@704c32149151:/myapp# echo $ARQUIVO_PADRAO
test_3.txt
```
A variável de ambiente `ARQUIVO_PADRAO` é persistida na execução do container pois foi declarada no `Dockerfile` de criação. O seu valor já não é mais o original (`test_1.txt`) pois o mesmo foi sobreescrito na incialização do container.

#### Uso de ARG

O `ARG` deve ser definido durante o **build** do `Dockerfile`, no fim do processo os `ARG` não são persistidos como variáveis de ambiente, mas podem ser utilizados como variáveis durante a construção da imagem.

```shell
$ docker build . --build-arg MIN=1 --build-arg MAX=20 -t teste3:arg
Successfully built 841b770c03c8
$ docker run -it --rm teste3:arg ls
test_1.txt   test_13.txt  test_17.txt  test_20.txt  test_6.txt
test_10.txt  test_14.txt  test_18.txt  test_3.txt   test_7.txt
test_11.txt  test_15.txt  test_19.txt  test_4.txt   test_8.txt
test_12.txt  test_16.txt  test_2.txt   test_5.txt   test_9.txt
```

Veja que no `Dockerfile` temos dois `ARG` pré definidos que são utilizados por padrão. Caso precise sobreescrever esses valores, use a *flag* `--build-arg NOME_DO_ARG=NOVO_VALOR` como sintaxe durante o **build**.

O `ARG` também pode ser definido por uma variável de ambiente no host, sem precisar ser definida ao passar **flags** `--build-arg`.

```shell
$ export MAX=10
$ echo $MAX
10
$ export MIN=2
$ echo $MIN
2
$ docker build . --build-arg MIN --build-arg MAX -t teste4:arg
Successfully built e2e0cb7b14f8
$ docker run -it --rm e2e0cb7b14f8 ls
test_10.txt  test_3.txt  test_5.txt  test_7.txt  test_9.txt
test_2.txt   test_4.txt  test_6.txt  test_8.txt
```
