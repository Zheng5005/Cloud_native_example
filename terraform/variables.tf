variable "kubeconfig_path" {
  description = "Ruta al kubeconfig local"
  type        = string
}

variable "namespace" {
  description = "Namespace a crear"
  type        = string
  default     = "calc-tf"
}

variable "image_server" {
  description = "Imagen del servidor gRPC"
  type        = string
}
