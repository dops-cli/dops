# extract

> Extracts data from other data

Extract contains multiple data extractors, which can be used to extract data from a file, URL or stdin.

## Usage

> `dops [options] extract subcommand [arguments...]`

**Category:** Data Analysis  
## Submodules

### text

> Extracts text from data

This can be used to extract text using a predefined or a custom regex.

#### Usage

> `dops [options] extract text [options] subcommand [arguments...]`

**Category:**   
**Aliases:** `t, string, strings, s`  

##### Options
```flags
--regex PATTERN, -r PATTERN  |  extracts matching strings with PATTERN  
--input FILE, -i FILE        |  use FILE as input  
--output DIR, -o DIR         |  outputs to directory DIR  
--append, -a                 |  append instead of overriding output (default: false)  
```
#### Submodules

#### predefined

> 



##### Usage

> `dops [options] text [options] predefined subcommand [arguments...]`

**Category:** Text Processing  
##### Submodules

##### email

> returns emails

The email command finds all email`s in the input and returns them.

###### Usage

> `dops [options] predefined email [options] [arguments...]`

**Category:** Text Processing  

####### Options
```flags
--input value, -i value   |  Input accepts a file, URL or stdin if not set  
--output value, -o value  |  Writes to a file, if not set it writes to stdout  
--append, -a              |  append instead of overriding output (default: false)  
```
## Examples

### Extract all email´s from INPUT

```command
dops extract text predefined email --input file.txt
```

