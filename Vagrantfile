# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "chef/centos-7.1"

  config.vm.define "goserver" do |goserver|
    goserver.vm.network "private_network", ip: "192.168.50.2"
    goserver.vm.network "forwarded_port", guest: 8153, host: 8153
  end

  config.vm.define "artifactory" do |artifactory|
    artifactory.vm.network "private_network", ip: "192.168.50.4"
    artifactory.vm.network "forwarded_port", guest: 8081, host: 8081
  end

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "site.yml"
    ansible.groups = {
        "buildservers" => ["goserver"]
    }
  end
end
