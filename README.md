# Autorizador

## Decisões técnicas e arquiteturais

**Golang** foi utilizado para desenvolver o projeto por trabalhar facilmente com a manipulação de stdin e stdout além de gerar executáveis de código para várias plataformas. Go também tem características gerais atraentes como simplicidade e performance que podem ser pontos chaves em uma gama de aplicações.

**Clean Architecture**  é utilizada como design para projeto por fornecer uma série de vantagens como: código facilmente testável, independência de frameworks, UI e banco de dados ou  qualquer agente externo. Sua divisão de camadas priorizando separação de responsabilidades tornando o software muito coeso e manutenível. 

## Dependências
- [uuid](https://github.com/google/uuid)
    - O pacote uuid gera e inspeciona UUIDs com base no [RFC 4122](https://datatracker.ietf.org/doc/html/rfc4122).
    - Utilizado como identificador único de entidades
- [gomock](https://github.com/golang/mock)
    - Gomock é uma estrutura de simulação para linguagem Go. 
    - Utilizado na geração de código que simula objetos para testes
- [testify](https://github.com/stretchr/testify)
    -  Fornece muitas ferramentas para testar que seu código se comportará como você deseja.
    - Utilizado para fazer asserções nos testes

## Executando projeto

2º Construíndo build da aplicação

	make build

3º Executando aplicação 

	./autorizador < operations

4º Executando testes

	make test

4º Executando testes de integração

	make test-integration

5° Verificação da cobertura de teste

	make test-cov

### Outros comandos

Gerando/Atualizando mocks
    
    make build-mocks

Download das dependencias

    make dependencies

Executando aplicação sem build

    make run