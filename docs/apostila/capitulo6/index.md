---
hide:
  - toc
---

## Introdução

Aplicações que são compostas por mais de um container são chamadas de multi container.

Por exemplo, aplicações que são compostas por um front-end, back-end e um banco de dados.

``` mermaid
graph LR
  A[request] --> |1| B[APP Front-end];
  B --> |2| C[APP Back-end];
  C -->|3| D[DB];
  D --> |4| C;
  C --> |5| B;
  B --> |6| A;
```

Como boa prática cada container deve ser apenas um único processo em execução, alguns dos motivos são:

- Melhor controle de versão entre aplicações
- Escalabilidade. Cada aplicação pode ser escalado de forma independente.
- Menor complexidade para *rollbacks*.


