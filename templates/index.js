echo("<!doctype HTML>")
echo("<html><head><meta charset=\"utf-8\"/><title>Noah's Blog</title><link href='https://fonts.googleapis.com/css?family=Slabo+27px&subset=latin,latin-ext' rel='stylesheet' type='text/css'><link rel=\"stylesheet\" type=\"text/css\" href=\"../styles/stylesheet.css\"></head>");
echo("<body>");
echo("<div class=\"title\"><h1>NOAH'S BLOG</h1></div>");
echo("<div class=\"page-content\">");
echo("<p>Recent posts:</p>");
echo("<ul>");
for(var i = posts.length-1; i > posts.length-11 && i >= 0; i --){
	echo("<li>");

	echo("<h5><a href=\"" + posts[i].Location + "\">" + posts[i].Meta.Title + "</a></h5>");
	echo("<ul>")
	var cats = posts[i].Meta.Categories;
	if (cats.length > 0) {

		echo("<li>Categories: <a href=\"./categories/" + cats[0] + ".html\">" + cats[0] + "</a>");

		var cats = posts[i].Meta.Categories;
		for (var j = 1; j < cats.length; j ++){
			//echo(posts[j].Meta.Categories);
			echo(", <a href=\"./categories/" + cats[j] + ".html\">" + cats[j] + "</a>");
			//echo("lol");
		}
		echo("</li>");
	}

	var tags = posts[i].Meta.Tags;
	if (tags.length > 0) {

		echo("<li>Tags: <a href=\"./tags/" + tags[0] + ".html\">" + tags[0] + "</a>");

		var cats = posts[i].Meta.Categories;
		for (var j = 1; j < tags.length; j ++){
			//echo(posts[j].Meta.Categories);
			echo(", <a href=\"./tags/" + tags[j] + ".html\">" + tags[j] + "</a>");
			//echo("lol");
		}
		echo("</li>");
	}
	
	echo("</ul>");
	echo("</li>");
}
echo("</ul></div>")
echo("</body>");
echo("</html>");