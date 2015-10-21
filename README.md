# terraform-provider-kubernetes

Example usage:

provider "kubernetes" {
    endpoint = "http://kube.domain.test.com:8080"
}

resource "template_file" "some_service_json_file" {
  filename = "service.json"
  lifecycle { create_before_destroy = true }
}

resource "kubernetes_service" "terraform-service-resource" {
    name = "servtest"
    config = "${template_file.user_data.rendered}"
}

service.json example:

{
    "metadata":
    {
        "name": "test",
        "labels":
        {
            "k8s-app": "test"
        }
    },
    "spec":
    {
        "selector":
        {
            "k8s-app": "test-select"
        },
        "ports":[
            {
            "port": 80,
            "targetPort": 8080
            }
        ]
    }
}
