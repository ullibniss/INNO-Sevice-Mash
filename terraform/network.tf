resource "twc_vpc" "mesh-vpc" {
  name = "Mesh VPC"
  description = "VPC for service mash"
  subnet_v4 = "192.168.7.0/24"
  location = "ru-1"
}