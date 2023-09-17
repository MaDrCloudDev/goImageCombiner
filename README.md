# go_image_combiner

Build the program:
```
go build -o imagecombiner
```
Run the program with your images:
```
./imagecombiner images/image1.png images/image2.png output.png
```

We're making the images the same size (which sometimes results in offset images, I'll try to fix that), then combining the images using every-other column of pixels from each image. I adapted it from my [rust_image_combiner](https://github.com/MaDrCloudDev/rust_image_combiner), which was a code-along to one of the freeCodeCamp tutorials.