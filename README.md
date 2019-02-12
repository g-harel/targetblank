<!--

TODO
- preserve cursor position in editor
- keyboard shortcuts for quick links
- release
    - remove "under construction" banner
    - take application secrets from environment
    - wipe database (possibly contains old formats)
    - set up regular database backups

requirements
- text-based template for links, labels, categories, etc.
- optional search bar with a few search providers
- short url using 6 alphanumeric chars (https://targetblank.org/aB7pPo)
- submit email and receive a link to a new page
- temp password that can be included in the url
- email used to get new temp password
- pages can be made public at their existing url
- credentials stored
- collapsible labels
- open all tabs button

notes
- frontend spa served from s3 + cloudfront
- api gateway + functions backed by dynamodb

endpoints (/api/v1..)
- authenticate per page  (POST   /auth/:address        password)
- change page password   (PUT    /auth/:address [auth] password)
- reset page password    (DELETE /auth/:address        email   )
- create new page        (POST   /page                 email   )
- validate page document (POST   /page/validate        doc     )
- fetch page             (GET    /page/:address [auth]         )
- edit page document     (PUT    /page/:address [auth] data    )

dynamodb schema {
    addr: string (6 alphanumeric chars),
    password: string (hashed),
    email: string (hashed),
    published: bool,
    page: string
}

links
- https://undraw.co/illustrations
- https://gauger.io/fonticon/
- http://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda
- https://github.com/nzoschke/gofaas
- https://read.acloud.guru/how-to-keep-your-lambda-functions-warm-9d7e1aa6e2f0
- https://gist.github.com/prwhite/8168133
- https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-general-nosql-design.html
- https://scene-si.org/2018/05/08/protecting-api-access-with-jwt/
- https://www.terraform.io/docs/providers/aws/guides/serverless-with-aws-lambda-and-api-gateway.html
- https://github.com/hashicorp/best-practices/tree/master/terraform

-->

# [targetblank](https://targetblank.org)

[![Build Status](https://travis-ci.org/g-harel/targetblank.svg?branch=master)](https://travis-ci.org/g-harel/targetblank)
[![Test Coverage](https://img.shields.io/codecov/c/github/g-harel/targetblank.svg)](https://codecov.io/gh/g-harel/targetblank)

Targetblank is an in-browser tool to organize links. Pages are defined by a markdown-like document and can be shared with others.

## Development

```bash
$ npm install
```

```bash
$ npm run dev
```

This will launch a local server which watches and serves contents from [`./website`](./website).

_The api is not mocked in dev mode, it will use production data._

## Deployment

This project is hosted on AWS and uses Terraform to manage the cloud resources.

To recompile Lambda Functions and transpile frontend assets:

```bash
$ make build
```

To update deployed resources after changes to the terraform files or to the source code:

```bash
$ terraform apply
```

_The AWS credentials will be [loaded from the environment](https://www.terraform.io/docs/providers/aws/#environment-variables)._

## License

[MIT](./LICENSE)
