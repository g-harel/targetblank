<!--

TODO
- take application secrets from environment
- scroll area in editor

-->

# [targetblank](https://targetblank.org)

[![Build Status](https://travis-ci.org/g-harel/targetblank.svg?branch=master)](https://travis-ci.org/g-harel/targetblank)
[![Test Coverage](https://img.shields.io/codecov/c/github/g-harel/targetblank.svg)](https://codecov.io/gh/g-harel/targetblank)

Browser tool to organize links. Pages are defined by a [structured document](#document-format) which allows links to be labelled, nested and grouped.

- **Productivity focused** Follow links by typing their label and pressing enter.

- **Shareable** Pages are publicly readable by anyone with the link.

- **Optimized for performance** Aggressive caching, small code bundles and native font stacks.

## Document Format

#### Minimal

```shell
version 1
===
```

#### Simple

```shell
version 1
title=bookmarks
===
example.com
email [mail.google.com]
```

#### Detailed

```shell
# Everything after a pound character (#), trailing whitespace and empty lines are ignored.

# Documents must start with their version (currently only 1).
version 1

# Document metadata key-value pairs can be added at the top of the document.
key=value

# The "title" key can be used to name the document.
title=Hello World

# The first group starts after the header line.
===

# Group metadata key-value pairs can be added at the start of each group.
# These values are currently ignored, but may be used in the future.
key=value

# Groups hold entries containing a label and a link.
labelled link [example.com]

# Both the label and the link are optional.
label without link
[google.com]
amazon.com

# New groups are started using the group delimiter.
---

# Group entries can be nested using indentation.
entry 1
    entry 2
        entry 3
    entry 4
```

## Development

```bash
$ npm run dev
```

This will launch a local server which watches and serves contents from [`./website`](./website).

_The api is not mocked in dev mode, it points to production endpoints._

The built-in `go test` command can be used to validate changes to backend code.

## Deployment

This project is hosted on AWS and uses [Terraform](https://www.terraform.io/) to manage the cloud resources.

The [deployment workflow](./.github/main.workflow) uses [GitHub Actions](https://developer.github.com/actions/) to package and apply changes on every change to master.

## License

[MIT](./LICENSE)

<!--

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
    document: string
    email: string (hashed),
    password: string (hashed),
    published: bool,
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
