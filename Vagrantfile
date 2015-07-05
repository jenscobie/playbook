# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "chef/centos-7.1"

  config.vm.network "forwarded_port", guest: 8081, host: 8081

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "artifactory.yml"
  end
end
