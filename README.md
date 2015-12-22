# betelgeuse

What better place for a programmer to keep their blog than in the same place they keep their code?

# What

This is a static site generator to create the pages for my blog. I want to post my blog entries in Markdown to Github, where they will be pulled and compiled into HTML automatically by my server.

# Why

I want to create a simple and easy to use blog system for myself. I chose Go for the system as it is reasonably low-level, full-featured and nice to use, and I want to practise it to use it more in future. I chose Javascript for the page generation as Go has a very nice library, otto, to interface with it. I also want to use my Javascript before I forget it all. In future, I could build in support for cooffeescript, typescript or equivalents if I want to get fancy.

# How
`setup.sh` should install everything needed to build the site.
`generate.go` will build the site, using Markdown files from ./posts, and outputting complete html files to ./site using the code in ./templates to generate the blog entry HTML pages.

# Where
The blog will be live at http://mycode.doesnot.run as soon as I get my act together.