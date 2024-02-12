# API Loja de Pets

Esta é uma API de exemplo para uma loja de pets, desenvolvida em Go, utilizando o framework Mux para roteamento e o banco de dados memDB.

## Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas em seu sistema:

- Go (versão 1.21 ou superior)
- Git
- Postman (para testar a API)

## Instalação e Execução

1. Clone este repositório:

```bash
git clone https://git@github.com:renatinhafront/pet-store.git
Navegue até o diretório do projeto:
bash
Copy code
cd pet-store
Instale as dependências:
bash
Copy code
go mod tidy
Execute a aplicação:
bash
Copy code
go run main.go
A API estará disponível em http://localhost:8080.

Endpoints
Listar todos os Pets
bash
Copy code
GET /pets
Retorna uma lista de todos os pets na loja.

Adicionar um novo Pet
bash
Copy code
POST /pets
Adiciona um novo pet à loja. Envie os dados do pet no corpo da solicitação no formato JSON.

Obter informações de um Pet específico
bash
Copy code
GET /pets/{id}
Retorna informações sobre o pet com o ID fornecido.

Atualizar informações de um Pet
bash
Copy code
PUT /pets/{id}
Atualiza as informações do pet com o ID fornecido. Envie os dados atualizados no corpo da solicitação no formato JSON.

Excluir um Pet
bash
Copy code
DELETE /pets/{id}
Remove o pet com o ID fornecido da loja.

Contribuição
Contribuições são bem-vindas! Se você encontrar algum problema ou tiver sugestões de melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.

Licença
Este projeto está licenciado sob a licença MIT. Consulte o arquivo LICENSE para obter mais detalhes.

Golang
Copy code

## Me encontre aqui:

[<img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" />](https://www.linkedin.com/in/renata-saraiva-santos/)

Muito Obrigada!!
