## sifnoded add-genesis-validators

add genesis validators to genesis.json

### Synopsis

add validator to genesis.json. The provided account must specify
the account address or key name. If a key name is given, the address will be looked up in the local Keybase. 


```
sifnoded add-genesis-validators [address_or_key_name] [flags]
```

### Options

```
  -h, --help                     help for add-genesis-validators
      --home string              node's home directory (default "/Users/ivanverchenko/.sifnoded")
      --keyring-backend string   Select keyring's backend (os|file|test) (default "os")
```

### SEE ALSO

* [sifnoded](sifnoded.md)	 - app Daemon (server)

###### Auto generated by spf13/cobra on 2-Jul-2021
