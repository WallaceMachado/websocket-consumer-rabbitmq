# websocket-consumer-rabbitmq

> status:	🚧 api-bank-transfers 🚀 em construção..  🚧

Um serviço, consome os eventos gerados pela [PERSON-API](https://github.com/WallaceMachado/challenge-go-rabbitmq) em uma fila do rambbitmq, e atualiza
um websocket para exibição dos mesmos.



## Indice

* <p><a href="#pré-requisitos">Pré Requisitos</a> </p>
* <p><a href="#iniciando-projeto">Iniciando Projeto</a></p>
* <p><a href="#variáveis-de-ambiente">Variáveis de Ambiente</a></p>
* <p><a href="#rotas">Rotas</a></p>
* <p><a href="#controle-de-versão">Controle de versão</a></p>
* <p><a href="#testes">Testes</a></p>
* <p><a href="#autor">Autor</a></p>




## Pré Requisitos

Antes de começar, você precisará ter as seguintes ferramentas instaladas em sua máquina:
* [Git](https://git-scm.com)

Para rodar via docker
* [Docker](https://docs.docker.com/)

Para rodar Local
* [Go](https://golang.org/) versão 1.16.7
* [RabbitMQ](https://www.rabbitmq.com/)

Além disso, é bom ter um editor para trabalhar com o código como: [VSCode](https://code.visualstudio.com/)



## Iniciando Projeto 

### Local

Deverá criar/rodar rabbitmq com as configurações informadas no arquivo ``` .env ```.  
Altere a varável de ambiente RABBITMQ_DEFAULT_HOST  para localhost
 

```bash
# Clone este repositório
$ git clone https://github.com/WallaceMachado/websocket-consumer-rabbitmq.git

# Acesse a pasta do projeto no terminal / cmd
$ cd websocket-consumer-rabbitmq

# Instale as dependências e rode o projeto
$ go run main.go

# Server is running
```


### Docker

```bash
# Clone este repositório
$ git clone https://github.com/WallaceMachado/websocket-consumer-rabbitmq.git

# Acesse a pasta do projeto no terminal / cmd
$ cd websocket-consumer-rabbitmq

# Instale as dependências e rode o projeto
$ docker-compose up -d --build

```
#### Configuração RabbitMQ
* Acessar via navegador http://localhost:15672, 
* Efetuar login com usuário e senha do arquivo ``` .env ```
* Criar nova exchange com o mesmo nome da variável de ambiente RABBITMQ_EXCHANGE_PERSON
* Criar nova queue com o mesmo nome da variável de ambiente RABBITMQ_QUEUE_PERSON
* Fazer o bind da queue com a exchange

## Variáveis de Ambiente

Após clonar o repositório, renomeie o ``` .env.example ``` no diretório raiz para ``` .env ``` e atualize com suas configurações.


| Chave  |  Descrição  | Predefinição  |
| :---: | :---: | :---: | 
|  PORT |  Número da porta em que o aplicativo será executado. | 5001  |
|  RABBITMQ_DEFAULT_USER | Usuário RabbitMQ.  |  -   |
|  RABBITMQ_DEFAULT_PASS | Senha RabbitMQ.  |  -   |
|  RABBITMQ_DEFAULT_VHOST | virtual host do RabbitMQ.  |  /   |
|  RABBITMQ_DEFAULT_HOST | host do RabbitMQ.  |  rabbit   |
|  RABBITMQ_DEFAULT_PORT | Porta em que o RabbitMQ está sendo excecutado  |  5672   |
|  RABBITMQ_EXCHANGE_PERSON | Exchanche do RabbitMQ  |  person_ex   |
|  RABBITMQ_QUEUE_PERSON  | Fila do RabbitMQ  |  person_queue   |

## Controle de versão
Para contrele de versão, foi inserida a versão ``` v1 ``` após o  ``` host ```

```
GET http://localhost:5001/api/v1/

```

 
## Rotas

| Rotas  |  HTTP Method  | Params  |  Descrição  | 
| :---: | :---: | :---: | :---: |
|  / |  GET |  -  | vizualizar, via navegagor, os eventos recebidos via websocket |
|  /person-api|  GET | - | connect-se via websocket com o serviço que consome a fila do rabbitmq das operações feitas na person-api |



## Testes
Para executar os testes :

### Local

Altere a varável de ambiente RABBITMQ_DEFAULT_HOST  para localhost

```bash

  # Acesse a pasta do projeto no terminal / cmd
  $ cd challenge-go-rabbitmq
  
  # Rode os testes
  $ go test ./...
  
```


### Docker


```bash

    # Acesse a pasta do projeto no terminal / cmd
    $ cd challenge-go-rabbitmq

    # Instale as dependências e rode o projeto
    $ docker-compose up --build

    # Pegar o CONTNAINER ID  da imagem github.com/wallacemachado/challenge-go-rabbitmq
    $ docker ps
    # Exemplo de retorno:
    # CONTAINER ID      IMAGE                                             COMMAND                  
    # xxxxxxxxxxxxxxx   github.com/wallacemachado/websocket-consumer-rabbitmq      "go run main.go"         0.0.0.0:5003->5000/tcp, :::5003->5003 ....  
    # yyyyyyyyyyyyyyy   github.com/wallacemachado/websocket-consumer-rabbitmq   "docker-entrypoint.s…"   0.0.0.0:5432->5432/tcp, :::5432->5432 ....

    # acesse o contaniner
    $ docker exec -it <CONTAINER ID> /bin/bash
    # Exemplo:
    # docker exec -it xxxxxxxxxxxxxxx /bin/bash

    #root@xxxxxxxxxxxxxxx:/go/src/app# 
    #rode os testes
    $ go test ./...

```



## Autor


Feito com ❤️ por [Wallace Machado](https://github.com/WallaceMachado) 🚀🏽 Entre em contato!

[<img src="https://img.shields.io/badge/linkedin-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" />](https://www.linkedin.com/in/wallace-machado-b2054246/)
