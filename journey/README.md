packages:
goapi-gen
sqlc
goqu*
squirrel*
pgx
tern


# Desenvolvimento de API RESTful com FastAPI, Docker, e PostgreSQL

## Sumário

1. [Descrição do projeto](#descrição-do-projeto)

2. [Descrição técnica](#descrição-técnica)

    2.1 [Tecnologias utilizadas](#tecnologias-utilizadas)
    
    2.2 [Principais funcionalidades da API](#principais-funcionalidades-da-api)

3. [Destaques do desenvolvimento](#destaques-do-desenvolvimento)

4. [Estrutura de pastas](#estrutura-de-pastas)

5. [Como rodar o projeto](#como-rodar-o-projeto)

    5.1 [Instalação](#instalação)

    5.2 [Rodando banco de dados](#rodando-banco-de-dados)

    5.3 [Rodando o servidor de desenvolvimento](#rodando_o-servidor-de-desenvolvimento)
    
    5.4 [Testes](#testes)

    5.4 [Conteinerização](#conteinerização)

    5.5 [Migrações](#migrações)

6. [Conclusão](#conclusão)


## Descrição do projeto

Esta é uma API genérica, criada com base em um "Travel Manager", criado pela Rocketseat no NLW de Julho/2024. Essa API está disponibilizada aqui para servir como fonte de estudos para mim e para a comunidade, portanto será constantemente atualizada. Sintam-se livres para enviar sugestões e comentários, o debate é o mais importante.

O objetivo principal é criar um sistema eficiente que aborde os principais pontos do desenvolvimento de um CRUD em GoLang que com o tempo vá se atualizando e se tornando mais eficiente.

## Descrição técnica

### Tecnologias Utilizadas

- **Go:** Linguagem de programação eficiente e robusta, ideal para desenvolvimento de serviços web escaláveis.
- **Postgresql:** Sistema de gerenciamento de banco de dados relacional avançado e de alta performance.
- **pgAdmin:** Ferramenta de administração e desenvolvimento de banco de dados Postgresql.
- **Mailpit:** Solução simples e eficiente para teste de envio de emails em ambiente de desenvolvimento.
- **Docker & Docker Compose:** Ferramentas essenciais para criar ambientes de desenvolvimento consistentes e facilmente replicáveis.

### Dependências Principais

**[GoAPI Gen](https://github.com/discord-gophers/goapi-gen/releases/tag/v0.3.0)**
  - Gerador de código Go a partir de especificações OpenAPI, facilitando a criação de APIs RESTful.

**[OpenAPI](https://github.com/getkin/kin-openapi/releases/tag/v0.126.0)**
  - Biblioteca para análise e manipulação de documentos OpenAPI, essencial para garantir a conformidade da API com os padrões.

**[Chi](https://github.com/go-chi/chi/releases/tag/v5.1.0)**
  - Framework leve e flexível para roteamento em Go, ideal para construir APIs HTTP.

**[Chi Render](https://github.com/go-chi/render/releases/tag/v1.0.3)**
  - Pacote complementar ao Chi para simplificar o processo de renderização de respostas HTTP.

**[Validator](https://github.com/go-playground/validator/v10/realeases/tag/v10.22.0)**
  - Biblioteca de validação de dados rica em recursos, utilizada para garantir a integridade dos dados de entrada na API.

**[Google UUID](https://github.com/google/uuid/releases/tag/v1.6.0)**
  - Gerador de UUIDs (Identificadores Únicos Universais), utilizado para criação de identificadores únicos e seguros.

**[pgx.v5](https://github.com/jackc/pgx/releases/tag/v5.6.0)**
  - Driver e biblioteca de conexão para Postgresql em Go, conhecido por sua performance e flexibilidade.

**[GUtils](https://github.com/phenpessoa/gutils)**
  - Coleção de utilitários úteis para desenvolvimento em Go, facilitando várias tarefas comuns.

**[GoMail](github.com/wneessen/go-mail/realeases/tag/v0.4.2)**
  - Biblioteca para envio de emails em Go, utilizada para implementar funcionalidades de notificação por email na API.

**[Zap Logger](https://go.uber.org/zap/realease/tag/v1.27.0)**
  - Logger de alta performance para Go, utilizado para registro eficiente de logs e diagnósticos.


### **Principais funcionalidades da API:**
- **CRUD Completo**: Operações de criação, leitura, atualização e exclusão de dados.
- **Geração de código**: Implementação de pacotes que gerem código Go a partir de especificações da OpenAPI.
- **Documentação Automática**: Utilização do OpenAPI para geração automática da documentação da API via Swagger.
- **Validação de Dados**: Validação rigorosa de entradas utilizando Validator.
- **Migrações de Banco de Dados**: Gerenciamento de mudanças e comunicações com o banco de dados utilizando SQLC.
- **AdminPage**: PgAdmin para gerar uma página para administrar o PostgreSQL

## **Destaques do Desenvolvimento:**
- **Desempenho**: Adoção de práticas para garantir alta performance, como consultas otimizadas e uso de caching.
- **Escalabilidade**: Arquitetura projetada para escalar horizontalmente com facilidade.
- **Segurança**: Implementação de padrões de segurança, incluindo sanitização de inputs e prevenção contra ataques comuns como SQL Injection e XSS.
- **Testes Automatizados**: Cobertura abrangente com testes unitários e de integração para garantir a confiabilidade do sistema.
- **Monitoramento e Logging**: Integração com ferramentas de monitoramento e logging para rastreamento e diagnóstico em tempo real.


## Estrutura de pastas e arquivos
- **cmd:**
    - **/journey/***journey.go*: Arquivo *main*. É ele quem irá gerar o binário para rodar o servidor da aplicação
- **internal/:** Pasta principal do projeto
   - **api/:** Diretório contendo a lógica da API assim como as suas especificações.
        - **spec/:** Diretório de especificações OpenAPI
            - */journey.gen.spec.go:* Arquivo gerado com toda a estrutura da API, métodos que devem ser implementados estão descritos na interface `ServerInterface` 
            - */journey.spec.json:* Arquivo JSON que servirá de orientação para gerar a estrutura da API.
        - */api.go:* Arquivo onde a lógica dos endpoints será implementada.
    - **mailer/mailpit/:** Diretório para configuração do servidor mailer local.
        - */mailpit.go:* Configuração do mailer escolhido.
    - **pgstore/:** Diretório do driver de conexão do PostgreSQL, aqui serão gerados vários arquivos de código.
        - **migrations/:** Arquivos descrevendo as migrações que serão aplicadas ao Banco de Dados utilizando o Tern.
            - */tern.conf:* Descrição das configurações do Tern, como as variaveis de ambiente.
        - **queries/:** Descrição das instruções que serão aplicadas no banco de dados em cada método.
            - */queries.sql:* Instruções em formato SQL.
        - */models.go:* Descrição dos modelos que serão levados ao banco de dados.
        - */sqlc.yaml:* Especificações sobre como o SQLC trabalhará, quais ferramentas, tags, drivers e até quais tipos serão sobrescritos.
- *compose.yaml:* Arquivo para orquestração dos containers de: aplicação, banco de dados, mailpit e pgadmin.
- *compose-db.yaml:* Arquivo inicialização somente do container do banco de dados.
- *Dockerfile:* Arquivo descrevendo as camadas da imagem Docker a ser buildada.
- *go.mod & go.sum:* Gerenciadores de dependências do projeto.
- *gen.go:* Arquivo de diretrizes do `go generate`.
- *Makefile* Arquivo com tasks úteis para o projeto.


## CLI do projeto
### Pré-requisitos:
1. (Git)[]

2. (GNU Make)[]

3. (Docker e plugin Docker Compose)[]


### Instalação

Copie o repositório
```bash
git clone https://github.com/rmndvngrpslhr/travel_tracker_go
```

Instale as dependências
```bash
make go
```

### Rodando o banco de dados
Suba somente o banco de dados com o Docker Compose em modo deattach na porta `5432`
```bash
make db
```

### Rodando os servidores de desenvolvimento:
Pode buildar a aplicação e rodá-la na porta `8081`, junto com o banco de dados na `5432`, o pgadmin na `8081` e o mailpit na `1025`
```bash
make api-build-run
```

Para rodar todos os serviços com as imagens previamente construidas
```bash
make run-api
```

### Encerrando
Para encerrar as atividades do docker compose
```bash
make services-stop
```

### Operações com o banco de dados:
Pode instruir a geração de código do SQLC baseada no arquivo `.yaml` da pasta `/internal/pgstore`
```bash
make sqlc-generate
```

Para criar novas migrações com o Tern
```bash
make new-migration MIGRATION_NAME=${MIGRATION_NAME}
```

Para aplicar as migrações
```bash
make migrations
```

> Para mais detalhes sobre o que cada comando "make" faz, dê uma olhada no arquivo [Makefile](Makefile)


## Conclusão

Espero que este projeto sirva como uma base sólida para o estudo e compreensão das tecnologias envolvidas. Sinta-se à vontade para explorar, modificar e expandir o Travel Tracker conforme seu interesse e curiosidade.

Obrigado por ler e aproveitar o projeto!

> Caso precise de mais alguma informação ou ajuste, estou à disposição no e-mail: avgsolheiro@gmail.com!