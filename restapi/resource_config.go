package restapi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigCreate,
		Read:   resourceConfigRead,
		Update: resourceConfigUpdate,
		Delete: resourceConfigDelete,
		Schema: map[string]*schema.Schema{
			"subject": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"config": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConfigCreate(d *schema.ResourceData, m interface{}) error {
	client, err := make_client_config(d, m)

	if err != nil {
		return err
	}

	log.Printf("Create config '%s'.", client)

	err = client.create_config()

	if err != nil {
		return err
	}

	d.SetId(client.subject)
	d.Set("subject", client.subject)
	d.Set("config", client.config)
	return resourceConfigRead(d, m)
	//return nil
}

func resourceConfigRead(d *schema.ResourceData, m interface{}) error {
	client, err := make_client_config(d, m)

	if err != nil {
		return err
	}

	log.Printf("Read config '%s'.", client)

	configResponse, err2 := client.read_config()
	if err2 != nil {
		if strings.Contains(err2.Error(), "response code is 404") {
			d.SetId("")
		} else {
			return err2
		}
	}

	if err2 == nil {
		d.Set("config", configResponse)
	}

	return nil
}

func resourceConfigUpdate(d *schema.ResourceData, m interface{}) error {
	client, err := make_client_config(d, m)

	if err != nil {
		return err
	}

	log.Printf("Update subject '%s'.", client)

	err = client.update_config()

	if err == nil {
		d.Set("config", client.config)
	}
	if err != nil {
		return err
	}

	return resourceConfigRead(d, m)
	//return nil
}

func resourceConfigDelete(d *schema.ResourceData, m interface{}) error {
	client, err := make_client_config(d, m)

	if err != nil {
		return err
	}

	log.Printf("Delete subject '%s'.", client)

	err = client.delete_config()

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
