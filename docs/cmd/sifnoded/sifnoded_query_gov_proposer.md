## sifnoded query gov proposer

Query the proposer of a governance proposal

### Synopsis

Query which address proposed a proposal with a given ID.

Example:
$ <appd> query gov proposer 1

```
sifnoded query gov proposer [proposal-id] [flags]
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for proposer
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID
```

### SEE ALSO

* [sifnoded query gov](sifnoded_query_gov.md)	 - Querying commands for the governance module

###### Auto generated by spf13/cobra on 2-Jul-2021
