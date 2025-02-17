## sifnoded query clp lp

Get Liquidity Provider

### Synopsis

Query details for a liquidity provioder.
Example:
$ <appd> pool ETH sif1h2zjknvr3xlpk22q4dnv396ahftzqhyeth7egd

```
sifnoded query clp lp [External Asset symbol] [lpAddress] [flags]
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for lp
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID
```

### SEE ALSO

* [sifnoded query clp](sifnoded_query_clp.md)	 - Querying commands for the clp module

###### Auto generated by spf13/cobra on 2-Jul-2021
