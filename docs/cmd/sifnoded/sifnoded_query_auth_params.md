## sifnoded query auth params

Query the current auth parameters

### Synopsis

Query the current auth parameters:

$ <appd> query auth params

```
sifnoded query auth params [flags]
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for params
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID
```

### SEE ALSO

* [sifnoded query auth](sifnoded_query_auth.md)	 - Querying commands for the auth module

###### Auto generated by spf13/cobra on 2-Jul-2021
