---
title: Twitter Data!
categories: ["updates", "projects"]
tags: ["redbrick", "programming"]
---

Last weekend I went to Galway with Redbrick, for the first intersocs hackathon. It was good fun! Galway was very nice. I heard from the waiter in "Rockin' Joes" that there are 3 series of both Black Sails and The 100, whereas I'd only seen the first series of each. Plenty of dodgy TV to keep me busy there. 

However, the reason I was in Galway was to do analytics on some Twitter data. I'd had a Digital Ocean droplet running for about 10 days in the leadup to my trip, continuously logging all the geo-tagged data sent in the vicinity of the island of Ireland into a MongoDB collection. In the end, I had around 500k tweets; with half from the Republic of Ireland and the other half from Northern Ireland and some bits of Wales and Scotland that got caught in the bounding box. The (slightly shoddy) code for logging all the tweets is available [here](https://github.com/iandioch/mumsy).

I was very much looking forward to doing some analytics!

The first thing I wanted to do was physically map the tweets. The Twitter API gave coordinates for every tweet, so it was just a matter of lining that up with a map, right?

I looked into how best to do that. I didn't want to just fudge it by manually lining up a map image with the plotted data, so I used matplotlib with Python. It's a very mature and fully-featured tool, with a long and arduous install process when it's used with [basemap](http://matplotlib.org/basemap/). After a lot of work, I [had a result](https://twitter.com/iandioch/status/706148340581400576)! It wasn't what I expected at all. In Northern Ireland, the location data is great! However, in the South, as far as I can tell, it rounds every location off to the nearest county. I'd presume this is because Twitter doesn't have any data on placenames and the like in Ireland, but that's just a guess.

I then wanted a pretty graph to map the tweets over time. As I was logging the data, I was printing out how many tweets were being posted each minute. It was clear there was a pattern in the data; tweet frequencies increased during lunchtime, for example. However, it was also interesting during the election count to see the tweet numbers increase dramatically as news was announced. I was going to use d3 or something to graph the data, but I knew I didn't have enough time that weekend, so instead I opted to use [Processing](http://processing.org), with which I have years of experience. The built-in loadStrings() method couldn't handle the amount of data I had, so I wrote some in-between scripts in Python and C++ to make the data easier to handle, and then I used a plain Java BufferedReader to load the file.

In the end, I got [this graph](https://twitter.com/iandioch/status/706192733166309376)! It's easy to see the peak for the election count, which lines up with my own observation.

I would love to study the data a bit more. I had to shut down the server that was logging the tweets, as it ran out of disc space and crashed so I could no longer SSH in. However, the script is still sitting there on Github to start the ball rolling again. I'm interested in converting it to use an SQL database, which would be much more space-efficient, and would be a chance to improve my database knowledge a bit more. I only used Mongo in the first place because I had already written some Mongo database interaction in Python for another project that could almost direcltly be applied here.

I also want to do some actual word-based analysis. My original goal with the Twitter data was to see if dialectal language differences were visible even on such a small map as Ireland; in particular, I was going to count the usages of the words "mom", "mam", "mum", and their variations, and see if "mam" was more frequently used in the countryside and "mum" in the city, or anything like that.

In any case, I had lots of playing with the biggest dataset I've had to date. I think I'll set up a new server running in the near future to log even more data to fiddle with.

Ádh mór,

Noah
