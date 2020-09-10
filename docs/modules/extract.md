# extract

> Extracts data from other data

Extract contains multiple data extractors, which can be used to extract data from a file, URL or stdin.

## Usage

> `dops [options] extract subcommand [arguments...]`

**Category:** Data Analysis  
# text

> Extracts text from data

This can be used to extract text using a predefined or a custom regex.

## Usage

> `dops [options] extract text [options] [arguments...]`

**Category:**   
**Aliases:** `t, string, strings, s`  

### Options
```flags
--regex PATTERN, -r PATTERN  |  extracts matching strings with PATTERN  
--input FILE, -i FILE        |  use FILE as input  
--output DIR, -o DIR         |  outputs to directory DIR  
--append, -a                 |  append instead of overriding output (default: false)  
```
