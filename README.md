git-not-young
======

## Description

You are forced to use git checkout instead of git switch and git restore by `git-not-young`.

## Getting Started

### Install
```shell
go install github.com/ytkg/git-not-young@latest
```

### Settings

Add alias to your `.bash_profile`

- .bash_profile

```
alias git=git-not-young
```

```
$ source ~/.bash_profile
```


## Example

```shell
$ git switch main
Error: Use git switch or git restore instead of git switch.
```
