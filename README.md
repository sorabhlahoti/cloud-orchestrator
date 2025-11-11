```markdown
# Cloud Resource Orchestrator 🚀

[![CI/CD Pipeline](https://github.com/sorabhlahoti/cloud-orchestrator/actions/workflows/ci-cd.yaml/badge.svg)](https://github.com/sorabhlahoti/cloud-orchestrator/actions/workflows/ci-cd.yaml)
[![Go Version](https://img.shields.io/badge/Go-1.24.3-00ADD8?logo=go)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker)](https://www.docker.com/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Ready-326CE5?logo=kubernetes)](https://kubernetes.io/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A production-ready, cloud-native resource orchestration API built with Go, demonstrating modern microservices architecture, containerization, and observability patterns.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [Monitoring](#monitoring)
- [Development](#development)
- [Testing](#testing)
- [CI/CD Pipeline](#cicd-pipeline)
- [Project Structure](#project-structure)
- [Future Enhancements](#future-enhancements)
- [Contributing](#contributing)
- [License](#license)

---

## 🎯 Overview

The **Cloud Resource Orchestrator** is a RESTful API service that simulates cloud resource lifecycle management. It provides a control plane for provisioning, tracking, and managing resources through a clean HTTP interface, demonstrating the architectural patterns used in production cloud platforms.

### What This Project Demonstrates

- ✅ **Backend Development**: Building scalable REST APIs with Go
- ✅ **Cloud-Native Design**: Containerization and Kubernetes-ready deployment
- ✅ **Observability**: Integrated Prometheus metrics and health monitoring
- ✅ **DevOps Practices**: Complete CI/CD pipeline with automated testing
- ✅ **Production Patterns**: Error handling, logging, concurrency management

### Use Case

While this is a proof-of-concept that manages logical resource entries rather than actual infrastructure, it demonstrates the **complete architecture** of a cloud orchestration system - from API design to deployment strategies.

---

## ✨ Features

### Core Functionality
- 🔄 **Resource Lifecycle Management**: Create, list, and delete resources via REST API
- 🏥 **Health Monitoring**: Built-in health check endpoint for service monitoring
- 📊 **Metrics Export**: Prometheus-compatible metrics for observability
- 🔒 **Thread-Safe Operations**: Concurrent request handling with mutex locks
- 📝 **Request Logging**: Comprehensive logging with timestamps and duration tracking

### DevOps & Infrastructure
- 🐳 **Docker Support**: Fully containerized application
- ☸️ **Kubernetes Ready**: Production-ready deployment manifests with health probes
- 🔄 **CI/CD Automation**: GitHub Actions pipeline with testing, linting, and security scanning
- 🎯 **Resource Management**: Kubernetes resource limits and requests configured
- 📈 **Horizontal Scaling**: Multi-replica deployment support with load balancing

### Code Quality
- ✅ **Unit Tests**: Comprehensive test coverage with Go's testing framework
- 🎨 **Code Formatting**: Automated `gofmt` checks in CI pipeline
- 🔍 **Static Analysis**: Go vet integration for code quality
- 🛡️ **Security Scanning**: Trivy vulnerability scanning
- 📊 **Coverage Reports**: Automated test coverage generation

---

## 🏗️ Architecture

### System Design

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │ HTTP/REST
       ▼
┌─────────────────────────────────┐
│   Cloud Orchestrator API        │
│   (Go HTTP Server)              │
│                                 │
│  ┌──────────────────────────┐  │
│  │  REST Handlers           │  │
│  │  ├─ Provision            │  │
│  │  ├─ List                 │  │
│  │  ├─ Delete               │  │
│  │  └─ Health Check         │  │
│  └──────────────────────────┘  │
│                                 │
│  ┌──────────────────────────┐  │
│  │  Middleware Layer        │  │
│  │  ├─ Logging              │  │
│  │  └─ Metrics Collection   │  │
│  └──────────────────────────┘  │
│                                 │
│  ┌──────────────────────────┐  │
│  │  In-Memory Storage       │  │
│  │  (Thread-Safe Map)       │  │
│  └──────────────────────────┘  │
└─────────────────────────────────┘
       │
       ▼
┌─────────────┐
│ Prometheus  │
│  /metrics   │
└─────────────┘
```

### Deployment Architecture (Kubernetes)

