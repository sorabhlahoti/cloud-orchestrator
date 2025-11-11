### Cloud Resource Orchestrator 🚀

```markdown

[![CI/CD Pipeline](https://github.com/sorabhlahoti/cloud-orchestrator/actions/workflows/ci-cd.yaml/badge.svg)](https://github.com/sorabhlahoti/cloud-orchestrator/actions/workflows/ci-cd.yaml)
[![Go Version](https://img.shields.io/badge/Go-1.24.3-00ADD8?logo=go)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker)](https://www.docker.com/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Ready-326CE5?logo=kubernetes)](https://kubernetes.io/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

> A production-ready, cloud-native resource orchestration API demonstrating modern microservices architecture, containerization, and observability patterns.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [Monitoring](#monitoring)
- [Development](#development)
- [Testing](#testing)
- [CI/CD Pipeline](#cicd-pipeline)
- [Project Structure](#project-structure)
- [Technology Stack](#technology-stack)
- [Future Enhancements](#future-enhancements)
- [License](#license)

---

## 🎯 Overview

The **Cloud Resource Orchestrator** is a RESTful API service built with Go that demonstrates cloud resource lifecycle management patterns. It provides a control plane for provisioning, tracking, and managing resources through a clean HTTP interface.

### What This Project Demonstrates

✅ **Backend Development** - Scalable REST API design with Go  
✅ **Cloud-Native Architecture** - Containerization with Docker, Kubernetes-ready deployment  
✅ **Observability** - Prometheus metrics integration and health monitoring  
✅ **DevOps Best Practices** - Complete CI/CD pipeline with automated testing and security scanning  
✅ **Production Patterns** - Error handling, logging, concurrent request management  

### Purpose

This is a **proof-of-concept** that simulates cloud orchestration patterns without actual infrastructure provisioning. It demonstrates the architectural foundations used in production cloud platforms like AWS, Azure, and GCP.

---

## ✨ Features

### Core Functionality
- 🔄 **Resource Lifecycle Management** - Create, list, and delete resources via REST API
- 🏥 **Health Monitoring** - Built-in health check endpoint for service monitoring
- 📊 **Metrics Export** - Prometheus-compatible metrics endpoint
- 🔒 **Thread-Safe Operations** - Concurrent request handling with mutex locks
- 📝 **Request Logging** - Comprehensive logging with timestamps and duration tracking
- ⚡ **Fast Response Times** - In-memory storage for instant operations

### DevOps & Infrastructure
- 🐳 **Fully Containerized** - Docker support with multi-stage builds
- ☸️ **Kubernetes Ready** - Production-grade deployment manifests
- 🔄 **CI/CD Automation** - GitHub Actions pipeline with 5 validation stages
- 🎯 **Resource Management** - Configured CPU/memory limits and requests
- 📈 **Auto-Scaling Ready** - Multi-replica deployment support
- 🛡️ **Security Scanning** - Automated vulnerability detection with Trivy

### Code Quality
- ✅ **Unit Tests** - Comprehensive test coverage
- 🎨 **Code Formatting** - Automated `gofmt` checks
- 🔍 **Static Analysis** - Go vet integration
- 📊 **Coverage Reports** - Automated generation in CI
- 📐 **Best Practices** - Clean architecture and error handling

---

## 🏗️ Architecture

### High-Level Design

```
┌─────────────────────────────────────────────────────────┐
│                        Client                           │
│                (Browser/curl/Postman)                   │
└────────────────────┬────────────────────────────────────┘
                     │ HTTP/REST
                     ▼
