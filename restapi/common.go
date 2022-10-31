package restapi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func make_client_schema(d *schema.ResourceData, m interface{}) (*schema_registry_client_schema, error) {
	uri := m.(string)
	subject := d.Get("subject").(string)
	schema := d.Get("schema").(string)

	client, err := NewSchemaRegistryClientSchema(uri, subject, schema)

	return client, err
}

func make_client_config(d *schema.ResourceData, m interface{}) (*schema_registry_client_config, error) {
	uri := m.(string)
	subject := d.Get("subject").(string)
	config := d.Get("config").(string)

	client, err := NewSchemaRegistryClientConfig(uri, subject, config)

	return client, err
}
