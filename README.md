# go-gateway

Features:

- Cada cliente da gateway terá uma account e uma secret
- Cada transção será realizada via API REST
- Serviço processador de pagamentos
- Quando uma transação é enviada via API, seus dados serão encaminhados para o processador de pagamentos
- painel onde poderemos acompanhar as transações sendo processadas em tempo real
- Comunicação síncrona e assincrona
- tudo é uma simulação para que possamos fazer os microsserviços se comunicarem
  faremos com que os sistemas consigam conversar de forma síncrona e assíncrona para que possamos explorar conceitos e
  tecnologias utilizadas em grandes empresas que necessitam de performance e resiliencia

Detalhes:

- Não misturar responsabilidade de receber as transações com o processamento das mesmas
- Ao submetermos as transações para o microsserviço de processamento, precisaremos de resiliencia para evitar que
  qualquer dado seja perdido
- O microsserviço que recebe a solicitação enviará os dados da transação para o kafka.
- do outro lado, o microsserviço responsavel por realizar a transação fará a leitura do kafka.
- ao processar, o retorno do resultado será publicado em um outro tópico do kafka.

Regras do processador de pagamentos:

- Uma transação terá apenas dois status: "approved" ou "rejected"
- O valor mínimo para cada transação é de $1,00
- O valor máximo para que uma transação seja aprovada é de $1000
- Toda transação entre os valores de $1 e $1000 sempre será aprovadas
- Para que uma transação seja aprovada, os dados do cartão de crédito deve ser válidos

* Caso de uso: "process transaction"

1. Receberá os dados de uma transação
2. criará uma transa~~ao
3. adicionará o cartão de credito a essa transação
    - se o cartão for invalido:
        - os dados da transação serão inseridos no banco de dados com o status=rejected contendo a mensadem do erro
        - a transação será publicada no kafka

4. caso a transação não seja aprovada:
    - os dados da transação serão inseridos no banco de dados com o status=rejected contendo a a mensagem do erro
    - a transação será publicada no kafka

5. caso a transação seja aprovada:
    - os dados da transação serao inseridos no banco de dados com o status=approved
    - a transacao sera publicada no kafka

* Adaptadores (software que vai falar com o mundo externo):
    - apache kafka
        - consumo de mensagens
        - produção das mensagens
    - banco de dados
        - conexao com o banco de dados
        - insercao da transacao no banco
        - SQLite/MySQL
        - repository
    - presenter
        - definira o padrao da mensagem a ser enviada via apache kafka

- Tecnologias a serem utilizadas
  - Frontend Painel: Next.js
  - Backend Painel: Nest.js
  - Microsserviço de processamento: Goland 
  - Apache Kafka
  - Prometheus e Grafana
  - Docker e Kubernetes

## Como Compilar e Rodar o Projeto

1. Clone o repositório:
   ```sh
   git clone https://github.com/patrickcmserrano/go-gateway.git
   cd go-gateway
   ```

2. Inicialize o módulo Go:
   ```sh
   go mod init github.com/patrickcmserrano/go-gateway
   ```

3. Compile o projeto:
   ```sh
   go build -o bin/go-gateway main.go
   ```

4. Configure as variáveis de ambiente necessárias:
   ```sh
   export DB_HOST=localhost
   export DB_USER=usuario
   export DB_PASSWORD=senha
   export DB_NAME=go_gateway
   export KAFKA_BROKER=localhost:9092
   ```

5. Execute o projeto:
   ```sh
   ./bin/go-gateway
   ```

## Como Rodar o Projeto com Docker

1. Clone o repositório:
   ```sh
   git clone https://github.com/patrickcmserrano/go-gateway.git
   cd go-gateway
   ```

2. Construa e execute os containers com `docker-compose`:
   ```sh
   docker-compose up -d
   ```

3. Se você atualizar as dependências no container, reconstrua e execute os containers:
   ```sh
   docker-compose up -d --build
   ```

4. Acesse o container em execução:
   ```sh
   docker exec -it go-gateway-app-1 bash
   ```

## Como Rodar os Testes

1. Execute todos os testes:
   ```sh
   go test ./...
   ```

2. Execute testes específicos:
   ```sh
   go test -v ./path/to/your/package
   ```

