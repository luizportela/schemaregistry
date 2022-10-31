package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type schema_registry_client_schema struct {
	http_client *http.Client
	create_uri  string
	update_uri  string
	read_uri    string
	delete_uri  string
	subject     string
	schema      string
}

type Reference struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Version int    `json:"version"`
}

type schemaResponse struct {
	Subject    string      `json:"subject"`
	Version    int         `json:"version"`
	Schema     string      `json:"schema"`
	ID         int         `json:"id"`
	References []Reference `json:"references"`
}

func NewSchemaRegistryClientSchema(uri string, subject string, schema string) (*schema_registry_client_schema, error) {
	client := schema_registry_client_schema{
		create_uri:  uri + "/subjects/" + subject + "/versions",
		update_uri:  uri + "/subjects/" + subject + "/versions",
		read_uri:    uri + "/subjects/" + subject + "/versions" + "/latest",
		delete_uri:  uri + "/subjects/" + subject,
		subject:     subject,
		schema:      schema,
		http_client: &http.Client{},
	}

	return &client, nil
}

func (client schema_registry_client_schema) create_subject() error {
	jsonData := map[string]string{"schema": client.schema}
	jsonValue, err := json.Marshal(jsonData)

	if err != nil {
		return err
	}

	response, err := http.Post(client.create_uri, "application/vnd.schemaregistry.v1+json", bytes.NewBuffer(jsonValue))

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
		log.Println("ERROR happened")
		return err
	}

	defer response.Body.Close()

	return nil
}

func (client schema_registry_client_schema) delete_subject() error {
	request, err := http.NewRequest("DELETE", client.delete_uri, nil)

	if err != nil {
		return err
	}

	response, err := client.http_client.Do(request)

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
		return err
	}

	defer response.Body.Close()

	return err
}

func (client schema_registry_client_schema) read_config() (*schemaResponse, error) {
	request, err := http.NewRequest("GET", client.read_uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/vnd.schemaregistry.v1+json")
	response, err := client.http_client.Do(request)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(data)

	var schemaResponse schemaResponse
	json.Unmarshal([]byte(bodyString), &schemaResponse)

	log.Println("Body string:::::" + bodyString)
	log.Println("Body Json string:::::" + schemaResponse.Schema)

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
		return nil, err
	}

	defer response.Body.Close()

	return &schemaResponse, err
}

func (client schema_registry_client_schema) update_subject() error {
	jsonData := map[string]string{"schema": client.schema}
	jsonValue, err := json.Marshal(jsonData)

	if err != nil {
		return err
	}

	response, err := http.Post(client.create_uri, "application/vnd.schemaregistry.v1+json", bytes.NewBuffer(jsonValue))

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
		return err
	}

	defer response.Body.Close()

	return nil
}
