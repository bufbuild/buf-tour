Contributing
============

## Syncing tutorial modules to the BSR

When adding a module for a tutorial that should be available on the BSR, you need to set that module up with the [buf-push-action](https://github.com/bufbuild/buf-push-action). 

Ensure that your module is included in the `start` directory [workspace](../../start/buf.work.yaml). You can ensure that you have correctly set up the module in the workspace by running `buf build` from the `start` directory.

Add the new module to the [buf workflow](../workflows/buf.yaml), similarly to how the tutorial-breaking module is configured there. For example, if you were adding a `tutorial-foo` module with its protobuf files defined in a `proto` subdirectory, you'd add the following clause to the workflow:
```
      # Push the tutorial-foo module
      - uses: bufbuild/buf-push-action@v1
        input: tutorial-foo/proto
        with:
          buf_token: ${{ secrets.BUF_TOKEN }}
```
