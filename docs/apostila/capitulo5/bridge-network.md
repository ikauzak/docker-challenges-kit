---
hide:
  - toc
---

## Rede Bridge em prática

Neste exemplo criaremos três containers do tipo `alpine` mas somente dois deles farão parte da mesma rede `bridge` não padrão, e um container na rede `brige` padrão `Docker`.

1. Crie uma rede com o nome `lab-net`.
```shell
$ docker network create --driver bridge lab-net
147b754ad33e3d56428dd98e4313a73f06120979cb66b64d828234d243a19624
$ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
c89e6df8cd76   bridge    bridge    local
cb0231b7b183   host      host      local
147b754ad33e   lab-net   bridge    local
4a7c31b39a4e   none      null      local
```

2. Inspeciona a rede `lab-net` para ver detalhes:
```shell
$ docker network inspect lab-net

[
    {
        "Name": "lab-net",
        "Id": "147b754ad33e3d56428dd98e4313a73f06120979cb66b64d828234d243a19624",
        "Created": "2022-10-02T19:31:42.381718141-03:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "172.18.0.0/16",
                    "Gateway": "172.18.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {},
        "Options": {},
        "Labels": {}
    }
]
```

3. Crie três containers e anexe somente dois na rede `lab-net`:
```shell
$ docker run -dit --name alp-lab1 --network lab-net alpine:3.16.2 ash

$ docker run -dit --name alp-lab2 --network lab-net alpine:3.16.2 ash

$ docker run -dit --name alp-lab3 alpine:3.16.2 ash
```
Verifique se todos estão em execução:
```shell
$ docker container ls

CONTAINER ID   IMAGE           COMMAND                  CREATED          STATUS          PORTS                                       NAMES
1d03e07d6281   alpine:3.16.2   "ash"                    3 seconds ago    Up 2 seconds                                                alp-lab3
2917bba2413c   alpine:3.16.2   "ash"                    9 seconds ago    Up 9 seconds                                                alp-lab2
d6ec6fee21b0   alpine:3.16.2   "ash"                    15 seconds ago   Up 14 seconds                                               alp-lab1
```

4. Inspecione a rede `lab-net` para verificar quais containers estão conectadas à ela.
```shell
$ docker inspect network lab-net

[
    {
        "Name": "lab-net",
        "Id": "147b754ad33e3d56428dd98e4313a73f06120979cb66b64d828234d243a19624",
        "Created": "2022-10-02T19:31:42.381718141-03:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "172.18.0.0/16",
                    "Gateway": "172.18.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "2917bba2413c42e47d35191a2fee82257f0484678b5c6b8e98932ce582ac9759": {
                "Name": "alp-lab2",
                "EndpointID": "49d6ce178be94cb582305c2d1607b919db0bbcc28cabe39457b63b14bac9fe6c",
                "MacAddress": "02:42:ac:12:00:03",
                "IPv4Address": "172.18.0.3/16",
                "IPv6Address": ""
            },
            "d6ec6fee21b00cd64ca9616f6d125a56a32c135ae41875e4bc174be376b103b0": {
                "Name": "alp-lab1",
                "EndpointID": "e54bec00bb61b6c4abd4b7822dee5e91c8daec0e2dce590a213aa4a41e0e4fae",
                "MacAddress": "02:42:ac:12:00:02",
                "IPv4Address": "172.18.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {}
    }
]
```

5. Acesse o container `alp-lab1` e tente se comunicar com os outros containers através do nome de cada um:
```shell
$ docker exec -ti alp-lab1 sh
/ # ping -c 2 alp-lab2
PING alp-lab2 (172.18.0.3): 56 data bytes
64 bytes from 172.18.0.3: seq=0 ttl=64 time=0.087 ms
64 bytes from 172.18.0.3: seq=1 ttl=64 time=0.129 ms

--- alp-lab2 ping statistics ---

/ # ping -c 2 alp-lab3
ping: bad address 'alp-lab3'
```
Perceba que o `alp-lab1` só consegue chegar até o `alp-lab2` pelo motivo de ambos estarem na mesma rede `lab-net`. Já o `alp-lab3` não é alcançável pois está numa rede diferente da `lab-net`.

6. Por via das dúvidas, conecte ao `alp-lab3` e tente chegar ao container `alp-lab1` ou `alp-lab2`:
```shell
$ docker exec -ti alp-lab3

/ # ping -c 2 alp-lab2
ping: bad address 'alp-lab2'
/ # 
/ # ping -c 2 alp-lab1
ping: bad address 'alp-lab1'
```

### Incluir containers em redes

Seguindo a infraestrutura anterior, vamos incluir o container `alp-lab3` na rede `lab-net` sem precisar recriá-lo.

1. Conecte o container `alp-lab3` na rede `lab-net`:
```shell
$ docker network connect lab-net alp-lab3
```

2. Agora acesse o container `alp-lab3` e tente alcançar os containers `alp-lab1` e `alp-lab2`:
```shell
$ docker exec -ti alp-lab3 sh            
/ # ping -c 2 alp-lab1
PING alp-lab1 (172.18.0.2): 56 data bytes
64 bytes from 172.18.0.2: seq=0 ttl=64 time=0.087 ms
64 bytes from 172.18.0.2: seq=1 ttl=64 time=0.146 ms

--- alp-lab1 ping statistics ---
2 packets transmitted, 2 packets received, 0% packet loss
round-trip min/avg/max = 0.087/0.116/0.146 ms

/ # ping -c 2 alp-lab2
PING alp-lab2 (172.18.0.3): 56 data bytes
64 bytes from 172.18.0.3: seq=0 ttl=64 time=0.181 ms
64 bytes from 172.18.0.3: seq=1 ttl=64 time=0.144 ms

--- alp-lab2 ping statistics ---
2 packets transmitted, 2 packets received, 0% packet loss
round-trip min/avg/max = 0.144/0.162/0.181 ms
```

3. Acesse os containers `alp-lab1` e  `alp-lab2` para testar a conexão com o `alp-lab3`:
```shell
$ docker exec -ti alp-lab2 sh
/ # 
/ # ping -c 2 alp-lab3
PING alp-lab3 (172.18.0.4): 56 data bytes
64 bytes from 172.18.0.4: seq=0 ttl=64 time=0.161 ms
64 bytes from 172.18.0.4: seq=1 ttl=64 time=0.142 ms

--- alp-lab3 ping statistics ---
2 packets transmitted, 2 packets received, 0% packet loss
round-trip min/avg/max = 0.142/0.151/0.161 ms
/ # exit

$ docker exec -ti alp-lab1 sh
/ # 
/ # ping -c 2 alp-lab3
PING alp-lab3 (172.18.0.4): 56 data bytes
64 bytes from 172.18.0.4: seq=0 ttl=64 time=0.130 ms
64 bytes from 172.18.0.4: seq=1 ttl=64 time=0.055 ms

--- alp-lab3 ping statistics ---
2 packets transmitted, 2 packets received, 0% packet loss
round-trip min/avg/max = 0.055/0.092/0.130 ms
```

Dessa forma conseguimos conectar o container `alp-lab3` que inicialmente não fazia parte da rede `lab-net` para comunicação com os containers `alp-lab1` e `alp-lab2` com sucesso.
