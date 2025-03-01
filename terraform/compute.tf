data "twc_presets" "mesh-preset-1-2-20" {
  location = "ru-1"
  disk_type = "nvme"
  disk = 1024 * 30
  cpu = 1
  cpu_frequency = 3.3
  ram = 1024 * 2
}

data "twc_os" "mesh-os" {
  name = "ubuntu"
  version = "22.04"
}

resource "twc_server" "master-lb-server" {
  name = "Master Load Balancer"
  os_id = data.twc_os.mesh-os.id
  preset_id = data.twc_presets.mesh-preset-1-2-20.id
  cloud_init = templatefile("./cloud-init/others.yaml", {ssh_pass: local.master-lb-pass , hostname: "master-lb" })

  local_network {
    id = twc_vpc.mesh-vpc.id
    ip = "192.168.7.100"
  }
}

resource "twc_server" "coffee-server" {
  name = "Coffee"
  os_id = data.twc_os.mesh-os.id
  preset_id = data.twc_presets.mesh-preset-1-2-20.id
  cloud_init = templatefile("./cloud-init/coffee.yaml", {ssh_pass: local.coffee-pass, hostname: "coffee" })

  local_network {
    id = twc_vpc.mesh-vpc.id
    ip = "192.168.7.101"
  }
}

resource "twc_floating_ip" "coffee-floating-ip" {
  availability_zone = "spb-3"

  comment = "Coffee floating ip"

  resource {
    type = "server"
    id   = twc_server.coffee-server.id
  }
}

resource "twc_server" "drink-server" {
  name = "Drink"
  os_id = data.twc_os.mesh-os.id
  preset_id = data.twc_presets.mesh-preset-1-2-20.id
  cloud_init = templatefile("./cloud-init/others.yaml", {ssh_pass: local.drink-pass, hostname: "drink" })

  local_network {
    id = twc_vpc.mesh-vpc.id
    ip = "192.168.7.102"
  }
}

resource "twc_server" "milk-server" {
  name = "Milk"
  os_id = data.twc_os.mesh-os.id
  preset_id = data.twc_presets.mesh-preset-1-2-20.id
  cloud_init = templatefile("./cloud-init/others.yaml", {ssh_pass: local.milk-pass, hostname: "milk" })

  local_network {
    id = twc_vpc.mesh-vpc.id
    ip = "192.168.7.103"
  }
}

resource "twc_server" "syrup-server" {
  name = "Syrup"
  os_id = data.twc_os.mesh-os.id
  preset_id = data.twc_presets.mesh-preset-1-2-20.id
  cloud_init = templatefile("./cloud-init/others.yaml", {ssh_pass: local.syrup-pass, hostname: "syrup" })

  local_network {
    id = twc_vpc.mesh-vpc.id
    ip = "192.168.7.104"
  }
}