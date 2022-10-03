---
hide:
  - toc
---

## Rede Docker

Com `Docker` podemos isolar quais containers possuem comunicação entre si, além de também ter um registro de DNS interno entre as aplicações pra se acharem em execução.

### Drivers de rede

Existem diversos drivers de rede `Docker` que gerenciam o comportamento de comunicação dos containers que estão em execução. Dentre todos, os mais utilizados são do tipo `bridge`, `host` e `overlay`.

- `bridge`: É o driver de rede padrão utilizado no `Docker`, esse driver gerencia o tráfego entre containers na camada `Docker`, podendo incluir ou remover containers em redes do tipo `bridge` a qualquer momento. Por padrão, containers que **fazem parte da mesma rede `bridge`** têm todas as suas portas expostas entre si.
- `host`: Mais utilizada para aplicações standalone, remove o isolamento de comunicação entre o host `Docker` e o container. É importante publicar a porta do container com a flag `-p` para o mesmo ser acessível externamente.
- `overlay`: É uma rede virtual que é criada para a comunicação entre diversos hosts `Docker`. Usada principalmente quando existe um cluster `Swarm` (múltiplos hosts `Docker` comunicando entre si).

Importante destacar que a rede padrão `bridge` (pré configurada na instalação) não resolve DNS por nomes de containers, a comunicação só funciona por endereçamento IP, que não é a melhor opção para um ambiente produtivo `Docker`.
