# Optimaze
A tiny image optimizer written in GO (currently supported JPEG)
(tiny step towards my GOPHER jouerny :) )

## Installation

Download binary file and put file path to your environment variable
Make a folder 'input' in your working directory and put your all JPG images into it
then type
```bash
optimaze width height quality
```
for default width/height just put a '0'
quality must be 0-100

```bash
optimaze 0 0 50
```

Reduced 92% size without a big impact on quality

| Original image (2,759 KB) | Optimized image (212 KB) |
|:--:|:--:|
| ![original](https://github.com/SharifClick/optimaze/blob/master/image.jpg) | ![optimized](https://github.com/SharifClick/optimaze/blob/master/image_reduced.jpg) | 

