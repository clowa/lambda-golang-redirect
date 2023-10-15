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

1. Run `yarn install` to download dependencies.
2. Change provider configuration of `serverless.yaml` as needed.

   ```yaml
   provider:
     name: aws
     runtime: go1.x
     region: eu-central-1
     architecture: arm64
     memorySize: 128
     stage: prod
     # Function environment variables
     environment:
       REDIRECT_TO: https://example.org # Change me
       HSTS_ENABLED: false
     # Duration for CloudWatch log retention (default: forever)
     logRetentionInDays: 7
     stackTags:
       app: ${self:service}
       stage: ${self:provider.stage}
       deploymentMethod: serverless
       repository: clowa/lambda-golang-redirect

   custom:
     domains:
       # References to 'prod' stage
       prod:
         domainName: clowa.de # Change me
         certificateName: clowa.de # Change me

     customCertificate:
       # Route 53 Hosted Zone name
       # don't forget the dot on the end!
       hostedZoneNames: "clowa.de." # Change me
   ```

3. Run `make deploy`
4.

### Environment variables

| Variable                  | Default                         | Description                                                                                                                                                                                                      |
| ------------------------- | ------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `REDIRECT_TO`             | `https://example.org`           | URI of the redirect target                                                                                                                                                                                       |
| `HSTS_ENABLED`            | `true`                          | Wether or not [Strict-Transport-Security](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) should be enabled                                                                 |
| `HSTS_MAX_AGE`            | `300 * 24 * 60 * 60` _300 days_ | The time, in seconds, that the browser should remember that a site is only to be accessed using HTTPS.                                                                                                           |
| `HSTS_INCLUDE_SUBDOMAINS` | `false`                         | If this parameter is `true`, this rule applies to all of the site's subdomains as well.                                                                                                                          |
| `HSTS_PRELOAD`            | `false`                         | See Preloading [Strict Transport Security](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security#preloading_strict_transport_security) for details. Not part of the specification. |

## Deployment

// todo

## License

MIT
