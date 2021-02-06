# quickbase-textmate

A command line tool that automates the creation of a TextMate grammar for Quickbase formulae. The primary goal is to make it easier to maintain IDE extentions like the one at https://github.com/jdklub/vscode-quickbase-formula for VS Code.

## Installation

With a [correctly configured](https://golang.org/doc/install#install) Go toolchain, clone the repository to a directory of your choosing, change into it, and run `make`:

```
git clone https://github.com/cpliakas/quickbase-textmate.git
cd ./quickbase-textmate
make
```

## Usage

The `data/Formula_Functions_Reference.csv` file is the CSV file downloaded from the [Quickbase Formula Functions Reference](https://login.quickbase.com/db/6ewwzuuj?a=td) documentation.

To build the [snippets](https://code.visualstudio.com/docs/editor/userdefinedsnippets) file:

```
./quickbase-textmate snippets --file data/Formula_Functions_Reference.csv
```

To build the list of functions for the `support.function.quickbase` pattern:

```
./quickbase-textmate functions --file data/Formula_Functions_Reference.csv
```

To build the list of operators for the `keyword.operator.quickbase` pattern:

```
./quickbase-textmate operators --file data/Formula_Functions_Reference.csv
```
