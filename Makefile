SHELL := /bin/bash
IMG_PREFIX ?= local
KIND_CLUSTER ?= calc

PROTO_DIR := proto
GEN_DIR := gen

.PHONY: proto server client docker-build kind-up k8s-apply k8s-run-client

proto:
	@protoc --go_out=$(GEN_DIR) --go-grpc_out=$(GEN_DIR) \
		$(PROTO_DIR)/calculator.proto
	@echo "✅ Stubs generados en $(GEN_DIR)"

server:
	go build -o bin/server ./server

client:
	go build -o bin/client ./client

docker-build:
	docker build -t $(IMG_PREFIX)/grpc-server:latest -f server/Dockerfile .
	docker build -t $(IMG_PREFIX)/grpc-client:latest -f client/Dockerfile .
	@if [ "$(IMG_PREFIX)" = "local" ]; then \
		echo "📦 Cargando imágenes en kind..."; \
		kind load docker-image $(IMG_PREFIX)/grpc-server:latest --name $(KIND_CLUSTER); \
		kind load docker-image $(IMG_PREFIX)/grpc-client:latest --name $(KIND_CLUSTER); \
	fi

kind-up:
	kind create cluster --name $(KIND_CLUSTER) --config kind/kind-config.yaml || true

k8s-apply:
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/server-deployment.yaml
	kubectl apply -f k8s/server-service.yaml

k8s-run-client:
	kubectl -n calc delete job grpc-client --ignore-not-found
	kubectl apply -f k8s/client-job.yaml

k8s-see-result:
	kubectl -n calc logs job/grpc-client -f

k8s-destroy:
	kubectl delete ns calc --ignore-not-found
	kind delete cluster --name calc

.PHONY: all-local

all-local: proto kind-up docker-build k8s-apply
	@echo "✅ Flujo completo ejecutado correctamente"

