[![Build Status](https://travis-ci.com/pulumi/pulumi-godaddy.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-godaddy)

# godaddy Resource Provider

The godaddy resource provider for Pulumi lets you manage godaddy resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/godaddy

or `yarn`:

    $ yarn add @pulumi/godaddy

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_godaddy

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-godaddy/sdk/go/...

## Configuration

The following configuration points are available:

- `godaddy:key` - (Required) Key used to authenticate to the GoDaddy Customer API. May be set via the `GODADDY_API_KEY` environment variable.
- `godaddy:secret` - (Optional) API secret. May be set via the `GODADDY_API_SECRET` environment variable.

## Reference

For detailed reference documentation, please visit [the API docs](https://pulumi.io/reference/pkg/nodejs/@pulumi/godaddy/index.html).
