# GraphQL

GraphQL é uma linguagem de consulta de dados e uma runtime para execução dessas consultas, desenvolvida pelo Facebook em 2012 e posteriormente liberada como código aberto em 2015. Ela oferece uma abordagem mais eficiente e flexível para obter e manipular dados em comparação com abordagens tradicionais como REST.

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
