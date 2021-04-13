# Cofy

<!-- [![RELEASE](https://img.shields.io/github/v/release/blyndusk/cofy)](https://github.com/blyndusk/cofy/releases)
[![RELEASE DATE](https://img.shields.io/github/release-date/blyndusk/cofy)](https://github.com/blyndusk/cofy/commits/main) -->
[![ISSUES](https://img.shields.io/github/issues/blyndusk/cofy)](https://github.com/blyndusk/cofy/issues)
[![PULL REQUESTS](https://img.shields.io/github/issues-pr/blyndusk/cofy)](https://github.com/blyndusk/cofy/pulls) 
[![LAST COMMIT](https://img.shields.io/github/last-commit/blyndusk/cofy)](https://github.com/blyndusk/cofy/commits/main)
[![LICENSE](https://img.shields.io/github/license/blyndusk/cofy)](https://github.com/blyndusk/cofy/blob/main/LICENSE) 
[![REPO SIZE](https://img.shields.io/github/repo-size/blyndusk/cofy)](https://github.com/blyndusk/cofy) 
<!-- 
[![DOCKER](https://github.com/blyndusk/cofy/actions/workflows/docker.yml/badge.svg)](https://github.com/blyndusk/cofy/actions/workflows/docker.yml) 
[![PYTHON](https://github.com/blyndusk/cofy/actions/workflows/python.yml/badge.svg)](https://github.com/blyndusk/cofy/actions/workflows/python.yml)
[![RELEASE](https://github.com/blyndusk/cofy/actions/workflows/release.yml/badge.svg)](https://github.com/blyndusk/cofy/actions/workflows/release.yml) -->

- [Cofy](#cofy)
  - [I - Goal](#i---goal)
    - [FR](#fr)
    - [EN](#en)
  - [II - Conventions, templates and guidelines](#ii---conventions-templates-and-guidelines)
    - [A - Commit conventions](#a---commit-conventions)
    - [B - Issue template](#b---issue-template)
    - [C - Branch naming convention](#c---branch-naming-convention)
    - [D - Pull request template](#d---pull-request-template)
  - [IV - Project use](#iv---project-use)
    - [Help](#help)
    - [Start](#start)
    - [Stop](#stop)
    - [Restart](#restart)
    - [Display logs](#display-logs)
  - [III - License](#iii---license)

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


## II - Conventions, templates and guidelines

### A - Commit conventions

```
tag(scope): #issue_id message
```

See [commit_conventions.md](.github/commit_conventions.md) for more informations.

### B - Issue template

See [user-story.md](.github/ISSUE_TEMPLATE/user-story.md) for more informations.

### C - Branch naming convention

```
type_scope-of-the-work
```

See [branch_naming_convention.md](.github/branch_naming_convention.md) for more informations.

### D - Pull request template

See [pull_request_template.md](.github/pull_request_template.md) for more informations.


## IV - Project use

### Help

```bash
$ make help
```

### Start

```bash
$ make start
```

### Stop

```bash
$ make start
```

### Restart

```bash
$ make restart
```

### Display logs

```bash
$ make logs
```



## III - License

Under [MIT](./LICENSE) license.
