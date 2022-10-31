package restapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type schema_registry_client_config struct {
	http_client *http.Client
	create_uri  string
	update_uri  string
	read_uri    string
	delete_uri  string
	global_uri  string
	subject     string
	config      string
}

func NewSchemaRegistryClientConfig(uri string, subject string, config string) (*schema_registry_client_config, error) {
	client := schema_registry_client_config{
		create_uri:  uri + "/config/" + subject,
		update_uri:  uri + "/config/" + subject,
		read_uri:    uri + "/config/" + subject,
		delete_uri:  uri + "/config/" + subject,
		global_uri:  uri + "/config/",
		subject:     subject,
		config:      config,
		http_client: &http.Client{},
	}

	return &client, nil
}

func (client schema_registry_client_config) create_config() error {

	// set the HTTP method, url, and request body
	request, err := http.NewRequest(http.MethodPut, client.create_uri, strings.NewReader(client.config))
	if err != nil {
		return err
	}

	// set the request header Content-Type for json
	request.Header.Set("Content-Type", "application/vnd.schemaregistry.v1+json")
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

	return nil
}

func (client schema_registry_client_config) update_config() error {

	// set the HTTP method, url, and request body
	request, err := http.NewRequest(http.MethodPut, client.create_uri, strings.NewReader(client.config))
	if err != nil {
		return err
	}

	// set the request header Content-Type for json
	request.Header.Set("Content-Type", "application/vnd.schemaregistry.v1+json")
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

	return nil
}

func (client schema_registry_client_config) read_config() (*string, error) {
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
	responseBodyString := strings.Replace(string(data), "compatibilityLevel", "compatibility", 1)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
		return nil, err
	}

	defer response.Body.Close()

	return &responseBodyString, nil
}

// delete_config -confluence < 7.0.0
// Get the global config and set it to the schema before delete the state.
func (client schema_registry_client_config) delete_config() error {

	request, err := http.NewRequest("GET", client.global_uri, nil)
	if err != nil {
		return err
	}
	//request.Header.Add("Content-Type", "application/vnd.schemaregistry.v1+json")
	response, err := client.http_client.Do(request)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	responseBodyString := strings.Replace(string(data), "compatibilityLevel", "compatibility", 1)

	//log.Println("Body string:::::" + responseBodyString)

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
		return err
	}

	defer response.Body.Close()
	// Get the content return for the Get response client.global_uri

	request, err = http.NewRequest(http.MethodPut, client.create_uri, strings.NewReader(responseBodyString))
	if err != nil {
		return err
	}

	// set the request_put header Content-Type for json
	request.Header.Set("Content-Type", "application/vnd.schemaregistry.v1+json")
	response, err = client.http_client.Do(request)
	if err != nil {
		return err
	}

	data, err = ioutil.ReadAll(response.Body)
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

/* delete_config if confluence schema registry > 7.0.0
func (client schema_registry_client_config) delete_config() error {
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
*/
