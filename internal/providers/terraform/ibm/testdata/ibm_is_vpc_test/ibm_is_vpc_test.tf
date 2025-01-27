
terraform {
  required_providers {
    ibm = {
      source = "IBM-Cloud/ibm"
      version = "~> 1.12.0"
    }
  }
}

provider "ibm" {
    region = "us-south"
}

resource "ibm_is_vpc" "testVpc" {
  name = "test-vpc"
}
