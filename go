#!/bin/bash

set -e

VBoxManage -v >/dev/null 2>&1 || { echo >&2 "VirtualBox is required. Please install the latest version."; exit 1; }
vagrant -v >/dev/null 2>&1 || { echo >&2 "Vagrant is required. Please install the latest version."; exit 1; }
python -V >/dev/null 2>&1 || { echo >&2 "Python is required. Please install the latest version."; exit 1; }

[[ $(vagrant plugin list) == *vagrant-vbguest* ]] || { vagrant plugin install vagrant-vbguest; }

REQUIRED_RUBY=2.1.2

(rbenv versions | grep $REQUIRED_RUBY) || rbenv install $REQUIRED_RUBY
rbenv local $REQUIRED_RUBY
(rbenv exec gem list | grep bundler) || rbenv exec gem install bundler
bundle --path=vendor/bundle --quiet

ansible-galaxy install --role-file=rolefile --roles-path=roles --force

function helptext {
    echo "Usage: ./go <command>"
    echo ""
    echo "Available commands are:"
    echo "    setup        Install project dependencies"
    echo "    spec         Run acceptance tests against the local virtual machine"
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

function destroy {
    vagrant destroy -f
}

function execute {
    ansible-playbook artifactory.yml -i hosts --private-key ~/.ssh/ec2.pem -u ec2-user
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
    deploy
    spec
}

[[ $@ ]] || { helptext; exit 1; }

case "$1" in
    boot) boot
    ;;
    deploy) deploy
    ;;
    destroy) destroy
    ;;
    execute) execute
    ;;
    help) helptext
    ;;
    precommit) precommit
    ;;
    setup) setup
    ;;
    spec) spec
    ;;
esac