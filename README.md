# Microservices Architecture – Go, gRPC, Docker

A scalable, modular microservices system built in Go, leveraging gRPC for efficient inter-service communication and Docker for easy deployment. This project demonstrates clean architecture, robust API design, and production-ready engineering practices.

---

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Development](#development)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

This repository contains a set of microservices designed to operate as a distributed, cloud-native system. Each service is independently deployable, communicates via gRPC, and is containerized for seamless orchestration. The architecture is ideal for scalable backend applications, demonstrating best practices in Go, API design, and DevOps.

---

## Architecture

Client --> (HTTP/JSON) --> API Gateway
API Gateway --> (gRPC/REST) --> [Service A, Service B, Service C, ...]
Services <--> Database(s)


- **API Gateway**: Routes client requests to appropriate backend services.
- **gRPC**: Enables high-performance, type-safe communication between services.
- **Docker**: Provides containerization for local and cloud deployment.

---

## Features

- Modular, independently deployable microservices
- High-performance gRPC APIs
- Clean architecture and separation of concerns
- Dockerized for local and cloud deployment
- Environment-based configuration management
- Ready for extension with persistent storage, caching, and observability

---

## Tech Stack

- **Go** (Golang)
- **gRPC** & **Protocol Buffers**
- **Docker** & **Docker Compose**
- **PostgreSQL / MongoDB** (as needed)
- **Redis** (optional, for caching)
- **Makefile** for build automation

---

## Project Structure

├── api/ # Protobuf files for gRPC and API definitions
├── cmd/ # Entrypoints for each microservice
│ ├── serviceA/
│ ├── serviceB/
│ └── gateway/
├── internal/ # Shared/internal packages (business logic, models)
├── pkg/ # Service-specific logic
├── configs/ # Configuration files (env, YAML, etc.)
├── deploy/ # Docker and deployment scripts
├── tests/ # Integration and E2E tests
├── go.mod
├── go.sum
└── README.md


---

## Getting Started

### Prerequisites

- [Go 1.20+](https://golang.org/dl/)
- [protoc](https://grpc.io/docs/protoc-installation/) (Protocol Buffers compiler)
- [Docker](https://www.docker.com/)

### Clone the Repository

git clone https://github.com/Pavan-Pandya1/Microservices.git
cd Microservices


### Build and Run Locally

1. **Generate gRPC code:**

make proto

2. **Now, Run each service (in seprate terminals):**

go run cmd/serviceA/main.go
go run cmd/serviceB/main.go
go run cmd/gateway/main.go

3. **Or use Docker Compose:** 

docker-compose up --build


4. **Access the API Gateway:**
- Default: `http://localhost:8080`

---

## API Endpoints

### Example HTTP Endpoints (update as per your services)
- `POST /api/resource` — Create a new resource
- `GET /api/resource` — List resources
- `PUT /api/resource/:id` — Update a resource
- `DELETE /api/resource/:id` — Delete a resource

### Example gRPC Services
- `CreateResource`
- `GetResource`
- `UpdateResource`
- `DeleteResource`

---

## Development

- **Code Generation:**  
Run `make proto` to regenerate Go code from `.proto` definitions.
- **Testing:**  
Run `go test ./...` to execute unit tests.
- **Linting:**  
Run `make lint` to check code quality.
- **Extending:**  
Add new services or swap in different databases as needed.

---

## Testing

To run all tests:

make test


To generate a test coverage report:

make test_coverage


---

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss your ideas.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a pull request

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

**Built with ❤️ by [Pavan Pandya](https://github.com/Pavan-Pandya1)**

