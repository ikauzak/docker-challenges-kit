# -*- mode: ruby -*-
# vi: set ft=ruby :
# Fix VirtualBox windows
class VagrantPlugins::ProviderVirtualBox::Action::Network
  def dhcp_server_matches_config?(dhcp_server, config)
    true
  end
end

current_dir    = File.dirname(File.expand_path(__FILE__))
servers        = YAML.load_file("#{current_dir}/.vagrant_config.yaml")

$lab = <<-SCRIPT
wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz
tar -xzf go1.19.1.linux-amd64.tar.gz
mv go /usr/local/
rm -rf go1.19.1.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
apt-get update && apt-get install gcc -y
mkdir -p /usr/local/lib/docker/cli-plugins
curl -SL https://github.com/docker/compose/releases/download/v2.11.1/docker-compose-linux-x86_64 -o /usr/local/lib/docker/cli-plugins/docker-compose
chmod +x /usr/local/lib/docker/cli-plugins/docker-compose
cd /vagrant && make start
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
      
      srv.vm.hostname = "#{servers["name"]}.local"
      
      srv.vm.provision "docker" do |d|
        servers["images"].each do |image|
          d.pull_images "#{image}"
        end
      end
      
      case servers["name"]
        when "lab"
          srv.vm.provision "shell", inline: $lab
        when "client"
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
