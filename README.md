<p align="center">
  <img src="docs/logo.png" width="200"/>
</p>
<h1 align="center">Cofy</h1>
<p align="center">
  <p align="center">
    <a href="https://github.com/blyndusk/cofy/releases">
      <img src="https://img.shields.io/github/v/release/blyndusk/cofy"/>
    </a>
    <a href="https://github.com/blyndusk/cofy/commits/main">
      <img src="https://img.shields.io/github/release-date/blyndusk/cofy"/>
    </a>
    <a href="https://github.com/blyndusk/cofy/issues">
      <img src="https://img.shields.io/github/issues/blyndusk/cofy"/>
    </a>
    <a href="https://github.com/blyndusk/cofy/pulls">
      <img src="https://img.shields.io/github/issues-pr/blyndusk/cofy"/>
    </a>
    <a href="https://github.com/blyndusk/cofy/commits/main">
      <img src="https://img.shields.io/github/last-commit/blyndusk/cofy"/>
    </a>
    <a href="https://github.com/blyndusk/cofy/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/blyndusk/cofy"/>
    </a>
    <a href="https://github.com/blyndusk/cofy">
      <img src="https://img.shields.io/github/repo-size/blyndusk/cofy"/>
    </a>
  </p>
</p>

<p align="center">
  <a href="https://github.com/blyndusk/cofy/actions/workflows/go.yml">
      <img src="https://github.com/blyndusk/cofy/actions/workflows/go.yml/badge.svg"/>
    </a>
     <a href="https://github.com/blyndusk/cofy/actions/workflows/docker.yml">
      <img src="https://github.com/blyndusk/cofy/actions/workflows/docker.yml/badge.svg"/>
    </a>
     <a href="https://github.com/blyndusk/cofy/actions/workflows/release.yml">
      <img src="https://github.com/blyndusk/cofy/actions/workflows/release.yml/badge.svg"/>
    </a>
</p>

