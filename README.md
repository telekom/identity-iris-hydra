<!---
# SPDX-FileCopyrightText: 2025 Deutsche Telekom AG
#
# SPDX-License-Identifier: Apache-2.0
-->

# Iris-Hydra

Iris-Hydra is an Ory Hydra fork, with modifications required so that it can be
used as the "Iris" OAuth2.0 Authorization Server component.

Original ory-hydra [README](HYDRA_README.md)

## Branching

Active development takes place on the `iris-hydra-X.Y.Z` branch, where `X.Y.Z`
is the version of Hydra that Iris-Hydra is based on. All non-Hydra(Iris) commits
should appear AFTER the last Hydra commit in this branch.

**Important:** When switching to a new Hydra version, the branch should be
rebased on top of the new Hydra version.

**Note:** The `master` branch is read-only and tracks Hydra's master branch.

## Versioning

Iris-hydra uses the same versioning as Hydra, with the addition of a suffix to
indicate the Iris-Hydra specific part. Suffix is in the format `-iris-A.B.C`,
where `A.B.C` is the major, minor and patch version.

Example: `2.3.0-iris-0.1.0`

## Code formatting and code conventions

We follow the same formatting and code conventions used in Ory Hydra. The rules
are enforced for all pull requests by relevant GitHub actions.

To format locally prior to committing run:

```shell
make format
```

## License

Since Iris-Hydra is a fork of Hydra, it is licensed under same license, Apache
2.0.

**Important:** For new files, the license header should be added as per the
SPDX-License-Identifier tag at the top of this README file.

## Quickstart

To build and run Iris-Hydra locally, you can use the provided
`quickstart-iris-hydra` docker-compose file:

```shell
docker compose -f quickstart.yml -f quickstart-iris-hydra.yml up -d --build
```
