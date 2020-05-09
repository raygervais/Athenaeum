Contributing to Athenaeum

## Welcome!

This project is an example of building a software platform using modern practices!

## Environment and Setup

The following technologies are used throughout Athenaeum, and should be installed on the development environment.

| Name                                        | Stack          | Use                                                    |
| ------------------------------------------- | -------------- | ------------------------------------------------------ |
| [NodeJS LTS](https://nodejs.org/en/)        | Front-end      | Used to power ReactJS                                  |
| [Golang 1.14](https://golang.org/)          | Back-end       | Used as the REST API Language                          |
| [Docker](https://www.docker.com/)           | Infrastructure | Used for Containerizing the Services                   |
| [SQLite](https://www.sqlite.org/index.html) | Storage        | Used for Data Persistence and Configuration Management |

Once installed, follow the instructions found in the corresponding service directories `ex:` [src/back-end/README.md](https://github.com/raygervais/Athenaeum/src/back-end/README.md) to install the required frameworks and tooling.

## Commits

Leverage our git config by running `git config commit.template ~/.github/commit.template`, which will provide a template commit structure.
Our commit structure is as follows: `prefix: 50-character subject line`.

| Prefix    | Description                                                                                            |
| --------- | ------------------------------------------------------------------------------------------------------ |
| `adds`    | For when new features or code is being added to the codebase                                           |
| `fixes`   | For when we are providing a few line changes which fix common errors such as typos, expressions, etc.  |
| `removes` | for when we are removing a feature, or chunk of code which directly impacts the component or platform. |
| `reverts` | For the rare time when a commit has to be reverted.                                                    |

## Merge Pattern

In attempts to keep the git history clean, Athenaeum pull requests are always merged using the squash pattern, so that multiple commits are bundled into a single "feature" commit when put into `master`.
The merge commit message should follow the same format as the git commits.
