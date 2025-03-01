terraform {
    backend "s3" {
     endpoint = "https://s3.timeweb.cloud"
     region = "ru-1"
     bucket = "20c5f50c-d2991433-b1b5-47e2-81ef-f4b32d64e8c8"
     key = "fae-lab2-service-mesh.tfstate"
     access_key = "GLPOITH0ULTRQDZC0TXI"
     secret_key = "bTYcnjHn3laS9sjjxzirPF4VDjvCk5zVxisiYKqt"
     skip_region_validation = true
     skip_credentials_validation = true
     skip_metadata_api_check = true
    }
}