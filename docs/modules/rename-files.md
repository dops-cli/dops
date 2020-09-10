# rename-files

> Renames all selected files to a specific pattern

This module can be used to rename multiple files according to a specified pattern.
The pattern could be a timestamp, or the hashcode of the file, among others.

## Usage

> `dops [options] rename-files [options] [arguments...]`

**Category:** IO  
**Aliases:** `rf`  

### Options
```flags
--directory PATH, --dir PATH, -d PATH  |  PATH in which the files should be renamed  
--pattern OPTION, -p OPTION            |  Rename all files with OPTION  
--recursive, -r                        |  Rename files in subdirectories too (default: false)  
--disablebackup, --db                  |  Disable file name backups (default: false)  
--loadbackup, -l, --lb                 |  Reverts the filenames to the original (default: false)  
```
## Examples

