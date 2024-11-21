
# Carrinho de Compras - Backend

## Descrição do Projeto

Este é o backend da aplicação de **Carrito de Compras**, desenvolvido em **Golang**. O sistema permite a gestão de produtos no carrinho, como adicionar, remover, alterar quantidades e calcular totais. A aplicação consome dados da API externa do Mercado Livre para listar produtos.

---

## Funcionalidades

- **Adicionar produtos ao carrinho**
- **Remover produtos do carrinho**
- **Exibir o número total de produtos no carrinho**
- **Exibir o preço total do carrinho**
- **Consultar produtos via API Mercado Livre**
- **Calcular subtotal por produto**

---

## Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Framework HTTP**: Gin Gonic
- **Integração com API**: Consumo da API do Mercado Livre
- **Base de Dados**: MySQL
- **Docker**
- **Testes**: Postman para validação dos endpoints

---

## Endpoints Disponíveis

### **1. Listar Produtos**
- **Método**: `GET`
- **Endpoint**: `/products`
- **Parâmetros**: 
  - `q` (opcional): Filtro de busca por palavra-chave.
- **Exemplo de Requisição**:
  ```bash
  GET http://localhost:8080/products?q=notebook
  ```
- **Resposta**:
  ```json
  {
    "products": [
        {
            "id": "MLB5151418202",
            "title": "Notebook Gamer Dell G15-i1300-d35p 15.6  Fhd 13ª Geração Intel Core I5 16gb 512gb Ssd Nvidia Rtx 3050 Linux",
            "price": 5399.99,
            "thumbnail": "http://http2.mlstatic.com/D_759556-MLU79227331595_092024-I.jpg",
            "permalink": "https://www.mercadolivre.com.br/notebook-gamer-dell-g15-i1300-d35p-156-fhd-13-geraco-intel-core-i5-16gb-512gb-ssd-nvidia-rtx-3050-linux/p/MLB40690557#wid=MLB5151418202&sid=unknown"
        },
        {
            "id": "MLB5061213442",
            "title": "Notebook Vaio Fe15 Intel Core I5 - 1235u, 8gb Ram, 256gb Ssd, Windows 11 Home, Tela 15,6  Full Hd - Vjfe54f11x-b0111h",
            "price": 2499,
            "thumbnail": "http://http2.mlstatic.com/D_764214-MLA54863992522_042023-I.jpg",
            "permalink": "https://www.mercadolivre.com.br/notebook-vaio-fe15-intel-core-i5-1235u-8gb-ram-256gb-ssd-windows-11-home-tela-156-full-hd-vjfe54f11x-b0111h/p/MLB22530427#wid=MLB5061213442&sid=unknown"
        },
    ]
  }
  ```

---

### **2. Adicionar Produto ao Carrinho**
- **Método**: `POST`
- **Endpoint**: `/cart`
- **Corpo da Requisição**:
  ```json
  {
  "product_id": "MLB789123",
  "title": "Mouse Gamer Logitech G502",
  "quantity": 1,
  "price": 250.0
  }
  ```
- **Resposta**:
  ```json
  {
    "message": "Item adicionado ao carrinho com sucesso"
  }
  ```

---

### **3. Remover Produto do Carrinho**
- **Método**: `DELETE`
- **Endpoint**: `/cart/:id`
- **Exemplo de Requisição**:
  ```bash
  DELETE http://localhost:8080/cart/1
  ```
- **Resposta**:
  ```json
  {
    "message": "Item removido com sucesso"
  }
  ```
---

### **4. Exibir Carrinho**
- **Método**: `GET`
- **Endpoint**: `/cart`
- **Resposta**:
  ```json
  {
    "cart": [
        {
            "id": 2,
            "product_id": "MLB789123",
            "title": "Mouse Gamer Logitech G502",
            "quantity": 1,
            "price": 250,
            "created_at": "2024-11-20T22:27:13-03:00",
            "updated_at": "2024-11-20T22:27:13-03:00"
        },
        {
            "id": 3,
            "product_id": "MLB789123",
            "title": "Mouse Gamer Logitech G502",
            "quantity": 1,
            "price": 250,
            "created_at": "2024-11-20T22:41:52-03:00",
            "updated_at": "2024-11-20T22:41:52-03:00"
        },
        {
            "id": 4,
            "product_id": "MLB789123",
            "title": "Mouse Gamer Logitech G502",
            "quantity": 1,
            "price": 250,
            "created_at": "2024-11-21T01:42:49-03:00",
            "updated_at": "2024-11-21T01:42:49-03:00"
        }
    ]
  }
  ```

---

## Configuração do Projeto

### Pré-requisitos
- **Golang** (versão 1.18 ou superior)
- **Docker** (opcional)
- **Ferramentas de Teste**: Postman

---

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/jessikasousa/shopping-cart-backend.git
   cd shopping-cart-backend
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Execute o servidor:
   ```bash
   go run cmd/main.go
   ```

4. Acesse o servidor em `http://localhost:8080`.

---

## Próximos Passos

- Implementar autenticação JWT.
- Configurar testes automatizados.

---

## Autor
**Jéssika Sousa**
GitHub: [jessikasousa](https://github.com/jessikasousa)
