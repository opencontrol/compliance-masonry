---
permalink: /add-images/
title: Add images
---
First, create an `images/` directory, and put
your images inside it. You may want to use
[jpegoptim](https://github.com/tjko/jpegoptim) or
[optipng](http://optipng.sourceforge.net/) to optimize your images. On OS X,
both are available via [Homebrew](http://brew.sh/).

Second — from your documents — you can reference your images as outlined below and abiding by
the advice in the [Accessibility
Guide](http://18f.github.io/accessibility/images/):

```
<img src="{%raw%}{{site.baseurl}}{%endraw%}/images/images.png" alt="Example of
an included image">
```

<img src="{{site.baseurl}}/images/images.png" alt="Example of an included image">

## Next steps

Click the _Update the Config File_ entry in the table of contents.
