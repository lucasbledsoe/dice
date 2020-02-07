This test Lambda code is designed to be deployed using [localstack][localstack].

Set up IAM policies for the Lambda:

```sh
make iam-create
make iam-attach
```

Compile, zip, and upload the Lambda code:

```sh
make clean
make build
make zip
make deploy
```

Invoke the Lambda:

```sh
make invoke
```

[localstack]: https://github.com/localstack/localstack
