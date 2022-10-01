---
hide:
  - toc
---
## O que é Docker?

 O Docker ajuda a abstrair tarefas de configuração repetitivas, tornando o ciclo de desenvolvimento mais ágil e facilitando a portabilidade da aplicação entre ambientes distintos, seja em um desktop, nuvem ou data center.

O ciclo de vida:

- **Construa**: É onde tudo começa, a compilação (*build*) de um Dockerfile e gerando uma imagem.
- **Compartilhe**: Compartilhe imagens Docker através de repositórios [registry](https://hub.docker.com/).
- **Execute**: Execute a imagem em qualquer ambiente docker seja ele num notebook, servidor em nuvem, on-premises ou qualquer outra plataforma que suporta execução de containers como o [Kubernetes](https://kubernetes.io/docs/concepts/overview/).

## O que é um container?

Em sua definição, um container é uma aplicação com todas as suas dependências e códigos empacotados para a execução rápida de uma aplicação.

Na técnica, um container é um **processo** isolado que está em execução dentro do sistema.

Para entender na prática, vamos executar um simples container:

```sh
$ docker run -dit --name teste-proc --rm ubuntu:20.04
cf4243fb171f200c00f6ce261a6bef0ad4df59371b451297d1f079463d453095
```

> No momento, não se preocupe em **decorar** o comando e sim em entender o conceito.

Verifique se o container está em execução através dos processos abertos no sistema:

```{ .bash .annotate }
$ ps aux | grep bash
masuda    454295  0.0  0.0   9972   476 ?        S    set28   0:00 bash /usr/bin/remmina-file-wrapper
root      871174  0.0  0.0   4116  3320 pts/0    Ss+  21:52   0:00 bash
masuda    871337  0.0  0.0   9212  2224 pts/0    S+   21:52   0:00 grep --color bash

$ docker exec -ti teste-proc ps aux
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.0  0.0   4116  3320 pts/0    Ss+  00:52   0:00 bash
root          15  0.0  0.0   5900  2900 pts/1    Rs+  01:03   0:00 ps aux
```
Ao listar os processos do host, vemos um processo chamado **bash**, que apesar do **ID do processo** ser diferente do container, ambos possuem o mesmo comando (**bash**).

Caso o processo seja terminado no host, o Docker também encerra o container.

```bash
$ sudo kill -9 871174

$ docker container ls
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```

## Imagens docker

Todo container utiliza uma imagem pré-compilada com todos os pacotes necessários para a sua execução, ajudando assim o container a inicializar a aplicação com maior agilidade.

Imagens podem ser customizadas e criadas através de **Dockerfile**.

## Comparação entre VM e containers
<figure markdown>
  ![Image title](./img/container-vm-whatcontainer_2.png.webp){ width="400" align=left}
<figcaption>Host VM</figcaption>
</figure>

<figure markdown>
  ![Image title](./img/docker-containerized-appliction-blue-border_2.png.webp){ width="400" align=left}
<figcaption>Docker host</figcaption>
</figure>
