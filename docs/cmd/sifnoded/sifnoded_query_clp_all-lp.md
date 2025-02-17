## sifnoded query clp all-lp

Get all liquidity providers on sifnode 

```
sifnoded query clp all-lp [flags]
```

### Options

```
      --count-total       count total number of records in liquidityProviders to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for all-lp
      --limit uint        pagination limit of liquidityProviders to query for (default 100)
      --node string       <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of liquidityProviders to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of liquidityProviders to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of liquidityProviders to query for
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID
```

### SEE ALSO

* [sifnoded query clp](sifnoded_query_clp.md)	 - Querying commands for the clp module

###### Auto generated by spf13/cobra on 2-Jul-2021
