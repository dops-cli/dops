# DOPS - Modules
## bulkdownload  

> Download multiple files from a list  

Usage:  [options] [arguments...]<br/>
Aliases: `bd`<br/>
Category: Web

 ### Description

Bulkdownload downloads all files from a list. 
You can set how many files should be downloaded concurrently..

### Options

```
--input FILE, -i FILE	load URLs from FILE (default: "urls.txt")
--output DIR, -o DIR	save the downloaded files to DIR (default: current directory)
--concurrent NUMBER, -c NUMBER	downloads NUMBER files concurrently (default: 3)
```

## extract-text  

> Extracts text using regex from a file  

Usage:  [options] [arguments...]<br/>

Category: Text Processing

 ### Description

Extract-text can be used to extract text from a file using regex patterns.

### Options

```
--regex PATTERN, -r PATTERN	extracts matching strings with PATTERN
--input FILE, -i FILE	use FILE as input
--output DIR, -o DIR	outputs to directory DIR
--append, -a	append instead of overriding output (default: false)
```

## update  

> Updates the dops tool  

Usage:  [arguments...]<br/>

Category: Dops

 ### Description

NOTICE: This module is in progress. But you can already see it's usage for further use!



## demo  

> Demo module of dops  

Usage:  [options] [arguments...]<br/>

Category: Dops

 ### Description

NOTICE: This module does nothing, except showing all possible flags for an interactive demo.

### Options

```
--Boolean	(default: false)
--Duration value	(default: 0s)
--Float64 value	(default: 0)
--Float64 value	
--Int value	(default: 0)
--Int value	
--Path value	
--String value	
--String value	
--Timestamp value	(default: <nil>)
```
