```markdown
# Cloud Resource Orchestrator 🚀

[![CI/CD Pipeline](https://github.com/sorabhlahoti/cloud-orchestrator/actions/workflows/ci-cd.yaml/badge.svg)](https://github.com/sorabhlahoti/cloud-orchestrator/actions/workflows/ci-cd.yaml)
[![Go Version](https://img.shields.io/badge/Go-1.24.3-00ADD8?logo=go)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker)](https://www.docker.com/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Ready-326CE5?logo=kubernetes)](https://kubernetes.io/)

A production-ready, cloud-native resource orchestration API built with Go, demonstrating modern microservices architecture, containerization, and observability patterns.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Quick Start](#quick-start)
- [API Endpoints](#api-endpoints)
- [Running with Docker](#running-with-docker)
- [Kubernetes Deployment](#kubernetes-deployment)
- [Monitoring](#monitoring)
- [Testing](#testing)
- [CI/CD Pipeline](#cicd-pipeline)
- [Project Structure](#project-structure)
- [Future Enhancements](#future-enhancements)

---

## 🎯 Overview

The **Cloud Resource Orchestrator** is a RESTful API service that demonstrates cloud resource lifecycle management. It provides endpoints for provisioning, tracking, and managing resources through a clean HTTP interface.

**What This Demonstrates:**
- Building scalable REST APIs with Go
- Cloud-native architecture with Docker and Kubernetes
- Observability with Prometheus metrics
- Complete CI/CD pipeline with automated testing
- Production-ready error handling and logging

**Note:** This is a proof-of-concept that manages logical resource entries rather than actual infrastructure, demonstrating the architectural patterns used in production cloud platforms.

---

## ✨ Features

**Core Functionality:**
- REST API for resource CRUD operations
- Health monitoring endpoint
- Prometheus metrics export
- Thread-safe concurrent operations
- Request logging with timestamps

**DevOps & Infrastructure:**
- Fully containerized with Docker
- Kubernetes-ready deployment manifests
- CI/CD automation with GitHub Actions
- Multi-replica deployment support
- Automated security scanning

**Code Quality:**
- Unit tests with coverage reports
- Automated code formatting checks
- Static analysis with go vet
- Vulnerability scanning with Trivy

---

## 🛠️ Technology Stack

- **Language:** Go 1.24.3
- **Framework:** Standard library (net/http)
- **Monitoring:** Prometheus
- **Containerization:** Docker
- **Orchestration:** Kubernetes
- **CI/CD:** GitHub Actions
- **Testing:** Go testing framework
- **Security:** Trivy scanner

---

## 🚀 Quick Start

### Prerequisites

- Go 1.24.3+ ([Download](https://go.dev/dl/))
- Docker ([Download](https://www.docker.com/get-started))
- kubectl (optional, for Kubernetes)

### Local Development

```
# Clone repository
git clone https://github.com/sorabhlahoti/cloud-orchestrator.git
cd cloud-orchestrator

# Install dependencies
go mod download

# Run application
go run main.go

# Application starts on http://localhost:8080
```

### Test the API

```
# Check health
curl http://localhost:8080/health

# Create a resource
curl -X POST http://localhost:8080/provision

# List all resources
curl http://localhost:8080/resources

# Delete a resource (replace ID)
curl -X DELETE http://localhost:8080/resources/123456

# View metrics
curl http://localhost:8080/metrics
```

---

## 📡 API Endpoints

### Health Check
```
GET /health
```
Returns service health status.

**Response:**
```
{"status":"healthy"}
```

---

### Create Resource
```
POST /provision
```
Creates a new resource.

**Response:**
```
{
  "id": 123456,
  "name": "resource-123456",
  "created_at": "2025-11-11T10:00:00Z"
}
```

**Status Codes:**
- `201 Created` - Success
- `405 Method Not Allowed` - Invalid method

---

### List Resources
```
GET /resources
```
Returns all resources.

**Response:**
```
[
  {
    "id": 123456,
    "name": "resource-123456",
    "created_at": "2025-11-11T10:00:00Z"
  }
]
```

---

### Delete Resource
```
DELETE /resources/{id}
```
Deletes a specific resource by ID.

**Success Response:**
```
{"message": "Deleted resource 123456"}
```

**Error Response:**
```
{
  "error": "Resource not found",
  "message": "Resource with ID 123456 does not exist"
}
```

**Status Codes:**
- `200 OK` - Deleted successfully
- `400 Bad Request` - Invalid ID
- `404 Not Found` - Resource doesn't exist

---

### Prometheus Metrics
```
GET /metrics
```
Returns Prometheus-compatible metrics.

**Key Metrics:**
- `http_requests_total` - Total HTTP requests
- `resources_provisioned_total` - Total resources created
- `active_resources` - Current active resources

---

## 🐳 Running with Docker

### Build and Run

```
# Build Docker image
docker build -t cloud-orchestrator:latest .

# Run container
docker run -d -p 8080:8080 --name orchestrator cloud-orchestrator:latest

# View logs
docker logs -f orchestrator

