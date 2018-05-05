# targetblank

<!--

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

notes
- frontend spa served from s3 + cloudfront
- api gatweay + functions backed by dynamodb

endpoints (/api/v1..)
- create new homepage       (POST   /page                 email   )
- validate homepage spec    (GET    /page                 spec    )
- fetch homepage            (GET    /page/:address [auth]         )
- authenticate per homepage (POST   /auth/:address        password)
- change homepage password  (PUT    /auth/:address [auth] password)
- reset homepage password   (DELETE /auth/:address [auth] email   )
- edit homepage template    (PUT    /page/:address [auth] data    )
- delete homepage           (DELETE /page/:address [auth]         )
- make homepage public      (PATCH  /page/:address [auth]         )

nosql schema {
    addr: string (6 alphanumeric chars),
    password: string (hashed),
    email: string (hashed),
    temporary: string (hashed) || null,
    public: bool,
    rawSpec: string,
    rawSpecVersion: string,
    spec: ...
}

links
- http://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda
- https://github.com/nzoschke/gofaas
- https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-general-nosql-design.html
- https://read.acloud.guru/how-to-keep-your-lambda-functions-warm-9d7e1aa6e2f0

-->
