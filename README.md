# Trash-CLI

A conformant freedesktop's Trash implementation for linux terminal.
It supports deleting file with Trash-like behaviour in commandline.
Currently only support deleting local file correctly. That means it
can't put files to trash from networked drive, or removable drive.

> Still on development

## Usage

Trash files:
```
$ trash put foo bar
```

List trashed files:
```
$ trash list
```

Restore a trashed file:
```
$ trash restore
```

Remove all file permanently from trashcan:
```
$ trash empty
```

## Installation

WIP
