I don't use this any more, because of course rolling my own blog software was a bad idea. But it was fun while it lasted! You can see my new shiny blog [here](http://mycode.doesnot.run).

# betelgeuse

What better place for a programmer to keep their blog than in the same place they keep their code?

## What

This is a static site generator to create the pages for my blog. I want to post my blog entries in Markdown to Github, where they will be pulled and compiled into HTML automatically by my server.

## Why

I want to create a simple and easy to use blog system for myself. I chose Go for the system as it is reasonably low-level, full-featured and nice to use, and I want to practise it to use it more in future. I chose Javascript for the page generation as Go has a very nice library, [otto](https://github.com/robertkrimen/otto), to interface with it. I also want to use my Javascript before I forget it all. In future, I could throw in support for coffeescript, typescript or equivalents if I want to get fancy.

## How

`setup.sh` should install everything needed to build the site.

`generate.go` will build the site, using Markdown files from ./posts, and outputting complete html files to ./site using the code in ./templates to generate the blog entry HTML pages.

## Where

The void of the dead web.

