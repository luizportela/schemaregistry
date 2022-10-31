package restapi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSubject() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubjectCreate,
		Read:   resourceSubjectRead,
		Update: resourceSubjectUpdate,
		Delete: resourceSubjectDelete,

		Schema: map[string]*schema.Schema{
			"subject": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schema": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					oldString := "{\"type\":" + old + "}"
					return oldString == new
				},
			},
		},
	}
}

func resourceSubjectCreate(d *schema.ResourceData, m interface{}) error {
	client, err := make_client_schema(d, m)

	if err != nil {
		return err
	}

	log.Printf("Create subject '%s'.", client)

	err = client.create_subject()

	if err != nil {
		return err
	}

	d.SetId(client.subject)
	d.Set("subject", client.subject)
	d.Set("schema", client.schema)
	return resourceSubjectRead(d, m)
	//return nil
}

func resourceSubjectRead(d *schema.ResourceData, m interface{}) error {
	client, err := make_client_schema(d, m)

	if err != nil {
		return err
	}

	log.Printf("Read config '%s'.", client)

	schemaResponse, err2 := client.read_config()
	if err2 != nil {
		if strings.Contains(err2.Error(), "response code is 404") {
			d.SetId("")
		} else {
			return err2
		}
	}

	if err2 == nil {
		d.SetId(schemaResponse.Subject)
		d.Set("subject", schemaResponse.Subject)
		d.Set("schema", schemaResponse.Schema)
	}

	return nil
}

func resourceSubjectUpdate(d *schema.ResourceData, m interface{}) error {
	client, err := make_client_schema(d, m)

	if err != nil {
		return err
	}

	log.Printf("Update subject '%s'.", client)

	err = client.update_subject()

	if err == nil {
		d.Set("schema", client.schema)
	}
	if err != nil {
		return err
	}

	return resourceSubjectRead(d, m)
	//return nil
}

func resourceSubjectDelete(d *schema.ResourceData, m interface{}) error {
	client, err := make_client_schema(d, m)

	if err != nil {
		return err
	}

	log.Printf("Delete subject '%s'.", client)

	err = client.delete_subject()

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
