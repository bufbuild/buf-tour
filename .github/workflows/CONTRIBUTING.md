Contributing
============

## Syncing tutorial modules to the BSR

When adding a module for a tutorial that should be available on the BSR, there are a few prerequisite steps:

### Create the module in the BSR

Ensure that the module you plan on syncing has been created in the BSR. For example, if you were adding a `foo` module in the [tutorials](https://buf.build/tutorials) organization, you would need to make sure that a new module is created there before you start syncing the protobuf files.

### Set the module up with the [buf-push-action](https://github.com/bufbuild/buf-push-action)

Add the new module to the [buf workflow](../workflows/buf.yaml), similarly to how the tutorial-breaking module is configured there. For example, if you were adding a `tutorial-foo` module with its protobuf files defined in a `proto` subdirectory, you'd add the following clause to the workflow:
```
  foo-module:
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
      - uses: actions/checkout@v3
      - uses: bufbuild/buf-setup-action@v1
      - uses: bufbuild/buf-push-action@v1
        with:
          input: start/tutorial-foo/proto
          buf_token: ${{ secrets.BUF_TOKEN }}
```
