# Smart Coffee

Smart Coffee is a Go API for a coffee ordering flow, built to run locally and in k3d Kubernetes with basic observability via Prometheus metrics.

## APIs

- `GET /coffee/?id=<id>`
  - Returns a coffee response by id
  - Returns `400` if `id` is missing
- `GET /metrics`
  - Prometheus metrics endpoint (`prometheus/client_golang`)

## Run Locally (Go)

```bash
cd app
go mod download
go run .
```

API will be available at `http://localhost:8080`.

## Run in Local Kubernetes (k3d)

### 1) Create cluster

```bash
brew install k3d
k3d cluster create coffee-cluster \
  -p "8080:80@loadbalancer" \
  -p "3000:3000@loadbalancer" \
  --agents 2
```

### 2) Deploy MySQL

```bash
kubectl apply -f k8s/mysql.yaml
```

### 3) Build and import Coffee API image

```bash
docker build -t smart-coffee:latest ./app
k3d image import smart-coffee:latest -c coffee-cluster
```

### 4) Deploy Coffee API

```bash
kubectl apply -f k8s/coffee-api.yaml
```

### 5) Port-forward API

```bash
kubectl port-forward svc/coffee-api 8080:80
```

### 6) Verify

```bash
curl "http://localhost:8080/coffee/?id=123"
curl "http://localhost:8080/metrics"
```

## Observability with Helm (Prometheus + Grafana)

```bash
# Add and update Helm repo
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# Install kube-prometheus-stack (Prometheus, Grafana, Alertmanager)
helm upgrade --install obs prometheus-community/kube-prometheus-stack \
  -n monitoring --create-namespace
```

Check pods:

```bash
kubectl get pods -n monitoring
```

Grafana port-forward:

```bash
kubectl port-forward svc/obs-grafana -n monitoring 3000:80
```

## Optional: Combined port-forward (API + Grafana)

```bash
kubectl port-forward svc/coffee-api 8080:80 & kubectl port-forward svc/obs-grafana -n monitoring 3000:80 & wait
```

## TODO

- [ ] Add k6 load tests