```
                    ┌─────────────┐
                    │   Internet  │
                    └──────┬──────┘
                           │
                    ┌──────▼──────────┐
                    │  Load Balancer  │
                    │   (Service)     │
                    └──────┬──────────┘
                           │
            ┌──────────────┼──────────────┐
            │              │              │
      ┌─────▼─────┐  ┌────▼──────┐  ┌───▼──────┐
      │   Pod 1   │  │   Pod 2   │  │  Pod N   │
      │           │  │           │  │          │
      │ Container │  │ Container │  │Container │
      │  :8080    │  │  :8080    │  │ :8080    │
      └───────────┘  └───────────┘  └──────────┘
           │              │              │
           └──────────────┼──────────────┘
                          │
                  ┌───────▼────────┐
                  │   Prometheus   │
                  │   Scraper      │
                  └────────────────┘
```

---

## 🛠️ Technology Stack

| Category | Technology | Purpose |
|----------|-----------|---------|
| **Language** | Go 1.24.3 | Backend API implementation |
| **Framework** | net/http (stdlib) | HTTP server and routing |
| **Monitoring** | Prometheus | Metrics collection and monitoring |
| **Containerization** | Docker | Application packaging |
| **Orchestration** | Kubernetes | Container orchestration and scaling |
| **CI/CD** | GitHub Actions | Automated testing and deployment |
| **Testing** | Go testing | Unit and integration tests |
| **Security** | Trivy | Vulnerability scanning |

---

## 🚀 Getting Started

### Prerequisites

