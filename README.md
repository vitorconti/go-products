# Go Products API
This is a sample Go application for managing products. It uses a MySQL database to store the products and RabbitMQ for sending messages between components.

## Requirements
- Docker (https://docs.docker.com/engine/install/)
- Docker Compose (https://docs.docker.com/compose/install/)
- Go 1.19 or higher (https://golang.org/dl/)

## Running the Application
Clone the repository:

```sh

$ git clone https://github.com/vitorconti/go-products.git
$ cd go-products-api
Start the Docker containers:
```
with this following command you will create the necessary containers to run the application
```sh

$ docker-compose up -d

```

At the first time that you run this following command, will be created the products table in the database and the seeder of this table

```go

$ go run cmd/main.go cmd/wire_gen.go

```
 inside of the main folder you can found a post collection with the endpoints and requests examples
 you can found the .env of the application inside of the project feel free to change and test


Usage
The API has the following endpoints:

GET /products
Returns a list of all products.

GET /products/:id
Returns a single product by ID.

POST /products
Creates a new product. The request body should contain JSON data with the following properties:
```json
{
  "name": "Product Name",
  "description": "Product Description",
  "price": 10.99
}
```
PUT /products/:id
Updates an existing product. The request body should contain JSON data with the following properties:

```json
{
  "name": "Product Name",
  "description": "Product Description",
  "price": 10.99
}
```
DELETE /products/:id
Deletes an existing product.

# Troubleshooting
If you encounter any issues with the application, try the following steps:

- Make sure Docker and Docker Compose are installed correctly.
- Check if the required ports (8080, 3306, 5672, 15672) are available on your system.
- Restart Docker Compose with docker-compose down followed by docker-compose up.
- If the MySQL container fails to start, delete the mysql/data directory and restart Docker Compose.

License
This application is open source software licensed under the MIT License. See the LICENSE file for more information.
