# Acronym finder written in Go

## Usage
Assuming you have a file named `inputfile.txt` containing a single unique word for each line:

`go run main.go inputfile.txt`

All words that have acronyms will be printed to stdout along with their acronyms, e.g.
```
bra, bar
dro, ord, rod
løst, støl
...
```

## Build
* `go build`
* `./acronyms inputfile.txt`

## Performance analysis
```bash
time for i in {1..10}; do ./acronyms /usr/share/dict/words > /dev/null; done

real    0m1.304s
user    0m1.424s
sys     0m0.100s
```