- **Go**: Version 1.24.3 or higher ([Download](https://go.dev/dl/))
- **Docker**: Latest version ([Download](https://www.docker.com/get-started))
- **kubectl** (Optional): For Kubernetes deployment ([Install](https://kubernetes.io/docs/tasks/tools/))

### Quick Start

#### 1️⃣ Clone the Repository

```
git clone https://github.com/sorabhlahoti/cloud-orchestrator.git
cd cloud-orchestrator
```

#### 2️⃣ Run Locally (Development)

```
# Download dependencies
go mod download

# Run the application
go run main.go
```

The API will be available at `http://localhost:8080`

#### 3️⃣ Run with Docker

```
# Build Docker image
docker build -t cloud-orchestrator:latest .

# Run container
docker run -p 8080:8080 cloud-orchestrator:latest
```

#### 4️⃣ Verify Installation

```
# Check health
curl http://localhost:8080/health

# Expected output:
# {"status":"healthy"}
```

---

## 📡 API Documentation

### Base URL
```
http://localhost:8080
```

### Endpoints

#### 🏥 Health Check
```
GET /health
```

**Description**: Check service health status

**Response**:
```
{
  "status": "healthy"
}
```

**Status Codes**:
- `200 OK`: Service is healthy

---

#### ➕ Provision Resource
```
POST /provision
```

**Description**: Create a new resource

**Response**:
```
{
  "id": 123456,
  "name": "resource-123456",
  "created_at": "2025-11-11T10:00:00Z"
}
```

**Status Codes**:
- `201 Created`: Resource successfully created
- `405 Method Not Allowed`: Invalid HTTP method

**Example**:
```
curl -X POST http://localhost:8080/provision
```

---

#### 📋 List Resources
```
GET /resources
```

**Description**: Retrieve all resources

**Response**:
```
[
  {
    "id": 123456,
    "name": "resource-123456",
    "created_at": "2025-11-11T10:00:00Z"
  },
  {
    "id": 789012,
    "name": "resource-789012",
    "created_at": "2025-11-11T10:05:00Z"
  }
]
```

**Status Codes**:
- `200 OK`: Successfully retrieved resources
- `405 Method Not Allowed`: Invalid HTTP method

**Example**:
```
curl http://localhost:8080/resources
```

---

#### 🗑️ Delete Resource
```
DELETE /resources/{id}
```

**Description**: Delete a specific resource by ID

**Parameters**:
- `id` (path): Resource ID to delete

**Response**:
```
{
  "message": "Deleted resource 123456"
}
```

**Error Response**:
```
{
  "error": "Resource not found",
  "message": "Resource with ID 123456 does not exist"
}
```

**Status Codes**:
- `200 OK`: Resource successfully deleted
- `400 Bad Request`: Invalid resource ID format
- `404 Not Found`: Resource does not exist
- `405 Method Not Allowed`: Invalid HTTP method

**Examples**:
```
# Delete resource with ID 123456
curl -X DELETE http://localhost:8080/resources/123456

# Invalid ID (returns 400)
curl -X DELETE http://localhost:8080/resources/abc
```

---

#### 📊 Prometheus Metrics
```
GET /metrics
```

**Description**: Export Prometheus-compatible metrics

**Metrics Exposed**:
- `http_requests_total{method, endpoint, status}`: Total HTTP requests
- `resources_provisioned_total`: Total resources created
- `active_resources`: Current number of active resources

**Response Format**: Prometheus text-based exposition format

**Example**:
```
curl http://localhost:8080/metrics
```

**Sample Output**:
```
# HELP http_requests_total Total number of HTTP requests
# TYPE http_requests_total counter
http_requests_total{endpoint="/",method="GET",status="200"} 15
http_requests_total{endpoint="/provision",method="POST",status="200"} 5

# HELP resources_provisioned_total Total number of resources provisioned
# TYPE resources_provisioned_total counter
resources_provisioned_total 5

# HELP active_resources Current number of active resources
# TYPE active_resources gauge
active_resources 3
```

---

### Error Responses

All error responses follow this format:

```
{
  "error": "Error Type",
  "message": "Detailed error description"
}
```

**Common Error Status Codes**:
- `400 Bad Request`: Invalid input or malformed request
- `404 Not Found`: Resource does not exist
- `405 Method Not Allowed`: Incorrect HTTP method
- `500 Internal Server Error`: Server-side error

---

## ☸️ Deployment

### Docker Deployment

#### Build Image
```
docker build -t cloud-orchestrator:latest .
```

#### Run Container
```
docker run -d \
  --name orchestrator \
  -p 8080:8080 \
  cloud-orchestrator:latest
```

#### View Logs
```
docker logs -f orchestrator
```

#### Stop Container
```
docker stop orchestrator
docker rm orchestrator
```

---

### Kubernetes Deployment

#### Prerequisites
- Running Kubernetes cluster (minikube, Docker Desktop, or cloud provider)
- `kubectl` configured with cluster access

#### Deploy Application

```
# Apply deployment configuration
kubectl apply -f k8s/deployment.yaml

# Apply service configuration
kubectl apply -f k8s/service.yaml
```

#### Verify Deployment

```
# Check pod status
kubectl get pods

# Expected output:
# NAME                                  READY   STATUS    RESTARTS   AGE
# cloud-orchestrator-xxxxx              1/1     Running   0          30s
# cloud-orchestrator-yyyyy              1/1     Running   0          30s

# Check service
kubectl get services

# Get detailed pod information
kubectl describe pod cloud-orchestrator-xxxxx
```

#### Access Application

**LoadBalancer (Cloud providers):**
```
kubectl get services cloud-orchestrator
# Note the EXTERNAL-IP and access at http://EXTERNAL-IP
```

**NodePort (Local/Minikube):**
```
minikube service cloud-orchestrator
# Opens browser automatically
```

**Port Forward (Development):**
```
kubectl port-forward service/cloud-orchestrator 8080:80
curl http://localhost:8080/health
```

#### Scale Deployment

```
# Scale to 5 replicas
kubectl scale deployment cloud-orchestrator --replicas=5

# Verify scaling
kubectl get pods -w
```

#### View Logs

```
# Logs from all pods
kubectl logs -l app=cloud-orchestrator

# Logs from specific pod
kubectl logs cloud-orchestrator-xxxxx

# Follow logs
kubectl logs -f cloud-orchestrator-xxxxx
```

#### Update Deployment

```
# Edit deployment
kubectl edit deployment cloud-orchestrator

# Or apply updated YAML
kubectl apply -f k8s/deployment.yaml
```

#### Delete Deployment

```
kubectl delete -f k8s/deployment.yaml
kubectl delete -f k8s/service.yaml
```

---

### Kubernetes Configuration Details

#### Deployment Features

- **Replicas**: 2 instances for high availability
- **Health Probes**:
  - **Liveness**: Checks `/health` every 10s
  - **Readiness**: Checks `/health` every 5s
- **Resource Limits**:
  - Memory: 128Mi max
  - CPU: 500m (0.5 cores) max
- **Annotations**: Prometheus scraping enabled

#### Service Configuration

- **Type**: LoadBalancer (change to ClusterIP for internal-only access)
- **Port Mapping**: External 80 → Internal 8080
- **Selector**: Routes traffic to pods with `app: cloud-orchestrator` label

---

## 📊 Monitoring

### Prometheus Integration

The application exposes metrics at `/metrics` endpoint in Prometheus format.

#### Metrics Available

| Metric Name | Type | Description |
|------------|------|-------------|
| `http_requests_total` | Counter | Total HTTP requests by method, endpoint, status |
| `resources_provisioned_