- [I - Goal](#i---goal)
  - [FR](#fr)
  - [EN](#en)
- [II - Table of content](#ii---table-of-content)
- [III - Conventions, templates and labels](#iii---conventions-templates-and-labels)
  - [A - Commit conventions](#a---commit-conventions)
  - [B - Branch naming convention](#b---branch-naming-convention)
  - [C - Issue template](#c---issue-template)
  - [D - Pull request template](#d---pull-request-template)
  - [E - Custom issues labels preset](#e---custom-issues-labels-preset)
- [IV - CI/CD, release and container registry](#iv---cicd-release-and-container-registry)
  - [A - CI](#a---ci)
  - [B - CD](#b---cd)
  - [C - Release](#c---release)
- [V - Golang RESTful API](#v---golang-restful-api)
  - [A - Stack](#a---stack)
  - [B - Makefile](#b---makefile)
- [VI - License](#vi---license)

## I - Goal

### FR

"Cofy" est un robot discord avec comme but principal d'accumuler un maximum de café.

Pour ce faire, à chaque fois que l'utilisateur écrira un message sur un serveur ayant cofy, il gagnera différentes ressources, à savoir:

- de l'argent
- de l'expérience

L'expérience servira à monter en niveau, chaque pallier de niveau débloquant de nouveaux types de café (e.g. lvl5: longo, lvl10: macchiatto)

L'argent servira à acheter du café (e.g. 5c: expresso, 50c: macchiatto)

Il y aura différent jeux permettant de gagner de l'argent ou des cafés basé sur la chance (e.g. coinflip, dice roll).

L'utilisateur pourra gagner des badges pour différent palliers:

- badge pour un certain niveau gagné
- badge pour certain argent dépensé
- badge pour certain cafés obtenus

Chaque utilisateur aura un inventaire, contenant:

- expérience
- argent
- différent types de café
- badge et achievements

Les utilisateurs pourront voir un tableau des scores, par type de café, par expérience.

### EN

"Cofy" is a discord robot with the main goal of accumulating a maximum of coffee.

To do this, each time the user writes a message on a server with cofy, he will earn different resources, namely

- money
- experience

Experience will be used to level up, each level unlocking new types of coffee (e.g. lvl5: longo, lvl10: macchiatto)

Money will be used to buy coffee (e.g. 5c: espresso, 50c: macchiatto)

There will be different games to win money or coffee based on luck (e.g. coinflip, dice roll).

The user can earn badges for different levels:

- badge for a certain level earned
- badge for certain money spent
- badge for certain coffee obtained

Each user will have an inventory, containing:

- experience
- money
- different types of coffee
- badge and achievements

Users will be able to see a table of scores, by type of coffee, by experience.
<h1 align="center">Repo Template</h1>
<p align="center">
    <a href="https://github.com/blyndusk/repo-template/releases">
      <img src="https://img.shields.io/github/v/release/blyndusk/repo-template"/>
    </a>
    <a href="https://github.com/blyndusk/repo-template/commits/main">
      <img src="https://img.shields.io/github/release-date/blyndusk/repo-template"/>
    </a>
    <a href="https://github.com/blyndusk/repo-template/issues">
      <img src="https://img.shields.io/github/issues/blyndusk/repo-template"/>
    </a>
    <a href="https://github.com/blyndusk/repo-template/pulls">
      <img src="https://img.shields.io/github/issues-pr/blyndusk/repo-template"/>
    </a>
    <a href="https://github.com/blyndusk/repo-template/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/blyndusk/repo-template"/>
    </a>
</p>

<p align="center">
  <a href="https://github.com/blyndusk/repo-template/actions/workflows/go.yml">
      <img src="https://github.com/blyndusk/repo-template/actions/workflows/go.yml/badge.svg"/>
    </a>
     <a href="https://github.com/blyndusk/repo-template/actions/workflows/docker.yml">
      <img src="https://github.com/blyndusk/repo-template/actions/workflows/docker.yml/badge.svg"/>
    </a>
     <a href="https://github.com/blyndusk/repo-template/actions/workflows/release.yml">
      <img src="https://github.com/blyndusk/repo-template/actions/workflows/release.yml/badge.svg"/>
    </a>
</p>

<p align="center">
  <a href="https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-on-github/creating-a-repository-from-a-template">How to use this repo as a template for your project</a>
</p>

## II - Table of content

- [I - Goal](#i---goal)
  - [FR](#fr)
  - [EN](#en)
- [II - Table of content](#ii---table-of-content)
- [III - Conventions, templates and labels](#iii---conventions-templates-and-labels)
  - [A - Commit conventions](#a---commit-conventions)
  - [B - Branch naming convention](#b---branch-naming-convention)
  - [C - Issue template](#c---issue-template)
  - [D - Pull request template](#d---pull-request-template)
  - [E - Custom issues labels preset](#e---custom-issues-labels-preset)
- [IV - CI/CD, release and container registry](#iv---cicd-release-and-container-registry)
  - [A - CI](#a---ci)
  - [B - CD](#b---cd)
  - [C - Release](#c---release)
- [V - Golang RESTful API](#v---golang-restful-api)
  - [A - Stack](#a---stack)
  - [B - Makefile](#b---makefile)
- [VI - License](#vi---license)

## III - Conventions, templates and labels

### A - Commit conventions

```
tag(scope): #issue_id - message
```

See [commit_conventions.md](.github/commit_conventions.md) for more informations.

### B - Branch naming convention

```
type_scope-of-the-work
```

See [branch_naming_convention.md](.github/branch_naming_convention.md) for more informations.

### C - Issue template

See [user-story.md](.github/ISSUE_TEMPLATE/user-story.md) for more informations.

### D - Pull request template

See [pull_request_template.md](.github/pull_request_template.md) for more informations.

### E - Custom issues labels preset

The labels preset is located at [.github/settings.yml](.github/settings.yml).

You can **add, edit or remove** them. To automatically update these labels, you need to **install** the ["Settings" GitHub app](https://github.com/apps/settings), which will **syncs repository settings defined in the file above to your repository**.

## IV - CI/CD, release and container registry

### A - CI

[![GO](https://github.com/blyndusk/repo-template/actions/workflows/go.yml/badge.svg)](https://github.com/blyndusk/repo-template/actions/workflows/go.yml)

The **CI** workflow is located at [.github/workflows/go.yml](.github/workflows/go.yml). It's triggered a **each push** on **all branches**.

It consist of:

- **install Golang** on the VM
- get the dependancies **using the cache of other Actions run**
- **lint the files** to check or syntax errors
- **build** the application

### B - CD

[![DOCKER](https://github.com/blyndusk/repo-template/actions/workflows/docker.yml/badge.svg)](https://github.com/blyndusk/repo-template/actions/workflows/docker.yml)

The **CD** workflow is located at [.github/workflows/docker.yml](.github/workflows/docker.yml). It's triggered a **each push** on **`main` branch**.

It consist of:

- **login** into the GitHub container registry (ghcr.io)
- **build and push** the Golang api using the **production Dockerfile** located at [.docker/api/prod.Dockerfile](.docker/api/prod.Dockerfile)

After that, you can check the **pushed container** at: `https://github.com/<username>?tab=packages&repo_name=<repository-name>`

> IMPORTANT: you need to **update the production Dockerfile** with your **username** AND **_repository name_**. Otherwise, there will be errors at push:

```bash
LABEL org.opencontainers.image.source = "https://github.com/<username>/<repository-name>"
```

### C - Release

[![RELEASE](https://github.com/blyndusk/repo-template/actions/workflows/release.yml/badge.svg)](https://github.com/blyndusk/repo-template/actions/workflows/release.yml)

The **release** workflow is located at [.github/workflows/release.yml](.github/workflows/release.yml). It's triggered **manually by user input** at: [Actions > RELEASE](https://github.com/blyndusk/repo-template/actions/workflows/release.yml).

> IMPORTANT: you need to set the **image tag** in the action input, for the action to be able to push the docker image and create a release **with a specific version**. The image tag is a [SemVer](https://en.wikipedia.org/wiki/Software_versioning) tag, e.g. `1.0.2`.

It consist of:

- check if the **environment match the branch**
- do the CD (docker) action again, but **with a specific image tag**
- create a release **with the same tag**, filled with the **generated changelog as closed issues since the last release**

After that, you can check the release at `https://github.com/<username>/<repository-name>/releases`.

## V - Golang RESTful API

The project use **Docker** and **Docker Compose** to build and run local and distant images in our workspace.

### A - Stack

All the images use the **same network**, more informations at [docker-compose.yml](docker-compose.yml)

| CONTAINER | PORT        | IMAGE                                                    |
| :-------- | :---------- | :------------------------------------------------------- |
| GOLANG    | `3333:3333` | [.docker/api/dev.Dockerfile](.docker/api/dev.Dockerfile) |
| ADMINER   | `3334:8080` | [.docker/adminer/Dockerfile](.docker/adminer/Dockerfile) |
| POSTGRES  | `5432:5432` | [postgres:latest](https://hub.docker.com/_/postgres)     |

> Adminer is a GUI that allows us to **manage your database** by permetting to to **create, edit, delete** the different entities, tables, etc.

### B - Makefile

#### TL;DR <!-- omit in toc -->

```bash
make setup-env start logs
```

#### `make help` <!-- omit in toc -->

**Display** informations about other commands.

#### `make setup-env` <!-- omit in toc -->

**Copy** the sample environment files.

#### `make start` <!-- omit in toc -->

Up the containers with **full cache reset** to avoid cache errors.

#### `make stop` <!-- omit in toc -->

**Down** the containers.

#### `make logs` <!-- omit in toc -->

**Display and follow** the logs.

#### `make lint` <!-- omit in toc -->

**Lint** the Go files using `gofmt`.

## VI - License

Under [MIT](./LICENSE) license. 
