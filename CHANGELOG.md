## v2.4.0 ##

The first release under the policy in RELEASING.md: tags are what consumers
pin, and this one asserts that CI passed at the commit it names.

* feat(cli): create, list, revoke and authenticate with api keys (#21)
* ci: build/vet/test lane, the first CI this repository has had (#15)
* fix(encryption): restore aes_test.go, which had not compiled since 2024 (#16)
* fix(encryption): Decode returns an error rather than panicking on malformed
  ciphertext, which the oauth2 callback feeds from outside the process (#23)
* chore: tagged releases are the consumable unit; `v2.3.0+rc0` never was, since
  a Go module version may not carry build metadata (#22)

## v2.0.1 ##

#### Breaking Notice ####
This version is not going to support backward compability as the project is renamed as `kafeido`


* feat: add video datasource
* feat: support kafeido:// protocol for deploying models
* feat(kafeido): open prediction api


## v1.2.3 ##
* feat: add job cancel functions
* feat: upgrade underlying kubeflow to v1.6.0

## v1.0.0 ##

* initial release on client library
