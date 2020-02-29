🧳 gobundle
===========

> quickly update or install your go tools

Gobundle makes your Go tools portable and allows you to keep them up to date.

# Installing / Getting started

```shell
go get github.com/iwittkau/gobundle/cmd/gobundle@latest
gobundle -v
```

## Usage

```
$ gobundle -h
Usage of gobundle:
  -d    dumps installed tools and versions to gobundle.json
  -f string
        set configuation file
  -i    init configuration in home directory
  -v    print version
``` 

## Initial Configuration

```
gobundle -i
``` 

# Features

If you run `gobundle` without any flags, it tries to install the packages defined in your `.gobundle` configuration.

By using version `latest` for a package you can ensure you are always installing the latest available version of a tool. By specifying a specific version, you can pin a tool to this version. If you use `branch` as the version, gobundle installs the latest available version of that branch (e.g. `master`).

## Keeping your Go tools up to date

By choosing `latest` as the version and running `gobundle` regularly, you can keep your go tools up to date easily.

## Dumping your installed tools and their versions

- `gobundle -d` dumps a `gobundle.json` of all tools and their installed versions from the `$GOPATH/bin` directory. 
- Tools that were built from local source will be skipped. 
- To install tools from a dumped `gobundle.json` use `gobundle -f gobundle.json`.

# Configuration

If `gobundle -i` was previously run, a configuration file will exist in `$HOME/.gobundle`.

It is a very simple data structure, check [examples/gobundle.json](examples/gobundle.json).

You can start by dumping your curently installed tools and copy them to `$HOME/.gobundle`

```
$ gobundle -d
$ cp ./gobundle.json ~/.gobundle 
``` 

## Contributing

If you'd like to contribute, please fork the repository and use a feature
branch. Pull requests are warmly welcome. 

If you encounter any issue, please let me know [here](https://github.com/iwittkau/gobundle/issues).

## Licensing

The code in this project is licensed under MIT license.