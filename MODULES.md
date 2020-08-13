
### bulkdownload  

> download multiple files from a list  

Usage:  [options] [arguments...]<br/>
Aliases: `bd`\
Category: web



#### Options

```
--input FILE, -i FILE	load URLs from FILE (default: "urls.txt")
--output DIR, -o DIR	save the downloaded files to DIR (default: current directory)
--concurrent NUMBER, -c NUMBER	downloads NUMBER files concurrently (default: 3)
```

### update  

> updates dops  

Usage:  [arguments...]<br/>

Category: dops





### extract-text  

> extracts text using regex from a file  

Usage:  [options] [arguments...]<br/>

Category: text processing



#### Options

```
--regex PATTERN, -r PATTERN	extracts matching strings with PATTERN
--input FILE, -i FILE	use FILE as input
--output DIR, -o DIR	outputs to directory DIR
--stdout, -s	prints output to stdout instead of writing to a file (default: false)
```
