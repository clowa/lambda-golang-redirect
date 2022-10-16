# lambda-golang-redirect

Lightweight Configurable Lambda Redirection Service

Manage redirects using AWS Lambda functions.

## Requirements

- [Make](https://www.gnu.org/software/make/manual/make.html)
- [NodeJS](https://nodejs.org/en/download/package-manager/) v4 or greater
- [Yarn](https://yarnpkg.com/getting-started/install)
- [GoLang](https://golang.org/doc/install) v1.x
- [Serverless](https://serverless.com/framework/docs/providers/aws/guide/installation/)
- [AWS Account](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/) with Lambda management permissions
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)

## Installation

`lambda-golang-redirect` has a heavy dependency set listed [above](#requirements).

Once dependencies are met installation is a simple command:

    ```bash
    $ make deploy
    ```

This will build the Go binaries and deploy everything to AWS Lambda.

## Usage

Set environment variable `REDIRECT_TO` to desired redirect destination uri like `https://example.org`

## Deployment

// todo

## License

MIT
