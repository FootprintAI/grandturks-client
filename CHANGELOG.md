## v2.7.0 ##

Step 2 of moving the product API off `/api`, which it shares with the Kubeflow
central dashboard (FootprintAI/manifests#308). Nothing moves in this release:
the default is unchanged, and the point is to make the move possible at all.

* feat(cli): the API base path is configurable - `endpoint.apiBasePath` in the
  config file, `KAFEIDO_API_BASE_PATH` in the environment, or
  `config set endpoint --api_base_path`. It was the literal `"api"` compiled
  into root.go, so no deployment could serve the API anywhere else and no
  binary already in the field could be pointed at a new prefix. **The default
  is unchanged**: every existing config, scripted install and released binary
  keeps talking to `/api` (FootprintAI/manifests#308)
* feat(cli): every request now carries
  `User-Agent: kafeido-cli/<version> (<commit>)`. Retiring the old prefix
  blind would break whoever is still on it, and this repository had no way to
  report which versions were deployed - the same gap that made
  FootprintAI/grandturks#1251's committed key unrotatable. Istio's ingress
  access log already records the User-Agent *and* the pre-rewrite path, so one
  existing log now answers both "which prefix" and "which client", with no
  server-side change
* docs: fetch the CLI from where it is now published (#42)
* docs: drop three references to an issue that is about something else (#41)

## v2.6.0 ##

* feat(cli): the oauth2 login now asks for, and reads, an authenticated
  callback credential (#29, steps 1-3). The CLI mints an ephemeral X25519
  keypair per login and sends the public half; a server that understands it
  seals the access token with AES-256-GCM, and one that does not returns the
  legacy AES-CBC blob exactly as before. The format marker decides which
  decoder runs, so no client and server have to be upgraded together.
  Requires FootprintAI/grandturks#1210 server-side to take effect.
* feat(encryption): HasCredentialMarker, for dispatching on the format rather
  than on well-formedness - a truncated blob claiming the marker is now
  rejected as a sealed credential instead of being handed to the legacy
  decoder (#29)

## v2.5.0 ##

* feat(encryption): authenticated oauth2 callback credential - ephemeral
  X25519 + AES-256-GCM in a version-marked format, replacing unauthenticated
  AES-CBC under a build-time key (#29, step 1 of the rollout in
  docs/architecture/oauth2-callback-credential.md). Nothing calls it yet:
  the login command and the authentication service move in later steps.
* fix(encryption): Encryption reports "no encryptor configured" instead of
  dereferencing a nil interface (#33)
* fix(release): releases are library releases; no CLI binaries are attached,
  because this module cannot build a working one (#33)
* docs(architecture): the design for #29

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
