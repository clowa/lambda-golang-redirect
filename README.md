# go-redirect

Lightweight Configurable Lambda Redirection Service

Manage redirects using AWS Lambda functions.

## Requirements

- [Make](https://www.gnu.org/software/make/manual/make.html)
- [NodeJS](https://nodejs.org/en/download/package-manager/) v4 or greater
- [GoLang](https://golang.org/doc/install) v1.x
- [Docker](https://docs.docker.com/install/)
- [Serverless](https://serverless.com/framework/docs/providers/aws/guide/installation/)
- [AWS Account](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/) with Lambda management permissions 
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Installation

Go-Redirect has a heavy dependency set listed above (requirements).

Once dependencies are met installation is a simple command:

    $ make

This will build the Go binaries and export template.yml required for AWS SAM local testing.

## Usage

// todo

## Deployment

// todo

## Backlog

- Initial version
- Consolidate configuration
- Pattern matching redirects
- Analytics integration

## License

MIT