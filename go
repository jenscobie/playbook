#!/bin/bash

set -e

VBoxManage -v >/dev/null 2>&1 || { echo >&2 "VirtualBox is required. Please install the latest version."; exit 1; }
vagrant -v >/dev/null 2>&1 || { echo >&2 "Vagrant is required. Please install the latest version."; exit 1; }
python -V >/dev/null 2>&1 || { echo >&2 "Python is required. Please install the latest version."; exit 1; }

[[ $(vagrant plugin list) == *vagrant-vbguest* ]] || { vagrant plugin install vagrant-vbguest; }

function helptext {
    echo "Usage: ./go <command>"
    echo ""
    echo "Available commands are:"
    echo "    boot         Spin up a local virtual machine"
    echo "    destroy      Destroy the local virtual machine"
    echo "    execute      Provision EC2 servers"
    echo "    precommit    Run all validations before pushing code"
    echo "    provision    Provision the local virtual machine"
    echo "    setup        Install project dependencies"
    echo "    spec         Run acceptance tests against the local virtual machine"
}

function boot {
    vagrant up --no-provision
}

function provision {
    ansible-galaxy install --role-file=Galaxyfile --roles-path=roles --force

    boot
    vagrant provision
}

function destroy {
    vagrant destroy -f
}

function execute {
    ansible-playbook site.yml -i hosts --private-key ~/.ssh/ec2.pem -u ec2-user
}

function setup {
    which pip >/dev/null 2>&1 || sudo easy_install pip
    ansible --version  >/dev/null 2>&1 || sudo pip install ansible
}

function spec {
    bundle exec rake spec
}

function precommit {
    setup
    provision
    spec
}

[[ $@ ]] || { helptext; exit 1; }

case "$1" in
    boot) boot
    ;;
    destroy) destroy
    ;;
    execute) execute
    ;;
    help) helptext
    ;;
    precommit) precommit
    ;;
    provision) provision
    ;;
    setup) setup
    ;;
    spec) spec
    ;;
esac