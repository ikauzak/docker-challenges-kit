---
hide:
  - toc
---

## Usando volumes

Por padrão, quando um container é executado através de uma imagem, todos os arquivos do sistema são originários da imagem que foi usada em sua execução.
Quando um arquivo ou dado é criado dentro de um container, esse dado/arquivo é efêmero, ou seja, não possui persistência. Caso um container crie um arquivo durante sua execução e em seguida esse mesmo container ser deletado, arquivos que foram gerados ou que não pertencem à imagem original serão removidos juntamente com o container.

Apesar do conceito de dados efêmeros num container, existem aplicações que necessitam manter os dados, mesmo que depois o container seja deletado e/ou recriado no ambiente.

O uso de volumes também é importante para o uso de recursos de um container, pois ao usar volumes, evita o crescimento (consumo de recursos) do container dentro do host, pois os arquivos gerados serão gravados no volume ao invés da memória ram.

Para atender esse tipo de demandas e/ou técnicas para execução de container `Docker` podemos utilizar dois tipos de volumes:

- Named volumes
- Bind mounts
