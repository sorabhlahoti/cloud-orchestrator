
````markdown
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

```bash
# Clone repository
git clone https://github.com/sorabhlahoti/cloud-orchestrator.git
cd cloud-orchestrator

# Install dependencies
go mod download

# Run application
go run main.go

# Application starts on http://localhost:8080
````

### Test the API

```bash
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

```bash
GET /health
```

**Response:**

```json
{"status":"healthy"}
```

---

### Create Resource

```bash
POST /provision
```

**Response:**

```json
{
  "id": 123456,
  "name": "resource-123456",
  "created_at": "2025-11-11T10:00:00Z"
}
```

---

### List Resources

```bash
GET /resources
```

**Response:**

```json
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

```bash
DELETE /resources/{id}
```

**Response:**

```json
{"message": "Deleted resource 123456"}
```

---

### Prometheus Metrics

```bash
GET /metrics
```

**Key Metrics:**

* `http_requests_total`
* `resources_provisioned_total`
* `active_resources`

---

## 🐳 Running with Docker

### Build and Run

```bash
docker build -t cloud-orchestrator:latest .
docker run -d -p 8080:8080 --name orchestrator cloud-orchestrator:latest
```

### Docker Compose

```yaml
version: '3.8'
services:
  orchestrator:
    build: .
    ports:
      - "8080:8080"
    restart: unless-stopped
```

Run:

```bash
docker-compose up -d
```

---

## ☸️ Kubernetes Deployment

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl get pods
kubectl get services
```

**Port Forward (Development):**

```bash
kubectl port-forward service/cloud-orchestrator 8080:80
curl http://localhost:8080/health
```

---

## 📊 Monitoring

The application exposes metrics at `/metrics` endpoint.

Example:

```bash
curl http://localhost:8080/metrics
```

Prometheus can scrape metrics using deployment annotations:

```yaml
annotations:
  prometheus.io/scrape: "true"
  prometheus.io/port: "8080"
  prometheus.io/path: "/metrics"
```

---

## 🧪 Testing

```bash
go test -v
go test -v -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## 🔄 CI/CD Pipeline

**Stages:**

1. Test
2. Build
3. Docker
4. Lint
5. Security Scan

Triggered on every push to `main`.
View pipeline in GitHub Actions tab.

---

## 📁 Project Structure

```
cloud-orchestrator/
├── main.go
├── main_test.go
├── go.mod
├── go.sum
├── Dockerfile
├── README.md
├── k8s/
│   ├── deployment.yaml
│   └── service.yaml
├── prometheus/
│   └── prometheus.yaml
└── .github/
    └── workflows/
        └── ci-cd.yaml
```

---

## 🚀 Future Enhancements

* Add PostgreSQL or Redis storage
* JWT authentication
* API rate limiting
* API versioning
* Distributed tracing (Jaeger/OpenTelemetry)
* Graceful shutdown
* Swagger documentation
* Integration tests
* Cloud provider integration

---

## 📝 License

MIT License — see [LICENSE](LICENSE).

---

## 👤 Author

**Sorabh Lahoti**

* GitHub: [@sorabhlahoti](https://github.com/sorabhlahoti)
* Project: [https://github.com/sorabhlahoti/cloud-orchestrator](https://github.com/sorabhlahoti/cloud-orchestrator)

---

## 📚 Learning Resources

* [Go Tour](https://go.dev/tour/)
* [REST API Tutorial](https://restfulapi.net/)
* [Docker Docs](https://docs.docker.com/get-started/)
* [Kubernetes Basics](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
* [Prometheus Docs](https://prometheus.io/docs/introduction/overview/)

---

**⭐ If you find this project helpful, please consider giving it a star!**


