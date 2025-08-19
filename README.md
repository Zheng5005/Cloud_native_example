# Cloud-Native gRPC Calculator (Go)

Demostración **cloud-native**: microservicio gRPC en Go (servidor + cliente), contenedores Docker, despliegue en Kubernetes, e **IaC** con Terraform (usando el provider `kubernetes`) + cluster local con **kind**.

## Estructura
```
cloud-native-grpc-calculator/
├─ client/
│  ├─ Dockerfile
│  └─ main.go
├─ server/
│  ├─ Dockerfile
│  └─ main.go
├─ proto/
│  └─ calculator.proto
├─ gen/              # (se debe crear primero y luego ejecutar 'make proto')
│  └─ ... 
├─ k8s/
│  ├─ namespace.yaml
│  ├─ server-deployment.yaml
│  ├─ server-service.yaml
│  └─ client-job.yaml
├─ kind/
│  └─ kind-config.yaml
├─ terraform/
│  ├─ provider.tf
│  ├─ main.tf
│  ├─ variables.tf
│  └─ outputs.tf
├─ .github/workflows/
│  └─ ci.yaml
├─ Makefile
└─ go.mod
```

## Requisitos
- Go 1.22+
- `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc`
- Docker
- kind (o Minikube)
- kubectl
- Terraform (para IaC)

## Comandos rápidos

```bash
# 1) Generar stubs gRPC
make proto

# 2) Compilar binarios (opcional, permite testear el server y cliente sin Kubernetes)
make server client

# 3) Levantar cluster kind (local)
make kind-up

# 4) Construir imágenes Docker y cargarlas en kind
make docker-build IMG_PREFIX=local

# 5) Desplegar en Kubernetes
make k8s-apply

# 6) Ejecutar cliente como Job en el cluster
make k8s-run-client

# 7) Limpieza
make k8s-destroy

# Paso 100% opcional
# 8) Terraform (IaC) - namespace + despliegue alternativo -
cd terraform
terraform init
terraform apply -auto-approve   -var="kubeconfig_path=$HOME/.kube/config"   -var="namespace=calc-tf"   -var="image_server=tuusuario/grpc-server:latest"
```
## Ejecución local (sin Kubernetes)
En una terminal:
```bash
go run ./server
```
En otra terminal:
```bash
go run ./client add 3 5
go run ./client sub 3 5
go run ./client mul 3 5
go run ./client div 3 5
```

## Endpoints gRPC
- `Calculator.Add(a, b)` -> `result`
- `Calculator.Sub(a, b)` -> `result`
- `Calculator.Mul(a, b)` -> `result`
- `Calculator.Div(a, b)` -> `result`
