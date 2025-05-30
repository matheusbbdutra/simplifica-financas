# SimplificaFinancas

Sistema de gerenciamento financeiro pessoal, desenvolvido em Go, com arquitetura modular e autenticação JWT.

## Funcionalidades

- Cadastro de usuários (`/api/user/register`)
- Login de usuários com JWT (`/api/user/login`)
- Atualização de dados do usuário (`/api/user/update`)
- Documentação Swagger disponível em `/docs`

## Estrutura do Projeto

```
.
├── main.go           # Ponto de entrada da aplicação
├── config/jwt/       # Chaves privadas/públicas para JWT
├── database.db       # Banco de dados SQLite
├── docs/             # Documentação Swagger (JSON/YAML)
├── internal/         # Código principal (domínios, casos de uso, handlers)
│   ├── app/          # Inicialização de módulos e rotas
│   ├── finances/     # Módulo financeiro (em desenvolvimento)
│   └── user/         # Módulo de usuários
├── pkg/utils/        # Utilitários (JWT, senha, validação)
├── Makefile          # Comandos úteis (testes, geração de chaves)
└── ...
```

## Como rodar

1. **Pré-requisitos:** Go 1.24+, SQLite3, [swag](https://github.com/swaggo/swag) (para gerar docs)
2. **Variáveis de ambiente:**  
   Copie `.env.exemple` para `.env` e ajuste conforme necessário.
3. **Gerar chaves JWT (se necessário):**
   ```sh
   make generate-jwt-keys
   ```
4. **Rodar a aplicação:**
   ```sh
   go run cmd/app/main.go
   ```
5. **Rodar testes:**
   ```sh
   make test
   ```

## Endpoints principais

### POST `/api/user/register`
Cria um novo usuário.

**Body:**
```json
{
  "email": "email@exemplo.com",
  "password": "senha1234"
}
```

### POST `/api/user/login`
Realiza login e retorna um token JWT.

**Body:**
```json
{
  "email": "email@exemplo.com",
  "password": "senha1234"
}
```

### POST `/api/user/update`
Atualiza dados do usuário autenticado.

**Headers:**
```
Authorization: Bearer <JWT>
```
**Body:**
```json
{
  "name": "Novo Nome",
  "email": "novoemail@exemplo.com",
  "password": "novasenha1234"
}
```

## Documentação

Acesse a documentação Swagger em `/docs` após iniciar o servidor.

---

Projeto em desenvolvimento. Contribuições são bem-vindas!