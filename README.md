# golang-sample
 Este repositório contém uma coleção de programas e bibliotecas Go que demonstram a linguagem, bibliotecas padrão e ferramentas.

<p align="center">
![go](/assets/go.jpg)
</p>

## Índice
1. [Projeto e Conteúdo](#Projeto-e-Conteudo)
2. [Swagger](#Swagger)
3. [GORM e SQL Server](#SQL-Server-e-JDBC)

## Projeto e Conteúdo

### Dependences
```bash
go get -u github.com/gorilla/mux
go get -u gorm.io/driver/sqlserver
go get -u gorm.io/gorm
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```
### Como executar essa aplicação?

 - Faça o download do zip ou clone o repositório Git.
 - Descompacte o arquivo zip (caso tenha baixado o .zip) no diretorio `$home\go\src\github.com\fabiodelgadopereira\`
 - Abra o diretório Prompt de Comando e digite:

```bash
go build
golang-sample.exe
```

A aplicação deverá estar disponivel em seu navegador no endereço: http://localhost:8000/swaggerui/

![swagger](/assets/swagger.png)

## Swagger

O Swagger é uma aplicação open source que auxilia os desenvolvedores a definir, criar, documentar e consumir APIs REST;
É composto de um arquivo de configuração, que pode ser definido em YAML ou JSON;
Fornece ferramentas para: auxiliar na definição do arquivo de configuração (Swagger Editor), interagir com API através das definições do arquivo de configuração (Swagger UI) e gerar templates de código a partir do arquivo de configuração (Swagger Codegen).

fonte: https://swagger.io/resources/webinars/getting-started-with-swagger/

## GORM e SQL Server

O Grails Object Relational Mapper (GORM) é um excelente produto ORM. O que começou como um Groovy DSL sobre o Hibernate e modelado com Grails evoluiu para uma persistência agnóstica, uma biblioteca Grails independente para trabalhar com bancos de dados.
O GORM usa design via API e um pouco da magia Groovy por ser especialmente amigável. Ele faz isso através de cinco mecanismos de consulta de dados muito poderosos:
 - Finders dinâmicos
 - Cláusulas WHERE
 - Critérios
 - HQL
 - SQL