<p align="center">
    <a href="https://targetblank.org">
        <img src="https://svgsaur.us/?t=targetblank&o=b&s=26&c=332832&w=152&y=55" />
        <br>
        <img src="https://svgsaur.us/?t=organize_your_links&s=16&c=766873&w=152&y=12" />
    </a>
</p>

<!-- TODO example pages -->

- **Productivity focused** &nbsp; [Do everything from the comfort of your keyboard.](#keyboard-shortcuts)

- **Shareable** &nbsp; Pages are publicly readable by anyone with the link.

- **Optimized for performance** &nbsp; Aggressive caching, small code bundles and native font stacks.

<!-- TODO extension blurb

- **Browser Extension** An even snappier homepage and built-in conveniences.

 -->

## Document Format

The page structure is defined entirely by the _document_. Instead of adding, updating and removing links with buttons, targetblank uses text to represent your page. Although the syntax is completely different, it works similarly to other structured documents like [markdown](https://en.wikipedia.org/wiki/Markdown) or [yaml](https://en.wikipedia.org/wiki/YAML).

A minimal _document_ contains at least a _version_ line and a _header_ line. The _version_ is the first thing the parser reads, and it makes it easier to keep the format backwards-compatible if the syntax needs to change in the future.

```diff
+ version 1
+ ===
```

The space between the _version_ and _header_ lines is used for _metadata_ about the page. For now, this only allows the page's title to be configured, but please [open an issue](https://github.com/g-harel/targetblank/issues/new) to share your ideas (ex. background image). The format of a _metadata_ entry is a _name_, followed by `=` and finally the _value_.

```diff
  version 1
+ title = Hello World!
  ===
```

Your _links_ are added after the _header_ line and can be formatted in two ways. The simplest way is to simply write the _link_ on its own line, the parser will make sure it's detected and clickable. Alternatively, _links_ can be _labelled_, which lets you control what text is displayed in place of the full _link_.

```diff
  version 1
  title = Hello World
  ===
+ example.com?q=foo
+ labelled link [example.com?q=bar]
```

You can also add _labels_ on their own, without a link. These will not be clickable, but can be a convenient way to organize your _links_. For even more control, the _links_ and _labels_ can be indented into whatever shape makes sense for your use case.

```diff
  version 1
  title = Hello World
  ===
+ productivity
      example.com?q=foo
+         example.com/baz
      labelled link [example.com?q=bar]
```

When a single _group_ isn't enough, you can add a _group delimiter_ to create a new one and add more _links_. This will be reflected on the homepage by a new section side-by-side with the first one. After a few _groups_ (depending on the size of your screen) they will start wrapping underneath the first row of sections on the homepage.

```diff
  version 1
  title = Hello World
  ===
  productivity
      example.com?q=foo
          example.com/baz
      labelled link [example.com?q=bar]
  ---
+ communication
+     example.com/chat
```

## Keyboard Shortcuts

Targetblank is meant to be usable and productive using only a keyboard. The [document format](#document-format) goes a long way towards making that a reality, but these shortcuts help complete the story, adding quick navigation and useful text-editing commands. The editor shortcuts are inspired by common text editor keybindings and work on multi-line selections. If your favorite shortcut is missing, please [let me know](https://github.com/g-harel/targetblank/issues/new).

page     | shortcut       | description
-------- | -------------- | -----------------------
homepage | `ctrl + enter` | Navigate to editor
&nbsp;   | `<any>`        | Highlight link
&nbsp;   | `enter`        | Follow highlighted link
editor   | `ctrl + enter` | Navigate to homepage
&nbsp;   | `tab`          | Indent
&nbsp;   | `shift + tab`  | Un-indent
&nbsp;   | `ctrl + ]`     | Indent
&nbsp;   | `ctrl + [`     | Un-indent
&nbsp;   | `ctrl + /`     | Toggle comment
&nbsp;   | `alt + up`     | Move up
&nbsp;   | `alt + down`   | Move down

## Development

[![Build Status](https://travis-ci.org/g-harel/targetblank.svg?branch=master)](https://travis-ci.org/g-harel/targetblank)
[![Test Coverage](https://img.shields.io/codecov/c/github/g-harel/targetblank.svg)](https://codecov.io/gh/g-harel/targetblank)

```bash
$ npm run dev
```

This will launch a local server which watches and serves contents from [`./website`](./website).

_The api is not mocked in dev mode, it points to production endpoints._

The built-in `go test` command can be used to validate changes to backend code.

<!-- TODO extension dev -->

## Deployment

This project is hosted on AWS and uses [Terraform](https://www.terraform.io/) to manage the cloud resources.

The [deployment workflow](./.github/main.workflow) uses [GitHub Actions](https://developer.github.com/actions/) to package and apply changes on every change to master.

<!-- TODO browser deploy -->

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
