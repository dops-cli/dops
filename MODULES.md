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
--Float64List value	
--Int value	(default: 0)
--IntList value	
--Path value	
--String value	
--StringList value	
--Timestamp value	(default: <nil>)
```

## rename-files  

> Renames all selected files to a specific pattern  

Usage:  [options] [arguments...]<br/>
Aliases: `rf`<br/>
Category: IO

 ### Description

This module can be used to rename multiple files according to a specified pattern.
The pattern could be a timestamp, or the hashcode of the file, among others.

### Options

```
--directory PATH, --dir PATH, -d PATH	PATH in which the files should be renamed
--pattern OPTION, -p OPTION	Rename all files with OPTION
--recursive, -r	Rename files in subdirectories too (default: false)
--disablebackup, --db	Disable file name backups (default: false)
--loadbackup, -l, --lb	Reverts the filenames to the original (default: false)
```
