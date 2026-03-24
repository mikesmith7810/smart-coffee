# Smart Coffee

Smart Coffee is a Go API for a coffee ordering flow, built to run locally and in k3d Kubernetes with basic observability via Prometheus metrics.

Database settings now come from a config file, with secret values overridable via environment variables.

## APIs

- `GET /coffee/?id=<id>`
  - Returns a coffee response by id
  - Returns `400` if `id` is missing
- `GET /metrics`
  - Prometheus metrics endpoint (`prometheus/client_golang`)

## Run in Local Kubernetes (k3d)

### 1) Create cluster

```bash
brew install k3d
k3d cluster delete coffee-cluster  # if one already exists
k3d cluster create coffee-cluster \
  -p "8080:80@loadbalancer" \
  --agents 2
```

### 2) Build and import Coffee API image

```bash
docker build -t smart-coffee:latest ./app
k3d image import smart-coffee:latest -c coffee-cluster
```

### 3) Set up local values (first time only)

```bash
cp helm/smart-coffee/values.local.example.yaml helm/smart-coffee/values.local.yaml
# edit values.local.yaml and fill in your passwords
```

### 4) Deploy app (MySQL + Coffee API)

```bash
helm upgrade --install smart-coffee ./helm/smart-coffee \
  -f helm/smart-coffee/values.local.yaml
```

### 5) Deploy observability (Prometheus + Grafana)

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm upgrade --install obs prometheus-community/kube-prometheus-stack \
  -n monitoring --create-namespace \
  -f helm/obs-values.yaml
```

### 6) Deploy Loki (log aggregation)

Loki collects logs from all pods. Promtail runs as a DaemonSet and ships logs to Loki. Grafana is pre-configured to use Loki as a data source via `helm/obs-values.yaml`.

```bash
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm upgrade --install loki grafana/loki-stack \
  -n monitoring \
  -f helm/loki-values.yaml
```

### 7) Verify

```bash
kubectl get pods
kubectl get pods -n monitoring
curl "http://localhost:8080/coffee/?id=123"
curl "http://localhost:8080/metrics"
```

### 8) Port-forward Grafana

```bash
kubectl port-forward svc/obs-grafana -n monitoring 3000:80
```

Grafana is available at `http://localhost:3000`.

### 9) Log in to Grafana

Get the generated admin password:

```bash
kubectl get secret obs-grafana -n monitoring -o jsonpath="{.data.admin-password}" | base64 --decode
```

Log in with:
- **Username:** `admin`
- **Password:** output from the command above

## Viewing logs in Grafana

1. Open Grafana at `http://localhost:3000`
2. Go to **Explore** (compass icon in the left sidebar)
3. Select **Loki** from the data source dropdown
4. Use the label filters to find your pods, e.g.:
   - `app = coffee-api` for the API logs
   - `app = mysql` for the database logs
5. Hit **Run query**

## Shared environments

Pass secrets at deploy time rather than using a local values file:

```bash
helm upgrade --install smart-coffee ./helm/smart-coffee \
  --set mysql.auth.rootPassword="$MYSQL_ROOT_PASSWORD" \
  --set coffeeApi.database.password="$COFFEE_DB_PASSWORD"
```

For stronger secret management, use SOPS or an external secrets solution.

## TODO

- [ ] Add k6 load tests
