terraform {
  required_providers {
    stdlib = {
      source  = "mschuchard/stdlib"
      version = "~> 1.0"
    }
  }
}

provider "stdlib" {} # can be omitted
