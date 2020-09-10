# bulkdownload

> Download multiple files from a list

Bulkdownload downloads all files from a list. 
You can set how many files should be downloaded concurrently..

## Usage

> `dops [options] bulkdownload [options] [arguments...]`

**Category:** Web  
**Aliases:** `bd`  

### Options
```flags
--input FILE, -i FILE           |  load URLs from FILE (default: "urls.txt")  
--output DIR, -o DIR            |  save the downloaded files to DIR (default: current directory)  
--concurrent NUMBER, -c NUMBER  |  downloads NUMBER files concurrently (default: 3)  
```
## Examples

### Download all files from urls.txt, with 5 concurrent connections, to the current directory.

```command
dops bulkdownload -i urls.txt -c 5
```

