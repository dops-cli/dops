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

> Use a predefined regex to extract strings

Use the predefined submodule to choose from a set of regexes, which you can use to extract strings either from a website, a file or stdin.

##### Usage

> `dops [options] text [options] predefined subcommand [arguments...]`

**Category:** Text Processing  
##### Submodules

##### email

> returns emails

The email command finds all email`s in the input and returns them.

Regex: 
(?:[a-z0-9!#$%&'*+/=?^_\x60{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_\x60{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])

###### Usage

> `dops [options] predefined email [options] [arguments...]`

**Category:** Text Processing  

####### Options
```flags
--input value, -i value   |  Input accepts a file, URL or stdin if not set  
--output value, -o value  |  Writes to a file, if not set it writes to stdout  
--append, -a              |  append instead of overriding output (default: false)  
```
##### ipv4

> returns IP Version 4 addresses

The ipv4 command finds all ipv4`s in the input and returns them.

Regex: 
\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b

###### Usage

> `dops [options] predefined ipv4 [options] [arguments...]`

**Category:** Text Processing  

####### Options
```flags
--input value, -i value   |  Input accepts a file, URL or stdin if not set  
--output value, -o value  |  Writes to a file, if not set it writes to stdout  
--append, -a              |  append instead of overriding output (default: false)  
```
##### ipv6

> returns IP Version 6 addresses

The ipv6 command finds all ipv6`s in the input and returns them.

Regex: 
\b(?:[a-fA-F0-9]{1,4}:){7}[a-fA-F0-9]{1,4}\b

###### Usage

> `dops [options] predefined ipv6 [options] [arguments...]`

**Category:** Text Processing  

####### Options
```flags
--input value, -i value   |  Input accepts a file, URL or stdin if not set  
--output value, -o value  |  Writes to a file, if not set it writes to stdout  
--append, -a              |  append instead of overriding output (default: false)  
```
##### ipaddress

> returns IP addresses

The ipaddress command finds all ipaddress`s in the input and returns them.

Regex: 
\b(((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))|(?:[a-fA-F0-9]{1,4}:){7}[a-fA-F0-9]{1,4})\b

###### Usage

> `dops [options] predefined ipaddress [options] [arguments...]`

**Category:** Text Processing  
**Aliases:** `ip`  

####### Options
```flags
--input value, -i value   |  Input accepts a file, URL or stdin if not set  
--output value, -o value  |  Writes to a file, if not set it writes to stdout  
--append, -a              |  append instead of overriding output (default: false)  
```
##### url

> returns url

The url command finds all url`s in the input and returns them.

Regex: 
(?i)(http://|https://|www.)([-a-zA-Z0-9]{1,63}\.[a-zA-Z0-9()]{1,6}\b)([-a-zA-Z0-9()@:%_\+.~#?&//=]*)?

###### Usage

> `dops [options] predefined url [options] [arguments...]`

**Category:** Text Processing  

####### Options
```flags
--input value, -i value   |  Input accepts a file, URL or stdin if not set  
--output value, -o value  |  Writes to a file, if not set it writes to stdout  
--append, -a              |  append instead of overriding output (default: false)  
```
##### image url

> returns image url

The image url command finds all image url`s in the input and returns them.

Regex: 
(?m)(http(s?)://)([/\d\w\S-_])*\.(?:jpg|gif|png)

###### Usage

> `dops [options] predefined image url [options] [arguments...]`

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

### Extract all ipv4´s from INPUT

```command
dops extract text predefined ipv4 --input file.txt
```

### Extract all ipv6´s from INPUT

```command
dops extract text predefined ipv6 --input file.txt
```

### Extract all ipaddress´s from INPUT

```command
dops extract text predefined ipaddress --input file.txt
```

### Extract all url´s from INPUT

```command
dops extract text predefined url --input file.txt
```

### Extract all image url´s from INPUT

```command
dops extract text predefined image url --input file.txt
```

