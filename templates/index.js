echo("<!doctype HTML>")
echo("<html><head><meta charset=\"utf-8\"/><title>Noah Ó Donnaile</title><link href='https://fonts.googleapis.com/css?family=Slabo+27px&subset=latin,latin-ext' rel='stylesheet' type='text/css'><link rel=\"stylesheet\" type=\"text/css\" href=\"styles/stylesheet.css\"></head>");
echo("<body>");
echo("<div class=\"title\"><h1>Noah Ó Donnaile</h1><p>Hi, I'm Noah, and I'm from Monaghan, Ireland. I'm studying <a href=\"tags/cpssd.html\">Computational Problem Solving & Software Development</a> in <a href=\"http://dcu.ie\">Dublin City University</a>. I'm the current chairperson of <a href=\"http://redbrick.dcu.ie\">Redbrick</a>, DCU's Networking Society. I'm into languages, competitive programming, and poor films. See my resumé <a href=\"https://docs.google.com/document/d/1Nl6RUxY6QRPBq0aYQ85JeKBV4yBSr4X-rNM6PbhoVP4/\">here</a>. Read about my <a href=\"categories/projects.html\">projects</a>, check out my <a href=\"http://github.com/iandioch\">Github</a>, and follow me on <a href=\"http://twitter.com/iandioch\">Twitter</a>.</p></div>");
echo("<div class=\"page-content\">");
echo("<p>Recent posts:</p>");
//echo("<ul>");
for(var i = posts.length-1; i > posts.length-11 && i >= 0; i --){
	//echo("<li>");

	echo("<h2><a href=\"" + posts[i].Location + "\">" + posts[i].Meta.Title + "</a></h2>");
	//echo("<ul>")
	var cats = posts[i].Meta.Categories;
	if (cats.length > 0) {

		//echo("<li>Categories: <a href=\"./categories/" + cats[0] + ".html\">" + cats[0] + "</a>");
        echo("<p>Categories: <a href=\"./categories/" + cats[0] + ".html\">" + cats[0] + "</a>");

		var cats = posts[i].Meta.Categories;
		for (var j = 1; j < cats.length; j ++){
			//echo(posts[j].Meta.Categories);
			echo(", <a href=\"./categories/" + cats[j] + ".html\">" + cats[j] + "</a>");
			//echo("lol");
		}
        echo("</p>");
		//echo("</li>");
	}

	var tags = posts[i].Meta.Tags;
	if (tags.length > 0) {

		//echo("<li>Tags: <a href=\"./tags/" + tags[0] + ".html\">" + tags[0] + "</a>");
        echo("<p>Tags: <a href=\"./categories/" + tags[0] + ".html\">" + tags[0] + "</a>");

		var cats = posts[i].Meta.Categories;
		for (var j = 1; j < tags.length; j ++){
			//echo(posts[j].Meta.Categories);
			echo(", <a href=\"./tags/" + tags[j] + ".html\">" + tags[j] + "</a>");
			//echo("lol");
		}
        echo("</p>");
		//echo("</li>");
	}
    echo("<br>");
	
	//echo("</ul>");
	//echo("</li>");
}
//echo("</ul></div>");
echo("</div>");
echo("<div><p>Subscribe to <a href=\"rss\">my RSS feed</a></p></div>");
echo("</body>");
echo("</html>");
