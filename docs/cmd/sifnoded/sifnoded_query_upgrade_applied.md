## sifnoded query upgrade applied

block header for height at which a completed upgrade was applied

### Synopsis

If upgrade-name was previously executed on the chain, this returns the header for the block at which it was applied.
This helps a client determine which binary was valid over a given range of blocks, as well as more context to understand past migrations.

```
sifnoded query upgrade applied [upgrade-name] [flags]
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for applied
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID
```

### SEE ALSO

* [sifnoded query upgrade](sifnoded_query_upgrade.md)	 - Querying commands for the upgrade module

###### Auto generated by spf13/cobra on 2-Jul-2021
