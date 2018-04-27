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

    endpoints
    - create new homepage (email ->)
    - fetch homepage (address + jwt? -> data)
    - authenticate per homepage (address + password -> jwt)
    - change homepage password (address + jwt + password ->)
    - reset homepage password (address + jwt + email ->)
    - edit homepage template (address + jwt + data ->)
    - delete homepage (address + jwt ->)
    - make homepage public (address + jwt ->)

    schema {
        address: string (6 alphanumeric chars),
        password: string (hashed),
        email: string (hashed),
        temporary: string (hashed) || null,
        public: bool,
        rawSpec: string,
        rawSpecVersion: string,
        spec: ...
    }

 -->

<!-- https://github.com/nzoschke/gofaas -->
<!-- https://read.acloud.guru/how-to-keep-your-lambda-functions-warm-9d7e1aa6e2f0 -->
<!-- https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-general-nosql-design.html -->
<!-- https://hackernoon.com/introduction-to-aws-with-terraform-7a8daf261dc0?__s=vbsgitpwqxmgtteyr9km -->
<!-- https://www.terraform.io/docs/providers/aws/guides/serverless-with-aws-lambda-and-api-gateway.html -->
<!-- https://serverless.com/framework/docs/providers/aws/examples/hello-world/go/ -->
