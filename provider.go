package main

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mitchellh/mapstructure"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: "",
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: "",
			},

			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: "v1",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"kubernetes_replication_controller": resourceKubernetesReplicationController(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config Config
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &config); err != nil {
		return nil, err
	}
	log.Printf("[INFO] Initializing Kubernetes client")
	return config.Client()
}
