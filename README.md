
# Message Service

A simple Go-based microservice for managing messages, containerized with Docker and integrated with DynamoDB via LocalStack.

---

## **Getting Started**

Follow the instructions below to run the project locally using Docker and Docker Compose.

### **Prerequisites**

Ensure the following are installed on your system:
- Docker
- Docker Compose
- AWS CLI (optional, for configuration)

---

## **Setup and Configuration**

### **Environment Variables**

The following environment variables are required for AWS and LocalStack configuration:

- `AWS_ACCESS_KEY_ID` – Your AWS Access Key ID  
- `AWS_SECRET_ACCESS_KEY` – Your AWS Secret Access Key  
- `AWS_REGION` – AWS Region (e.g., `us-west-2`)  
- `LOCALSTACK_HOSTNAME` – Hostname for LocalStack (default: `localstack`)

---

## **Run Locally**

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Build and start the services:

   ```bash
   docker-compose up --build
   ```

   This will start:
   - The Go message service
   - LocalStack with DynamoDB

3. Verify that the application is running:

   ```bash
   curl -X POST http://localhost:8080/messages -H "Content-Type: application/json" -d '{"content": "Hello, world!"}'
   ```

   You should see a success response if the service is up and running.

---

## **Project Structure**

```plaintext
/message-service
├── Dockerfile                # Docker configuration for the Go app
├── docker-compose.yml        # Docker Compose configuration
├── go.mod                    # Go modules definition
├── go.sum                    # Go dependencies checksum
├── main.go                   # Main application logic
└── README.md                 # Project documentation (this file)
```

---

## **LocalStack Configuration**

- LocalStack runs DynamoDB locally to simulate AWS services.
- You can access it via port `4566`.

To create a table in DynamoDB:

```bash
aws dynamodb create-table     --table-name Messages     --attribute-definitions AttributeName=MessageID,AttributeType=S     --key-schema AttributeName=MessageID,KeyType=HASH     --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1     --endpoint-url http://localhost:4566
```

---

## **License**

This project is licensed under the [MIT License](LICENSE).

---