┌─────────────────────────────────────────────────────────┐
│          Cloud Orchestrator API (Go)                    │
│                                                         │
│  ┌────────────────────────────────────────────────┐   │
│  │           REST API Handlers                    │   │
│  │  ┌──────────────┐  ┌──────────────┐          │   │
│  │  │ Provision    │  │ List         │          │   │
│  │  │ /provision   │  │ /resources   │          │   │
│  │  └──────────────┘  └──────────────┘          │   │
│  │  ┌──────────────┐  ┌──────────────┐          │   │
│  │  │ Delete       │  │ Health Check │          │   │
│  │  │ /resources/* │  │ /health      │          │   │
│  │  └──────────────┘  └──────────────┘          │   │
│  └────────────────────────────────────────────────┘   │
│                                                         │
│  ┌────────────────────────────────────────────────┐   │
│  │         Middleware Layer                       │   │
│  │  -  Request Logging                             │   │
│  │  -  Metrics Collection                          │   │
│  │  -  Error Handling                              │   │
│  └────────────────────────────────────────────────┘   │
│                                                         │
│  ┌────────────────────────────────────────────────┐   │
│  │      In-Memory Storage (Thread-Safe)           │   │
│  │      map[int]Resource + sync.RWMutex           │   │
│  └────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
                     │
                     ▼
              ┌─────────────┐
              │ Prometheus  │
              │  /metrics   │
              └─────────────┘
```

### Kubernetes Deployment Architecture

```
                    Internet
                       │
                ┌──────▼──────┐
                │ LoadBalancer│
                │  (Service)  │
                └──────┬──────┘
                       │
        ┌──────────────┼──────────────┐
        │              │              │
   ┌────▼────┐    ┌───▼────┐    ┌───▼────┐
   │  Pod 1  │    │ Pod 2  │    │ Pod N  │
   │ :8080   │    │ :8080  │    │ :8080  │
   └────┬────┘    └───┬────┘    └───┬────┘
        │             │             │
        └─────────────┼─────────────┘
                      │
              ┌───────▼────────┐
              │   Prometheus   │
              │    Scraper     │
              └────────────────┘
```

---

## 🚀 Quick Start

### Prerequisites

- **Go**: 1.24.3 or higher - [Download](https://go.dev/dl/)
- **Docker**: Latest version - [Download](https://www.docker.com/get-started)
- **kubectl** (optional): For Kubernetes deployment - [Install](https://kubernetes.io/docs/tasks/tools/)

### Option 1: Run with Go (Development)

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

### Option 2: Run with Docker (Recommended)

```
# Build Docker image
docker build -t cloud-orchestrator:latest .

# Run container
docker run -p 8080:8080 cloud-orchestrator:latest

# Access at http://localhost:8080
```

### Verify Installation

```
# Check health
curl http://localhost:8080/health

# Expected: {"status":"healthy"}

# Create a resource
curl -X POST http://localhost:8080/provision

# Expected: {"id":123456,"name":"resource-123456","created_at":"..."}
```

---

## 📡 API Documentation

### Base URL
```
http://localhost:8080
```

### Endpoints Overview

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `POST` | `/provision` | Create new resource |
| `GET` | `/resources` | List all resources |
| `DELETE` | `/resources/{id}` | Delete specific resource |
| `GET` | `/metrics` | Prometheus metrics |

---

### 🏥 Health Check

```
GET /health
```

**Response:**
```
{
  "status": "healthy"
}
```

**Status Codes:**
- `200 OK` - Service is healthy

**Example:**
```
curl http://localhost:8080/health
```

---

### ➕ Create Resource

```
POST /provision
```

**Response:**
```
{
  "id": 123456,
  "name": "resource-123456",
  "created_at": "2025-11-11T10:00:00Z"
}
```

**Status Codes:**
- `201 Created` - Resource successfully created
- `405 Method Not Allowed` - Invalid HTTP method

**Example:**
```
curl -X POST http://localhost:8080/provision
```

---

### 📋 List Resources

```
GET /resources
```

**Response:**
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

**Status Codes:**
- `200 OK` - Resources retrieved successfully

**Example:**
```
curl http://localhost:8080/resources
```

---

### 🗑️ Delete Resource

```
DELETE /resources/{id}
```

**Parameters:**
- `id` (path) - Resource ID to delete

**Success Response:**
```
{
  "message": "Deleted resource 123456"
}
```

**Error Response:**
```
{
  "error": "Resource not found",
  "message": "Resource with ID 123456 does not exist"
}
```

**Status Codes:**
- `200 OK` - Resource deleted
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Resource doesn't exist
- `405 Method Not Allowed` - Invalid HTTP method

**Examples:**
```
# Delete resource
curl -X DELETE http://localhost:8080/resources/123456

# Invalid ID
curl -X DELETE http://localhost:8080/resources/abc
# Returns: 400 Bad Request
```

---

### 📊 Prometheus Metrics

```
GET /metrics
```

**Metrics Exposed:**

| Metric | Type | Description |
|--------|------|-------------|
| `http_requests_total` | Counter | Total HTTP requests by method/endpoint/status |
| `resources_provisioned_total` | Counter | Total resources created |
| `active_resources` | Gauge | Current active resources |

**Example Output:**
```
# HELP http_requests_total Total number of HTTP requests
# TYPE http_requests_total counter
http_requests_total{endpoint="/provision",method="POST",status="200"} 5

# HELP resources_provisioned_total Total number of resources provisioned
# TYPE resources_provisioned_total counter
resources_provisioned_total 5

# HELP active_resources Current number of active resources
# TYPE active_resources gauge
active_resources 3
```

**Example:**
```
curl http://localhost:8080/metrics
```

---

## ☸️ Deployment

### Docker Deployment

#### Build & Run

```
# Build image
docker build -t cloud-orchestrator:latest .

# Run container
docker run -d \
  --name orchestrator \
  -p 8080:8080 \
  cloud-orchestrator:latest

# View logs
docker logs -f orchestrator

# Stop container
docker stop orchestrator
```

#### Docker Compose (Optional)

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

Run:
```
docker-compose up -d
```

---

### Kubernetes Deployment

#### Prerequisites
- Kubernetes cluster (minikube, Docker Desktop with K8s, or cloud provider)
- `kubectl` configured

#### Deploy

```
# Apply configurations
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

# Check deployment
kubectl get pods
kubectl get services

# Expected output:
# NAME                                  READY   STATUS    RESTARTS   AGE
# cloud-orchestrator-xxxxx              1/1     Running   0          30s
# cloud-orchestrator-yyyyy              1/1     Running   0          30s
```

#### Access Application

**Using Port Forward (Development):**
```
kubectl port-forward service/cloud-orchestrator 8080:80
curl http://localhost:8080/health
```

**Using LoadBalancer (Cloud):**
```
kubectl get services cloud-orchestrator
# Note EXTERNAL-IP and access at http://EXTERNAL-IP
```

**Using Minikube:**
```
minikube service cloud-orchestrator
```

#### Scale Deployment

```
# Scale to 5 replicas
kubectl scale deployment cloud-orchestrator --replicas=5

# Verify
kubectl get pods
```

#### View Logs

```
# All pods
kubectl logs -l app=cloud-orchestrator

# Specific pod
kubectl logs cloud-orchestrator-xxxxx -f
```

#### Delete Deployment

```
kubectl delete -f k8s/deployment.yaml
kubectl delete -f k8s/service.yaml
```

---

### Kubernetes Configuration Details

**Deployment Features:**
- **Replicas**: 2 instances for high availability
- **Health Probes**:
  - Liveness: `/health` every 10s
  - Readiness: `/health` every 5s
- **Resource Limits**:
  - Memory: 64Mi request, 128Mi limit
  - CPU: 250m request, 500m limit
- **Prometheus**: Auto-discovery annotations enabled

**Service Configuration:**
- **Type**: LoadBalancer
- **Port**: 80 (external) → 8080 (internal)
- **Selector**: `app: cloud-orchestrator`

---

## 📊 Monitoring

### Prometheus Setup

The application exposes metrics at `/metrics` compatible with Prometheus.

#### Prometheus Configuration

```
# prometheus/prometheus.yaml
scrape_configs:
  - job_name: 'cloud-orchestrator'
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
        action: keep
        regex: true
```

#### Key Metrics

```