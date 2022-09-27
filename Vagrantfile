# -*- mode: ruby -*-
# vi: set ft=ruby :
# Fix VirtualBox windows
class VagrantPlugins::ProviderVirtualBox::Action::Network
  def dhcp_server_matches_config?(dhcp_server, config)
    true
  end
end

current_dir    = File.dirname(File.expand_path(__FILE__))
servers        = YAML.load_file("#{current_dir}/config.yaml")

$script = <<-SCRIPT
wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz
tar -xzf go1.19.1.linux-amd64.tar.gz
mv go /usr/local/
rm -rf go1.19.1.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
apt-get update && apt-get install gcc -y
mkdir -p /usr/local/lib/docker/cli-plugins
curl -SL https://github.com/docker/compose/releases/download/v2.11.1/docker-compose-linux-x86_64 -o /usr/local/lib/docker/cli-plugins/docker-compose
chmod +x /usr/local/lib/docker/cli-plugins/docker-compose
SCRIPT

$daemon = <<-SCRIPT
echo '{ "insecure-registries":["#{servers[0]["ip"]}:5000"] }' > /etc/docker/daemon.json
systemctl restart docker
SCRIPT

Vagrant.configure("2") do |config|

  config.vm.box = "hashicorp/bionic64"
  
  servers.each do |servers|
    config.vm.define servers["name"] do |srv|
      srv.vm.network "private_network", ip: servers["ip"]
      case servers["name"]
        when "lab"
              srv.vm.provision "shell", inline: $script
              srv.vm.provision "docker" do |d|
                d.pull_images "alpine:3.16.2"
                d.pull_images "ubuntu:20.04"
                d.pull_images "nginx:1.23.1"
                d.pull_images "registry:2"
              end
        when "client"
              srv.vm.provision "docker"
              srv.vm.provision "shell", inline: $daemon
      end
      srv.vm.provider :virtualbox do |vb|
        vb.name = servers["name"]
        vb.memory = servers["ram"]
        vb.cpus = servers["cpus"]
      end
    end
  end
end
