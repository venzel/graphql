# GraphQL

> Este guia foi elaborado por **Enéas Almeida** com o principal objetivo de facilitar os repasses de informações à equipe.

## O que é o GraphQL

<p align="center"><img src="./media/graphql.svg" width="300" /></p>

GraphQL é uma linguagem de consulta de dados e uma runtime para execução dessas consultas, desenvolvida pelo Facebook em 2012 e posteriormente liberada como código aberto em 2015. Ela oferece uma abordagem mais eficiente e flexível para obter e manipular dados em comparação com abordagens tradicionais como REST.

## Onde é ideal utilizar GraphQL

-   **Aplicações com múltiplos clientes**: Quando você tem diferentes clientes (por exemplo, aplicativos móveis, aplicativos da web, dispositivos IoT);

-   **Microserviços**: Em ambientes de microservices, onde há muitos serviços independentes responsáveis por diferentes partes de uma aplicação;

-   **BFF**: Uma excelente escolha para o padrão arquitetural BFF, fornecendo flexibilidade de consultas, redução de chamadas de APIs e desacoplamento cliente-servidor.

## Links importantes

-   [GQLGen](https://gqlgen.com/) - Biblioteca em Go para construir servidores GraphQL rapidamente.

### Passo 1 - Incializa a estrutra de pastas com os pacotes necessários

```bash
go mod init graphql

printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

go mod tidy
```

### Passo 2 - Inicializa o projeto com o GQLGen

```bash
go run github.com/99designs/gqlgen init

go mod tidy
```

### Passo 3 - Inializa o servidor

```bash
go run server.go
```

### Passo 4 - Acessa o playground do GraphQL (Apenas verificação)

http://localhost:8080/

## Passo 5 - Altera o schema GraphQL pré definido

```graphql
# Em /graph/schema.graphqls
scalar DateTime

type Transaction {
    id: ID!
    amount: Float!
    accounts: [Account!]!
    createdAt: DateTime!
    updatedAt: DateTime!
}

type Account {
    id: ID!
    name: String!
    email: String!
    createdAt: DateTime!
    updatedAt: DateTime!
}

input NewTransaction {
    amount: Float!
    accountId: ID!
}

input NewAccount {
    name: String!
    email: String!
}

type Query {
    transactions: [Transaction!]!
    accounts: [Account!]!
}

type Mutation {
    createTransaction(input: NewTransaction!): Transaction!
    createAccount(input: NewAccount!): Account!
}
```

Assim que alterado o schema como demonstrado no código acima, reconfigura o schema através do GQLGen com o comando abaixo:

```bash
go run github.com/99designs/gqlgen generate
```

⚠️ **Atenção**: após gerar o novo schema, lembrar de apagar os resolvers antigos em schema.resolvers.go

### Passo 6 - Testando a consulta no Playground

⚠️ **Atenção**: Antes de testar, implementar os métodos, pois, se não implementado, retorna um panic.

http://localhost:8080/

<img src="./media/p1.png" />

```graphql
# Cria a conta
mutation createAccount {
    createAccount(input: { name: "Tiago Rizzo", email: "tiago@gmail.com" }) {
        id
        name
        email
    }
}
```

```graphql
# Lista as transacoes
query queryTransactions {
    transactions {
        id
        amount
    }
}
```

## Sqlite3

### Instalação Sqlite3

```bash
# Instalação
sudo apt install sqlite3
# Versão
sqlite3 --version
```

### Comandos Sqlite3

```bash
# Acessa o banco
sqlite3 data.db
# Cria tabela accounts
sqlite> create table accounts (id string PRIMARY KEY, name string, email string);
# Cria a tabela transactions
sqlite> create table transactions (id string PRIMARY KEY, amount decimal, accountId string);
# Lista os dados da tabela
sqlite> select * from accounts;
# Para sair
sqlite> .quit
```

<details>
<summary>Mais comandos do Sqlite3</summary>

```bash
# Deleta todos os registros
sqlite> DELETE FROM accounts;
# Dropa a tabela
sqlite> DROP TABLE accounts;
# Insere um registro na tabela accounts
sqlite> INSERT INTO accounts (id, name, email) VALUES ('xx0011', 'tiago', 'tiago@gmail.com');
# Insere um registro na tabela transactions
sqlite> INSERT INTO transactions (id, amount, accountId) VALUES ('kk0033', 33.20, 'xx0011');
```

</details>

<div>
  <img align="left" src="https://imgur.com/k8HFd0F.png" width=35 alt="Profile"/>
  <sub>Made with 💙 by <a href="https://github.com/venzel">Enéas Almeida</a></sub>
</div>
