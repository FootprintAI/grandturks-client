# Releasing

This module **is** the kafeido CLI: grandturks builds its CLI's root command
from here, so a regression in this repository is one a customer sees directly.
Its releases are therefore meant to be consumed as releases — a tag someone
can pin, not a commit someone happened to be on.

That was not true when this document was written (grandturks-client#22), and
the reason turned out to be mechanical rather than cultural.

## What went wrong, so it is not repeated

grandturks pinned this module at

```
github.com/footprintai/grandturks-client/v2 v2.3.1-0.20260814134326-6255dd6e17fd
```

a **pseudo-version** — Go's notation for a commit with no tag on the path to
one — even though `v2.3.0+rc0` existed and was newer than anything else.

The tag was unusable. Build metadata (anything after a `+`) is not permitted in
a Go module version, so the proxy does not see such a tag as a version at all:

```
$ go list -m github.com/footprintai/grandturks-client/v2@v2.3.0+rc0
github.com/footprintai/grandturks-client/v2 v2.3.1-0.20241023162104-d0429b3a9b13
```

Nobody chose pseudo-versions. The newest tag simply could not be pinned, and Go
fell back — silently, which is why it went unnoticed for two years.

**A release candidate is `v2.4.0-rc.0`, never `v2.4.0+rc0`.** The first is a
pre-release: legal, pinnable, and sorts before `v2.4.0`. The second is build
metadata, which semver ignores for precedence anyway, so it does not even sort
the way its author intended.

## The rules

- **Tags are the consumable unit.** Consumers pin tags. A pseudo-version in a
  consumer's `go.mod` is a bug report about this repository, not a preference.
- **Semver, with the `v` prefix and no build metadata.** Major stays at `2`
  while the import path is `/v2`; a major bump means a new import path.
  - patch — bug fixes, no contract change
  - minor — new commands, new flags, regenerated client with new operations
  - major — anything that breaks an existing caller
- **A tag asserts that CI passed at that commit.** The release lane re-runs
  build, vet and test on the tag before publishing anything, because a tag can
  point at any commit. Before there was CI at all (#15), a tag asserted
  nothing, which is why option 1 of #22 was only worth adopting alongside it.
- **The module declares its own version**, in `declaredVersion` in
  `pkg/version/version.go`. The release lane refuses to publish a tag that
  disagrees with it. Without that check a `v2.5.0` tag can ship a binary that
  reports `2.3.0+rc0` — which is precisely what `kafeido version` printed for
  two years.
- **Tagging is a human decision.** No workflow creates tags.

## Cutting a release

1. **Bump `declaredVersion`** in `pkg/version/version.go` to the version being
   released, and add its section to `CHANGELOG.md`. Open this as an ordinary
   PR; the version bump is a reviewable change like any other.
2. **Merge it** and wait for `ci` to pass on `main`.
3. **Tag the merge commit** and push:

   ```bash
   git checkout main && git pull
   go run ./pkg/version/main         # prints the tag this tree expects, e.g. v2.4.0
   git tag -a v2.4.0 -m "v2.4.0"
   git push origin v2.4.0
   ```

4. **The `release` workflow does the rest**: it validates the tag shape,
   checks it against `declaredVersion`, re-runs build/vet/test at that commit,
   and publishes a GitHub Release. A tag with a pre-release suffix is published
   as a pre-release.

   These are **library releases** - a tag grandturks can pin. No CLI binaries
   are attached, because this module cannot build a working one: the kafeido
   CLI is assembled in grandturks, whose `main` injects the oauth2 callback
   decryptor (the key is shared with the authentication service) and registers
   the storage subcommands. See #33.

   Publishing the CLI itself is grandturks' job. It used to happen nowhere -
   the build existed and no lane ran it, so the newest CLI obtainable was from
   2024 (FootprintAI/grandturks#1221). Since that was fixed, every grandturks
   release attaches the CLI binaries and `copy-biries.sh` fetches them.
5. **Move the consumer.** In grandturks:

   ```bash
   go get github.com/footprintai/grandturks-client/v2@v2.4.0
   ```

   The point of all of the above is that this line names a release, and that
   "what changed" is its release notes rather than a commit range.

## Verifying a release afterwards

```bash
go list -m github.com/footprintai/grandturks-client/v2@v2.4.0
```

If that prints a pseudo-version rather than the tag you pushed, the tag is not
pinnable and the release did not happen, whatever the GitHub UI says.
