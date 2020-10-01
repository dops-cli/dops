# image

> Image modification

This module has a list of modules to modify images.

## Usage

> `dops [options] image subcommand [arguments...]`

**Category:** Image Processing  
## Submodules

### watermark

> Adds a watermark to an image

This module watermark adds a watermark to one or more images from the input with a custom text on one of the corners.

#### Usage

> `dops [options] image watermark [options] [arguments...]`

**Category:**   
**Aliases:** `wm`  

##### Options
```flags
--location value, -l value   |  Watermark location  
--text value, -t value       |  Watermark text  
--size value, -s value       |  Watermark size (default: 12)  
--color value, -c value      |  Watermark color (default: "#ffffff")  
--opacity value, --op value  |  Watermark opacity - range 0-100 (default: 100)  
--input FILE, -i FILE        |  use FILE as input  
--glob GLOB, -g GLOB         |  uses a GLOB pattern to input multiple files  
--output DIR, -o DIR         |  outputs to directory DIR  
```
## Examples

### Adds a watermark to the example.jpg image and saves it as example_watermarked.png

```command
dops image watermark --input example.jpg --text "example watermark text" --location "top left" --opacity 50 --output "example_watermarked.png"
```

### Adds a watermark to every image with .png ending in this path

```command
dops image watermark --glob c/images/*.png --text "example watermark text" --location "top left" --opacity 50
```

