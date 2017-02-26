# Scaleway dynamic inventory for Ansible

Designed to make using Ansible with Scaleway easy, this small script will
retrieve all servers associated with your account and group them by their
tags. This makes it extremely easy to associate your Ansible playbooks with
individual groups of servers (i.e. updating OpenVPN on only VPN enabled
servers).

## Installation

1. Go install this application:

`go install github.com/heraclmene/scaleway-dynamic-inventory`

2. Set the below environment variables and pass the binary reference into your 
ansible call

`SCALEWAY_ORG_TOKEN='' SCALEWAY_TOKEN='' ansible -i /path/to/gobin/scaleway-dynamic-inventory ...`

## Issues

Please raise any issues in the issue tracker associated with this project.

## Contribution guidelines

1. Ensure all code is gofmted and linted. 
2. Your build must be verified by TravisCI before the pull request will be
considered.
