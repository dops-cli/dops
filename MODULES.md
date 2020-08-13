
### bulkdownload  

> Download multiple files from a list  

Usage:  [options] [arguments...]<br/>
Aliases: `bd`<br/>
Category: web

 #### Description

Bulkdownload downloads all files from a list. 
You can set how many files should be downloaded concurrently..

#### Options

```
--input FILE, -i FILE	load URLs from FILE (default: "urls.txt")
--output DIR, -o DIR	save the downloaded files to DIR (default: current directory)
--concurrent NUMBER, -c NUMBER	downloads NUMBER files concurrently (default: 3)
```

### extract-text  

> Extracts text using regex from a file  

Usage:  [options] [arguments...]<br/>

Category: text processing

 #### Description

Extract-text can be used to extract text from a file using regex patterns.

#### Options

```
--regex PATTERN, -r PATTERN	extracts matching strings with PATTERN
--input FILE, -i FILE	use FILE as input
--output DIR, -o DIR	outputs to directory DIR
--stdout, -s	prints output to stdout instead of writing to a file (default: false)
```

### update  

> Updates the dops tool  

Usage:  [arguments...]<br/>

Category: dops

 #### Description

NOTICE: This module is in progress. But you can already see it's usage for further use!


