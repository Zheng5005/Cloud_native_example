resource "kubernetes_namespace" "ns" {
  metadata {
    name = var.namespace
  }
}

resource "kubernetes_deployment" "server" {
  metadata {
    name      = "grpc-server"
    namespace = kubernetes_namespace.ns.metadata[0].name
    labels = {
      app = "grpc-server"
    }
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        app = "grpc-server"
      }
    }
    template {
      metadata {
        labels = {
          app = "grpc-server"
        }
      }
      spec {
        container {
          name  = "server"
          image = var.image_server
          port {
            container_port = 50051
          }
          env {
            name  = "LISTEN_ADDR"
            value = ":50051"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "svc" {
  metadata {
    name      = "grpc-server"
    namespace = kubernetes_namespace.ns.metadata[0].name
  }
  spec {
    selector = {
      app = kubernetes_deployment.server.spec[0].template[0].metadata[0].labels.app
    }
    port {
      name        = "grpc"
      port        = 50051
      target_port = 50051
    }
    type = "ClusterIP"
  }
}
