#!/usr/bin/env bash
#
# Fetches the kafeido CLI.
#
# This repository is the CLI's source, not its distribution: a binary built
# from this module alone has no oauth2 callback decryptor and none of the
# storage subcommands, because grandturks' main is what injects them (#33).
# The CLI you run is built and published from grandturks, so this script
# fetches it from there.
#
# It used to `docker pull footprintai/grandturks-kafeido-client:v2.2.1` - a
# 2024 image, from a registry nothing publishes to any more. Nothing had
# published a CLI at all in between: the build existed and no lane ran it
# (FootprintAI/grandturks#1221). Since that was fixed, every grandturks
# release carries the binaries below, and the image goes to the same private
# registry as the service images.
#
# ACCESS: grandturks is a private repository, so this needs a GitHub account
# with read access to it, and `gh auth login` done once. If you have no such
# access, ask the team - there is no public download today, which is a known
# gap rather than an oversight.
set -euo pipefail

REPO="${REPO:-FootprintAI/grandturks}"
# Unset means the latest release. Set TAG=v2.4.4 to pin one.
TAG="${TAG:-}"
DEST="${DEST:-./kafeido}"

command -v gh >/dev/null || { echo "gh is required: https://cli.github.com" >&2; exit 1; }

case "$(uname -s)" in
  Linux)  os=linux ;;
  Darwin) os=darwin ;;
  *)      echo "unsupported OS $(uname -s) - fetch the windows asset by hand" >&2; exit 1 ;;
esac
case "$(uname -m)" in
  x86_64|amd64)  arch=amd64 ;;
  arm64|aarch64) arch=arm64 ;;
  *)             echo "unsupported architecture $(uname -m)" >&2; exit 1 ;;
esac

if [ -z "${TAG}" ]; then
  TAG="$(gh release view --repo "${REPO}" --json tagName -q .tagName)"
fi

mkdir -p "${DEST}"
cd "${DEST}"

asset="kafeido-${TAG}-${os}-${arch}"
echo "fetching ${asset} from ${REPO} ${TAG}"
gh release download "${TAG}" --repo "${REPO}" --pattern "${asset}" --pattern SHA256SUMS --clobber || true

# Checked rather than trusted: `gh release download` prints "no assets to
# download" and exits 0 when a release carries none, so set -e does not catch
# it and the script would sail on to checksum a file that is not there.
if [ ! -f "${asset}" ] || [ ! -f SHA256SUMS ]; then
  echo "release ${TAG} of ${REPO} carries no CLI assets." >&2
  echo "Releases before FootprintAI/grandturks#1225 predate the lane that attaches them;" >&2
  echo "pin a newer one with TAG=vX.Y.Z, or check that ${asset} exists on that release." >&2
  exit 1
fi

# The checksums cover every platform's asset; check only the one downloaded,
# since --ignore-missing is what keeps the other four from failing the check.
sha256sum --ignore-missing --check SHA256SUMS

chmod +x "${asset}"
ln -sf "${asset}" kafeido
echo "installed ${DEST}/${asset} (symlinked as ${DEST}/kafeido)"
"./${asset}" version
