<!-- Title section -->
<h1 align="center">
  Kamontat's Workspace for Golang Ecosystem

  <img src="https://simpleicons.org/icons/go.svg" width="24px">
  <img src="https://simpleicons.org/icons/sonarcloud.svg" width="24px">
</h1>

<!-- Description section -->
<p align="center">
  <strong>This monorepo included libraries, tools, and helpers that will make my life easier.</strong>
</p>

<!-- Badge setup -->
<p align="center">
  <a href="https://github.com/kc-workspace/go-lib">
    <img src="https://img.shields.io/github/commit-activity/m/kc-workspace/go-lib?style=flat-square" alt="Commit activity" />
  </a>
  <a href="https://sonarcloud.io/project/overview?id=kc-workspace_go-lib">
    <img src="https://img.shields.io/sonar/quality_gate/kc-workspace_go-lib/main?server=https%3A%2F%2Fsonarcloud.io&style=flat-square" alt="Quality gate" />
  </a>
  <a href="https://sonarcloud.io/project/issues?resolved=false&id=kc-workspace_go-lib">
    <img src="https://img.shields.io/sonar/violations/kc-workspace_go-lib/main?format=long&server=https%3A%2F%2Fsonarcloud.io&style=flat-square" alt="Violations" />
  </a>
  <a href="https://sonarcloud.io/project/overview?id=kc-workspace_go-lib">
    <img src="https://img.shields.io/sonar/tech_debt/kc-workspace_go-lib/main?server=https%3A%2F%2Fsonarcloud.io&style=flat-square" alt="Tech debt" />
  </a>
  <a href="https://sonarcloud.io/component_measures?metric=Coverage&view=list&id=kc-workspace_go-lib">
    <img src="https://img.shields.io/sonar/coverage/kc-workspace_go-lib/main?server=https%3A%2F%2Fsonarcloud.io&style=flat-square" alt="Code coverage" />
  </a>
  <a href="https://dashboard.mergify.com/github/kc-workspace/repo/go-lib/queues">
  <img src="https://img.shields.io/endpoint?label=mergify&logo=-&style=flat-square&url=https%3A%2F%2Fapi.mergify.com%2Fv1%2Fbadges%2Fkc-workspace%2Fgo-lib" alt="Mergify status" />
  </a>
</p>

<!-- External section -->
<h3 align="center">
  <a href="https://sonarcloud.io/project/overview?id=kc-workspace_go-lib">SONARCLOUD</a>
  <span> · </span>
  <a href="https://pkg.go.dev/search?q=kc-workspace%2Fgo-lib">PKG</a>
</h3>

## Prerequisite

1. Run `./scripts/main onboarding` for checking current environment

## Everyday command

Below is a list of command you usually use to develop this repository.

### Test packages

- `./scripts.sh test all` - test all packages
- `COVER=1 ./scripts.sh test all` - test and get coverage for all packages
- `./scripts.sh test <package>` - test only **<package>** package
- `COVER=1 ./scripts.sh test <package>` - test and get coverage for **<package>** package

### Build packages

`./scripts.sh build`: To build libraries for any compile errors.
Usually you don't need this as this repository contains only libraries,
so build will not compile to binary file.

### License

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B7211%2Fgithub.com%2Fkc-workspace%2Fgo-lib.svg?type=large)](https://app.fossa.com/projects/custom%2B7211%2Fgithub.com%2Fkc-workspace%2Fgo-lib?ref=badge_large)