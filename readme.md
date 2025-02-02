# Fetch Demo

This is a demo implementation of the application described in https://github.com/fetch-rewards/receipt-processor-challenge

## Prerequisites

- Go 1.22 or higher

## Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/SEary342/fetch-demo.git
   cd fetch-demo
   ```

2. **Install Dependencies**

    This project uses Go modules to manage dependencies. Run the following command to fetch and install them:

    ```bash
    go mod tidy
    ```

## Running the Application

1. **Start the Server**

    Run the following command to start the application:
   ```bash
   go run main.go
   ```

    The application will start the server on `0.0.0.0:8080`.

2. **Send Requests to the Application**

   The application will respond to:
   * `POST` requests at `/receipts/process` with reciepts of the format described in [reciept-processor-demo readme](https://github.com/fetch-rewards/receipt-processor-challenge/blob/main/README.md)
   * `GET` requests to `/receipts/{id}/points` with a UUID provided via the POST request