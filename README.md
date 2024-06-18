# Sistema de roteamento de mensagens RabbitMQ

Este repositório contém um sistema de roteamento de mensagens baseado em RabbitMQ com o intuito de mostrar de forma simplificada a criação de produtor(Producer), consumidor(Consumer), fila (Queue) e Exchange.

## Visão geral

Este projeto demonstra o uso do RabbitMQ para rotear mensagens com base em seu conteúdo. Os números de 0 a 10 são publicados em uma exchange e roteados para filas diferentes dependendo se são pares ou ímpares.

## Arquitetura
* Produtor: Publica números de 0 a 10 em uma exchange.
* Exchange: roteia mensagens para filas apropriadas com base no conteúdo da mensagem.
* Filas:
  * even_numbers: Recebe números pares.
  * odd_numbers: Recebe números ímpares.
* Consumidores: Escuta as filas e processa as mensagens.


## Installation
1. Clone the repository:
~~~bash
git clone https://github.com/henrique77/rabbitmq_producer_consumer.git
cd rabbitmq_producer_consumer
~~~
2. Instale dependências:
~~~bash
go mod tidy
~~~

## Configurando o RabbitMQ
1. Inicie o servidor RabbitMQ executando:
~~~bash
docker compose up -d
~~~
2. Acesse o servidor RabbitMQ:
```
http://localhost:15672/
```
user e password: guest

Obs.: Criar filas conforme a seção de arquitetura

## Executando Produtor e Consumidor
1. Execute o Produtor:
~~~bash
go run cmd/producer/main.go
~~~
Obs.: Nesse momento, se observar a aba **queue** do RabbitMQ verá 5 mensagens em cada fila criada.

2. Execute o Consumidor:
~~~bash
go run cmd/consumer/main.go
~~~
Obs.: Será consumida e impressa todas as mensagens que estavam nas filas

## Sobre mim

Sou Henrique Caires, desenvolvedor de software. Estou a disposição para dúvidas, esclarecimentos e sugestões. Me encontre no linkedin: [Henrique Caires](https://www.linkedin.com/in/henrique-caires)
