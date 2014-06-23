bash-init(1)
====

The easy way to start building BASH command-line application.


## Synopsis

`bash-init` is the easy way to start building BASH command-line application.
All you need to do is to set application name and its subcommands. You can forcus on core function of application.

## Defensive BASH Programming

`bash-init` has some inspirations from
[@kfiravi](https://github.com/kfirlavi)'s ["Defensive BASH programming"](http://www.kfirlavi.com/blog/2012/11/14/defensive-bash-programming/) post.

## Usage

You just need to set its application name:

```bash
$ bash-init [options] [application]
```

## Example

If you want to start to building `todo` application which has subcommands `add`, `list`, `delete`:

```bash
$ bash-init -s add,list,delete todo
```

You can see sample, [tcnksm/sample-bash-init](https://github.com/tcnksm/sample-bash-init).

## Artifact

`bash-init` generates **main.sh**, everything in one shellscript.

## Useful functions

`bash-init` generates some useful functions to write good bash script.

### Color messages

It makes outputs more easy to understand.

```bash

info() {
    echo -e "\033[34m$@\033[m" # blue
}

warn() {
    echo -e "\033[33m$@\033[m" # yellow
}

error() {
    echo -e "\033[31m$@\033[m" # red
}
```

### Code clarity

It makes bashscript more readable.

```bash
is_empty() {
    local var=$1

    [[ -z $var ]]
}

is_not_empty() {
    local var=$1

    [[ -n $var ]]
}

is_file() {
    local file=$1

    [[ -f $file ]]
}

is_dir() {
    local dir=$1

    [[ -d $dir ]]
}
```


## Author

[tcnksm](https://github.com/tcnksm)
