# random-generator

> Generate random values (string, integers, emails, etc..)

This module generates random values of specific types like string, integer, email etc.
You can set the number of generations and the seed.

!> **WARNING**
The generated random values are not cryptographically secure!  

## Usage

> `dops [options] random-generator [options] subcommand [arguments...]`

**Category:** Generators  
**Aliases:** `rg`  

### Options
```flags
--seed SEED, -s SEED  |  Uses SEED for the random generation (default: calculated by current time (nanoseconds))  
```
## Submodules

### string

> Generate random strings



#### Usage

> `dops [options] random-generator [options] string [options] [arguments...]`

**Category:**   
**Aliases:** `s`  

##### Options
```flags
--chars CHARS, -c CHARS     |  Use CHARS to generate a random string (default: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")  
--length LENGTH, -l LENGTH  |  Generate a random string of length LENGTH (default: 8)  
```
### integer

> Generate random integer



#### Usage

> `dops [options] random-generator [options] integer [options] [arguments...]`

**Category:**   
**Aliases:** `i, n, number`  

##### Options
```flags
--min NUMBER  |  Minimum NUMBER to be generated (default: 0)  
--max NUMBER  |  Maximum NUMBER to be generated (default: 2147483647)  
```
## Examples

### Generate a random string with 15 letters

```command
dops random-generator string --length 15
```
<img src="/_assets/example_svg/EpSJrKoCXduI.svg">

