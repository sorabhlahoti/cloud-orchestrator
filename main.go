package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Resource struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

var (
	resources   = make(map[int]Resource)
	resourceMux = sync.RWMutex{}

	// Prometheus metrics
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
	resourcesProvisioned = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "resources_provisioned_total",
			Help: "Total number of resources provisioned",
		},
	)
	activeResources = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_resources",
			Help: "Current number of active resources",
		},
	)
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(resourcesProvisioned)
	prometheus.MustRegister(activeResources)
}

// Logging middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
		log.Printf("Completed in %v", time.Since(start))
	}
}

// Metrics middleware
func metricsMiddleware(endpoint string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		httpRequestsTotal.WithLabelValues(r.Method, endpoint, "200").Inc()
	}
}

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

// Provision (create) resource
func provisionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only POST is supported")
		return
	}

	resourceMux.Lock()
	defer resourceMux.Unlock()

	id := rand.Intn(1000000)
	name := fmt.Sprintf("resource-%d", id)
	res := Resource{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
	}
	resources[id] = res

	// Update Prometheus metrics
	resourcesProvisioned.Inc()
	activeResources.Set(float64(len(resources)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// List resources
func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only GET is supported")
		return
	}

	resourceMux.RLock()
	defer resourceMux.RUnlock()

	resList := []Resource{}
	for _, res := range resources {
		resList = append(resList, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resList)
}

// Delete resource
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only DELETE is supported")
		return
	}

	idStr := r.URL.Path[len("/resources/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resource ID", "ID must be a number")
		return
	}

	resourceMux.Lock()
	defer resourceMux.Unlock()

	if _, ok := resources[id]; ok {
		delete(resources, id)
		activeResources.Set(float64(len(resources)))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Deleted resource %d", id)})
		return
	}

	respondWithError(w, http.StatusNotFound, "Resource not found", fmt.Sprintf("Resource with ID %d does not exist", id))
}

// Helper function for error responses
func respondWithError(w http.ResponseWriter, statusCode int, error string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   error,
		Message: message,
	})
}

func main() {
	http.HandleFunc("/", loggingMiddleware(metricsMiddleware("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Cloud Resource Orchestrator API is running 🚀\n")
	})))

	http.HandleFunc("/health", loggingMiddleware(metricsMiddleware("/health", healthHandler)))
	http.HandleFunc("/provision", loggingMiddleware(metricsMiddleware("/provision", provisionHandler)))
	http.HandleFunc("/resources", loggingMiddleware(metricsMiddleware("/resources", listHandler)))
	http.HandleFunc("/resources/", loggingMiddleware(metricsMiddleware("/resources/", deleteHandler)))

	// Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	log.Println("🚀 Cloud Resource Orchestrator starting on :8080")
	log.Println("📊 Metrics available at /metrics")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
