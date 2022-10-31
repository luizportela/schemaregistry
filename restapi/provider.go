package restapi

import (
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"uri": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Kafka schema registry endpoint. Example: http://localhost:8000",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"schemaregistry_subject": resourceSubject(),
			"schemaregistry_config":  resourceConfig(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("uri").(string)
	_, err := url.ParseRequestURI(endpoint)

	return endpoint, err
}
