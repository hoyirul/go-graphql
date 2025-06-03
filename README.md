# Golang GraphQL
> A simple Golang GraphQL server with a basic schema and resolvers.

## Features
- Basic GraphQL schema with queries and mutations
- Simple resolvers for handling queries and mutations
- Example of using GraphQL with Golang

## Getting Started
### Prerequisites
- Go 1.18 or later
- A working Go environment

### Installation
1. Clone the repository:
   ```bash
    git clone github.com/hoyirul/go-graphql.git
    cd go-graphql
   ```
2. Install dependencies:
   ```bash
    go mod tidy
   ```
3. Run the server:
   ```bash
    go run main.go
   ```
4. Open your browser and navigate to `http://localhost:8080/query` to access the GraphQL playground.

### Usage
1. Get 99designs gqlgen:
   ```bash
    go get github.com/99designs/gqlgen@v0.17.74
   ```
2. Generate the GraphQL server code:
   ```bash
    go run github.com/99designs/gqlgen generate
   ```

### Example Queries
- **Get all products:**
```graphql
query {
  products {
    id
    name
    price
  }
}
```
- **Get a product by ID:**
```graphql
query {
  product(id: "1") {
    id
    name
    price
  }
}
```
- **Add a new product:**
```graphql
mutation {
  addProduct(name: "New Product", price: 19.99) {
    id
    name
    price
  }
}
```

# License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

# Authors
- [Mochammad Hairullah](https://linkedin.com/in/madhai/)