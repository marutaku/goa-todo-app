data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./adapter/repository/",
    "--dialect", "postgres"
  ]
}

variable "dsn" {
  type        = string
  description = "The name of the tenant (schema) to create"
  default     = ""
}

env "gorm" {
  src = data.external_schema.gorm.url
  url = "${var.dsn}"
  dev = "docker://postgres/15/dev?search_path=public"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}