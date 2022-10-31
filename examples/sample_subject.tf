terraform {
  required_providers {
    schemaregistry = {
      source = "github.com/luizportela/schemaregistry"
      version = "1.0.1"
    }
  }
}

provider "schemaregistry" {
  uri = "http://localhost:8081"
}

resource "schemaregistry_subject" "schema_sample_from_string" {
  subject = "com.test.myapp.test-from-string"
  schema  = "{\"type\":\"record\",\"name\":\"paymentstest\",\"namespace\":\"my.examples\",\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"double\"}]}"
}

resource "schemaregistry_config" "schema_sample_from_string_config" {
  subject = "com.test.myapp.test-from-string"
  config  = "{\"compatibility\":\"BACKWARD\"}"
}

resource "schemaregistry_config" "schema_sample_config_2" {
  subject = "com.test.myapp.test-from-string-2"
  config  = "{\"compatibility\":\"FORWARD\"}"
}

resource "schemaregistry_subject" "schema_sample_key" {
  subject = "com.test.myapp-key"
  schema  ="{\"type\":\"long\"}"
}