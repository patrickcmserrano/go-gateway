# Backlog do Projeto go-gateway

## Funcionalidades Implementadas

- Validação de número de cartão de crédito
  - Arquivo: [domain/entity/credit_card.go](domain/entity/credit_card.go)
  - Testes: [domain/entity/credit_card_test.go](domain/entity/credit_card_test.go)

- Validação de transações
  - Arquivo: [domain/entity/transaction.go](domain/entity/transaction.go)
  - Testes: [domain/entity/transaction_test.go](domain/entity/transaction_test.go)

- Caso de uso "process transaction"
  - Arquivo: [usecase/process_transaction/process_transaction.go](usecase/process_transaction/process_transaction.go)
  - Testes: [usecase/process_transaction/process_transaction_test.go](usecase/process_transaction/process_transaction_test.go)

- Repositório de transações (mock)
  - Arquivo: [domain/repository/mock/mock.go](domain/repository/mock/mock.go)

- Migrações de banco de dados
  - Arquivos: [adapter/repository/fixture/sql/1-transactions.up.sql](adapter/repository/fixture/sql/1-transactions.up.sql), [adapter/repository/fixture/sql/1-transactions.down.sql](adapter/repository/fixture/sql/1-transactions.down.sql)
  - Fixture: [adapter/repository/fixture/fixture.go](adapter/repository/fixture/fixture.go)

## Funcionalidades a Implementar

- [ ] Implementar repositório de transações com banco de dados real
  - Arquivo: [adapter/repository/transaction_repository_db.go](adapter/repository/transaction_repository_db.go)
  - Testes: [adapter/repository/transaction_repository_db_test.go](adapter/repository/transaction_repository_db_test.go)

- [ ] Integração com Apache Kafka para produção e consumo de mensagens
  - [ ] Adaptador de produção de mensagens
  - [ ] Adaptador de consumo de mensagens

- [ ] Painel para acompanhar transações em tempo real
  - [ ] Frontend: Next.js
  - [ ] Backend: Nest.js

- [ ] Configuração de Prometheus e Grafana para monitoramento

- [ ] Configuração de Docker e Kubernetes para deploy e orquestração

## Melhorias e Refatorações

- [ ] Adicionar mais testes unitários e de integração
- [ ] Melhorar a documentação do código
- [ ] Implementar tratamento de erros mais robusto
- [ ] Refatorar código para seguir melhores práticas de design de software

## Tarefas Técnicas

- [ ] Configurar CI/CD para o projeto
- [ ] Configurar ambiente de desenvolvimento com Docker
- [ ] Configurar ambiente de produção com Kubernetes

## Bugs Conhecidos

- Nenhum bug conhecido no momento
