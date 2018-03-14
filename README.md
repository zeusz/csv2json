# csv2json

https://github.com/zeusz/csv2json.git

**csv2json** is a simple command line program to convert CSV to JSON.

You can clone/fork/whateveryouwant this repository.

It is made in Go! So, you can just 
...
go get github.com/zeusz/csv2json
...

## USAGE

```
  -d string
        delimeter in csv (default ",")
  -i string
        input
  -o string
        output (default "output.json")
  -v    Verbose, some debug info

```
Example:

```
csvtojson -i gitlog.csv -o gitlog.json -v
``
