terraform {
  required_providers {
    stdlib = {
      source  = "mschuchard/stdlib"
      version = "~> 2.0"
    }
  }
}

provider "stdlib" {} # can be omitted
