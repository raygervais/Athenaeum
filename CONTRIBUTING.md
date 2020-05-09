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

## Commits and Merge Requests

## Issues

Before creating an issue:

- Please look through the MVP Features listed below for inspiration-
  - Check [open issues](https://github.com/Seneca-CDOT/telescope/issues). Someone else may be working on the same thing!
  - Use [our Labels](https://github.com/Seneca-CDOT/telescope/labels) to help others quickly understand what an issue is all about.
-
- ## Environment Setup
-
- Telescope has many parts, and setup requires you to install a number of tools
- and dependencies. For instructions on how to setup your Telescope environment, please see
- the [Environment Setup documentation](https://github.com/Seneca-CDOT/telescope/blob/master/docs/environment-setup.md).
-
- ## Technologies
-
- Telescope uses quite a few different technologies, and we have some project specific
- docs available for each, including our:
-
- - database [Redis](redis.md)
- - frontend framework [GatsbyJS](gatsbyjs.md)
- - frontend query language [GraphQL](graphql.md)
- - single-sign-on (SSO) [login](login.md) using SAML2
- - backend logging framework [Pino](logging.md)
-
- If you're unsure about how something works, talk to one of us on [#telescope Slack channel](https://seneca-open-source.slack.com/archives/CS5DGCAE5).
-
- ## Workflow in Git and GitHub
-
- We use a number of tools and automated processes to help make it easier for
- everyone to collaborate on Telescope. This includes things like auto-formatting
- code, linting, and automated testing. We also use git and GitHub in particular
- ways.
-
- For more information on working with our tools and our workflows, see our [Git Workflow documentation](git-workflow.md).
-
- ## Reports
-
- We have a number of automated reports and audits that can be run on the code.
- These include things like checking accessibility and performance issues in our
- frontend, and determining test coverage for our automated tests.
-
- For more information on working with these automated reports, see our [Reports documentation](reports.md).
-
- ## Releases
-
- When doing a release of Telescope, a number of steps must be done. To help our
- maintainers do this properly, we have tools and information in our [Release documentation](release.md).
-
