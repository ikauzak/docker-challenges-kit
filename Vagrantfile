# -*- mode: ruby -*-
# vi: set ft=ruby :
# Fix VirtualBox windows
class VagrantPlugins::ProviderVirtualBox::Action::Network
  def dhcp_server_matches_config?(dhcp_server, config)
    true
  end
end

current_dir    = File.dirname(File.expand_path(__FILE__))
servers        = YAML.load_file("#{current_dir}/vagrant_config.yaml")

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

      srv.vm.provision "shell",
        inline: "make -C /vagrant start"

      srv.vm.provider :virtualbox do |vb|
        vb.name = servers["name"]
        vb.memory = servers["ram"]
        vb.cpus = servers["cpus"]
      end

    end
  end
end
