# Blockchain Simples em Go

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![SHA-256](https://img.shields.io/badge/SHA--256-4CAF50?style=for-the-badge&logo=security&logoColor=white)




#### Este projeto implementa uma blockchain simples em Go, com foco em demonstrar os conceitos básicos de como funciona uma blockchain: encadeamento de blocos, hash, e geração de um novo bloco a cada intervalo de tempo.

## Funcionalidade

- ### O projeto cria uma blockchain simples onde:

    - Um bloco gênese (o primeiro bloco da cadeia) é criado.

    - Após a criação do bloco gênese, um novo bloco é gerado automaticamente a cada 3 segundos, com base no bloco anterior.

    - Cada bloco contém:

        - Um índice (posição do bloco na cadeia).
        - Um timestamp (data e hora de criação).
        - Dados fictícios (aqui representados como Block Data com um número sequencial).
        - O hash do bloco anterior, garantindo a integridade da cadeia.
        - O próprio hash do bloco, que é recalculado toda vez que um novo bloco é criado.

## Como Funciona?

- Bloco Gênese: O primeiro bloco é criado manualmente. Ele não tem um bloco anterior e não possui hash (é um caso especial na blockchain).
- Blocos Subsequentemente Criados: A cada 3 segundos, um novo bloco é criado, com base no bloco anterior.
- Hashing: O hash de cada bloco é calculado usando a função de hash SHA256. Esse hash é a "impressão digital" do bloco, garantindo sua integridade.
- Encadeamento de Blocos: Cada bloco aponta para o bloco anterior por meio de seu hash, garantindo que, se qualquer bloco for alterado, todos os blocos subsequentes serão invalidos.

## Como Usar
- ### Requisitos:
  - #### Ter Go instalado.

- ### Clone o repositório:

        git clone https://github.com/Rodrigo-Kelven/Blockchain-Example
        cd blockchain-simples-go

- ### Execute o código:

        go run main.go

##### Saída esperada: O programa começará exibindo o bloco gênese, seguido de novos blocos gerados automaticamente a cada 3 segundos, imprimindo detalhes como índice, timestamp, dados, hash do bloco anterior e o hash do bloco atual.

### Exemplo de Saída

        Index: 0
        Timestamp: 2025-06-09 09:55:42 +0000 UTC
        Data: Genesis Block
        Previous Hash: 
        Hash: 8c0d77c8507bc557f64e2e3fcd1f9ec12a58f7b944928eb3e93c9fd689d40d4e

        Index: 1
        Timestamp: 2025-06-09 09:55:45 +0000 UTC
        Data: Block Data 1
        Previous Hash: 8c0d77c8507bc557f64e2e3fcd1f9ec12a58f7b944928eb3e93c9fd689d40d4e
        Hash: ecb9f4f55a3221cf84cbf7cc8e44a79eae1a747b9636da4f287c6c41ac1251a8

        Index: 2
        Timestamp: 2025-06-09 09:55:48 +0000 UTC
        Data: Block Data 2
        Previous Hash: ecb9f4f55a3221cf84cbf7cc8e44a79eae1a747b9636da4f287c6c41ac1251a8
        Hash: 1c0d389bfda5b0a98813b983c34a1150a7a084cb0d7f6f825234fd35b69861f5

### Tecnologias Utilizadas

- Go (Golang): Linguagem utilizada para a implementação da blockchain.
- SHA-256: Função de hash criptográfico usada para gerar os hashes dos blocos.

### Contribuições

Se você deseja contribuir para este projeto, fique à vontade para criar pull requests ou relatar issues. Melhorias como persistência de dados, maior segurança, e otimizações de desempenho são sempre bem-vindas.

## Autores
- [@Rodrigo_Kelven](https://github.com/Rodrigo-Kelven)
