# Cloud Resource Orchestrator 🚀

A lightweight, production-ready Kubernetes-based orchestration API built in Go with full observability.

## Features ✨

- **RESTful API** for cloud resource provisioning and management
- **In-memory resource storage** (easily extensible to databases)
- **Prometheus metrics** endpoint for monitoring
- **Request logging** with timestamps
- **JSON error responses** with proper HTTP status codes
- **Dockerized** and ready for container orchestration
- **Kubernetes manifests** included
- **CI/CD pipeline** with GitHub Actions
- **Unit tests** for core handlers

---

## API Endpoints 📡

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/` | Welcome message |
| `GET` | `/health` | Health check (returns JSON status) |
| `POST` | `/provision` | Create a new resource |
| `GET` | `/resources` | List all resources |
| `DELETE` | `/resources/{id}` | Delete resource by ID |
| `GET` | `/metrics` | Prometheus metrics |

---

## Quick Start 🏁

### Prerequisites
- Go 1.24.3+
- Docker
- (Optional) Kubernetes cluster

### Local Development

