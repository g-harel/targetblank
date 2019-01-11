# targetblank

[![Build Status](https://travis-ci.org/g-harel/targetblank.svg?branch=master)](https://travis-ci.org/g-harel/targetblank)
[![Test Coverage](https://img.shields.io/codecov/c/github/g-harel/targetblank.svg)](https://codecov.io/gh/g-harel/targetblank)

<!--

https://www.terraform.io/docs/providers/aws/guides/serverless-with-aws-lambda-and-api-gateway.html
https://github.com/hashicorp/best-practices/tree/master/terraform

TODO
- refactor
    - better error handling in all functions
- make terraform module to manage functions
    - rename functions?
    - add prefix/tag?
- random blob under logo on landing
- display uncaught network errors
- app ~ treat 5xx errors differently
- replace token secrets (eventually)
- offline page for playing around ("/localhost")

requirements
- text template specs for links, labels, categories, etc.
- optional search bar with a few search providers
- short url using 6 alphanumeric chars (https://targetblank.org/aB7pPo)
- submit email and receive a link to a new homepage
- temp password that can be included in the url
- email used to get new temp password
- homepages can be made public at their existing url
- credentials stored
- collapsible labels
- open all tabs button

notes
- frontend spa served from s3 + cloudfront
- api gateway + functions backed by dynamodb

endpoints (/api/v1..)
- authenticate per homepage (POST   /auth/:address        password)
- change homepage password  (PUT    /auth/:address [auth] password)
- reset homepage password   (DELETE /auth/:address        email   )
- create new homepage       (POST   /page                 email   )
- validate homepage spec    (POST   /page/validate        spec    )
- fetch homepage            (GET    /page/:address [auth]         )
- edit homepage template    (PUT    /page/:address [auth] data    )
- delete homepage           (DELETE /page/:address [auth]         )
- make homepage public      (PATCH  /page/:address [auth]         )

nosql schema {
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

-->
