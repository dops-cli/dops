# open

> Opens a file or URL in the default program assigned to it

Open finds the standard program, assigned to a file, and opens it with the found programm.

## Usage

> `dops [options] open subcommand [arguments...]`

**Category:** Execute  
## Submodules

### url

> Open `URL` with the standard browser

This modules locates the standard browser of the system and opens a specific URL.

#### Usage

> `dops [options] open url [options] [arguments...]`

**Category:**   
**Aliases:** `browser, u`  

##### Options
```flags
--input value, -i value  |  input takes a URL from a path, URL or stdin if it's not set  
```
## Examples

### Opens marvinjwendt.com in your standard browser

```command
dops open url https://marvinjwendt.com
```

### Opens URL from stdin in your standard browser

```command
echo "https://marvinjwendt.com" | dops open url
```

### Returns an error if no standard browser is set

```command
dops open url https://marvinjwendt.com
```
<img src="/_assets/example_svg/dopsopenurlhttpsmarvinjwendtcom.svg">