# Stop and remove
docker stop orchestrator
docker rm orchestrator
```

### Docker Compose

Create `docker-compose.yml`:

```
version: '3.8'
services:
  orchestrator:
    build: .
    ports:
      - "8080:8080"
    restart: unless-stopped
```

Run with:
```
docker-compose up -d
```

---

## ☸️ Kubernetes Deployment

### Deploy to Kubernetes

```
# Apply deployment and service
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

# Check status
kubectl get pods
kubectl get services

# View logs
kubectl logs -l app=cloud-orchestrator

# Scale replicas
kubectl scale deployment cloud-orchestrator --replicas=5
```

### Access the Application

**Port Forward (Development):**
```
kubectl port-forward service/cloud-orchestrator 8080:80
curl http://localhost:8080/health
```

**LoadBalancer (Cloud):**
```
kubectl get services
# Access via EXTERNAL-IP
```

**Minikube:**
```
minikube service cloud-orchestrator
```

### Delete Deployment

```
kubectl delete -f k8s/deployment.yaml
kubectl delete -f k8s/service.yaml
```

---

## 📊 Monitoring

### Prometheus Metrics

The application exposes metrics at `/metrics` endpoint in Prometheus format.

**Available Metrics:**
- `http_requests_total{method, endpoint, status}` - Request counter
- `resources_provisioned_total` - Total resources created
- `active_resources` - Current resource count

**Example:**
```
curl http://localhost:8080/metrics
```

### Kubernetes Monitoring

Prometheus can automatically discover and scrape metrics using the annotations in the deployment manifest:

```
annotations:
  prometheus.io/scrape: "true"
  prometheus.io/port: "8080"
  prometheus.io/path: "/metrics"
```

---

## 🧪 Testing

### Run Tests

```
# Run all tests
go test -v

# Run with coverage
go test -v -coverprofile=coverage.out

# View coverage report
go tool cover -html=coverage.out
```

### Test Coverage

The project includes unit tests for:
- Health check handler
- Resource provisioning
- Resource listing
- Resource deletion
- Error handling scenarios

---

## 🔄 CI/CD Pipeline

The project uses GitHub Actions for automated testing and deployment.

**Pipeline Stages:**

1. **Test** - Runs unit tests with coverage reporting
2. **Build** - Compiles Go binary
3. **Docker** - Builds and tests Docker image
4. **Lint** - Checks code formatting (gofmt, go vet)
5. **Security** - Scans for vulnerabilities with Trivy

**Trigger:** Automatic on every push to `main` branch

**View Status:** Check the Actions tab in GitHub repository

---

## 📁 Project Structure

```
cloud-orchestrator/
├── main.go                    # Main application code
├── main_test.go              # Unit tests
├── go.mod                    # Go module definition
├── go.sum                    # Go dependencies
├── Dockerfile                # Docker build instructions
├── README.md                 # Project documentation
├── k8s/
│   ├── deployment.yaml       # Kubernetes deployment config
│   └── service.yaml          # Kubernetes service config
├── prometheus/
│   └── prometheus.yaml       # Prometheus configuration
└── .github/
    └── workflows/
        └── ci-cd.yaml        # CI/CD pipeline
```

---

## 🚀 Future Enhancements

**Potential improvements for production use:**

- Add persistent storage (PostgreSQL/Redis) instead of in-memory map
- Implement JWT authentication for API endpoints
- Add API rate limiting to prevent abuse
- Implement API versioning (e.g., `/api/v1/resources`)
- Add distributed tracing with Jaeger or OpenTelemetry
- Implement graceful shutdown handling
- Add more comprehensive error handling and validation
- Create OpenAPI/Swagger documentation
- Add integration tests for end-to-end scenarios
- Implement actual cloud provider integration (AWS/Azure/GCP)

---

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

## 👤 Author

**Sorabh Lahoti**

- GitHub: [@sorabhlahoti](https://github.com/sorabhlahoti)
- Project Link: [https://github.com/sorabhlahoti/cloud-orchestrator](https://github.com/sorabhlahoti/cloud-orchestrator)

---

## 🙏 Acknowledgments

- Built with [Go](https://go.dev/)
- Containerized with [Docker](https://www.docker.com/)
- Orchestrated with [Kubernetes](https://kubernetes.io/)
- Monitored with [Prometheus](https://prometheus.io/)

---

## 📚 Learning Resources

If you want to understand the technologies used in this project:

- **Go Basics:** [Official Go Tutorial](https://go.dev/tour/)
- **REST APIs:** [REST API Tutorial](https://restfulapi.net/)
- **Docker:** [Docker Getting Started](https://docs.docker.com/get-started/)
- **Kubernetes:** [Kubernetes Basics](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
- **Prometheus:** [Prometheus Documentation](https://prometheus.io/docs/introduction/overview/)

---

**⭐ If you find this project helpful, please consider giving it a star!**
```

***

## ✅ **What's Included:**

- Professional badges at the top
- Clear table of contents
- Simple explanations without complex diagrams
- All necessary commands
- Clean formatting
- Easy to read sections
- Future enhancements section
- Learning resources

## 🎯 **This README Shows:**

- You understand the full stack
- You can document professionally
- You follow best practices
- You're thinking about production
- You're open to improvements

