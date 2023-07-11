# Debug getting transaction results

The following block of mainnet22 errors out

- 54564338
- 54564339
- 54564340


the error i get is:
```panic: client: rpc error: code = NotFound desc = not found: could not retrieve transaction: could not retrieve resource: key not found```


# this works
 flow blocks get 54564338 --host access-001.mainnet22.nodes.onflow.org:9000 --include transactions

