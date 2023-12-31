# Bitcoin Service API

This project provides an API for interacting with Bitcoin addresses and transactions, allowing you to obtain address and transaction details, check balances, and assemble UTXO (Unspent Transaction Output) for Bitcoin transactions.

## Functionalities

The API offers the following endpoints:

### 1. Get Details by Bitcoin Address

- **Endpoint**: `/address/{address}`
- **Method**: GET
- **Description**: Returns detailed details for a specific Bitcoin address.
- **Parameters**:
    - `address`: Bitcoin address (path parameter).

### 2. Get Balance by Bitcoin Address

- **Endpoint**: `/balance/{address}`
- **Method**: GET
- **Description**: Returns balance information for a specific Bitcoin address.
- **Parameters**:
    - `address`: Bitcoin address (path parameter).

### 3. Set up UTXO for Bitcoin Transaction

- **Endpoint**: `/send`
- **Method**: POST
- **Description**: Creates a UTXO set for a Bitcoin address and a specific value.
- **Request Body**:
    - `address`: Bitcoin address.
    - `amount`: Quantity of Bitcoin.

### 4. Get Details by Bitcoin Transaction ID

- **Endpoint**: `/tx/{tx}`
- **Method**: GET
- **Description**: Returns detailed details for a specific Bitcoin transaction ID.
- **Parameters**:
    - `tx`: Bitcoin Transaction ID (path parameter).

## How to run

- **Dependecies**: Install the dependencies using go mod download.
- **Run**: Run the service using go run main.go.


## Swagger
After run, access:http://localhost:8080/swagger/index.html#/ to see swagger documentation
