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

