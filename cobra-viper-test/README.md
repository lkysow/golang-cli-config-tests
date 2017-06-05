Root level help
```bash
./cobra-viper-test
Long description of cobra/viper test

Usage:
  cobra-viper-test [command]

Available Commands:
  help        Help about any command
  server      short server descrip

Flags:
  -h, --help   help for cobra-viper-test

Use "cobra-viper-test [command] --help" for more information about a command.
```

Subcommand help
```bash
./cobra-viper-test server -h
longer server descrip

Usage:
  cobra-viper-test server [flags]

Flags:
      --config string        config file
      --gh-username string   gh username
  -h, --help                 help for server
```

Config file
```bash
./cobra-viper-test server --config config.yaml
config:map[config:config.yaml gh-username:from yaml]
```

Flag override
```bash
cobra-viper-test ./cobra-viper-test server --config config.yaml --gh-username "from flag"
config:map[config:config.yaml gh-username:from flag]
```
