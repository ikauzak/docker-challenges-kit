---
hide:
  - toc
---

## Contador web server

Para ilustrar na prática o funcionamento de uma aplicação multi containers, vamos criar uma aplicação em Python para contar a quantidade de requests que foram feitas no navegador e salvar esses dados num banco [Redis](https://redis.io/docs/about/).

A aplicação é simples, o primeiro container (`py-web-counter`) precisa estar conectado na mesma rede `Docker` do banco Redis. No código da aplicação, a comunicação acontece através do nome de serviço `redis` que está pré-escrita no código.


```py linenums="1" title="capitulo6/exemplos/1/app.py"
--8<--
docs/apostila/capitulo6/exemplos/1/app.py
--8<--
```

```Dockerfile linenums="1" title="capitulo6/exemplos/1/Dockerfile"
--8<--
docs/apostila/capitulo6/exemplos/1/Dockerfile
--8<--
```

```linenums="1" title="capitulo6/exemplos/1/requirements.txt"
--8<--
docs/apostila/capitulo6/exemplos/1/requirements.txt
--8<--
```
