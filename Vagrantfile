# -*- mode: ruby -*-
# vi: set ft=ruby :
$script = <<-SCRIPT
wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz
tar -xzf go1.19.1.linux-amd64.tar.gz
mv go /usr/local/
rm -rf go1.19.1.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
apt-get install gcc -y
SCRIPT

Vagrant.configure("2") do |config|
  config.vm.box = "hashicorp/bionic64"

  config.vm.network "private_network", type: "dhcp"

  config.vm.provider "virtualbox" do |vb|
    vb.name = "docker-lab"
    vb.memory = "2048"
    vb.cpus = 2
  end
  config.vm.provision "docker" do |d|
    d.pull_images "alpine:3.16.2"
    d.pull_images "ubuntu:20.04"
    d.pull_images "nginx:1.23.1"
    d.pull_images "registry:2"
  end

  config.vm.provision "shell", inline: $script
end
