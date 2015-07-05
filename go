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
    echo "    setup        Install project dependencies"
    echo "    boot         Spin up a local virtual machine"
    echo "    deploy       Deploy all components to the local virtual machine"
    echo "    destroy      Destroy the local virtual machine"
    echo "    precommit    Run all validations before pushing code"
}

function boot {
    vagrant up --no-provision
}

function deploy {
    boot
    vagrant provision
}

function setup {
    which pip >/dev/null 2>&1 || sudo easy_install pip
    ansible --version  >/dev/null 2>&1 || sudo pip install ansible
    ansible-galaxy install --role-file=rolefile --roles-path=roles --force
}

function precommit {
    setup
    deploy
}

[[ $@ ]] || { helptext; exit 1; }

case "$1" in
    boot) boot
    ;;
    deploy) deploy
    ;;
    destroy) vagrant destroy -f
    ;;
    help) helptext
    ;;
    precommit) precommit
    ;;
    setup) setup
    ;;
esac