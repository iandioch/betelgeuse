echo("<!doctype HTML>")
echo("<html><head><meta charset=\"utf-8\"/><title>" + posts[currId].Meta.Title + "</title><link href='https://fonts.googleapis.com/css?family=Slabo+27px&subset=latin,latin-ext' rel='stylesheet' type='text/css'><link rel=\"stylesheet\" type=\"text/css\" href=\"../../../../styles/stylesheet.css\"></head>");
echo("<body>");
echo("<div class=\"title\"><h1>");
echo(posts[currId].Meta.Title.toUpperCase());
echo("</h1>");
echo("<div class=\"meta-data\">");
echo("<p>Posted on " + posts[currId].Date.Day + "/" + posts[currId].Date.Month + "/" + posts[currId].Date.Year + " with tags [");
for (var i = 0; i < posts[currId].Meta.Tags.length; i ++) {
	//echo("<a href=\"../../../../tags/" + posts[currId].Meta.Tags[i] + ".html\">" + posts[currId].Meta.Tags[i] + "</a> ");
}
if(posts[currId].Meta.Tags.length > 0){
	echo("<a href=\"../../../../tags/" + posts[currId].Meta.Tags[0] + ".html\">" + posts[currId].Meta.Tags[0] + "</a>");
}
for (var i = 1; i < posts[currId].Meta.Tags.length; i ++) {
	echo(", <a href=\"../../../../tags/" + posts[currId].Meta.Tags[i] + ".html\">" + posts[currId].Meta.Tags[i] + "</a>");
}
if(posts[currId].Meta.Categories.length > 0){
	if(posts[currId].Meta.Categories.length == 1){
		echo("] in category [");
	}else{
		echo("] in categories [");
	}
	echo("<a href=\"../../../../categories/" + posts[currId].Meta.Categories[0] + ".html\">" + posts[currId].Meta.Categories[0] + "</a>");
	for (var i = 1; i < posts[currId].Meta.Categories.length; i ++) {
		echo(", <a href=\"../../../../categories/" + posts[currId].Meta.Categories[i] + ".html\">" + posts[currId].Meta.Categories[i] + "</a>");
	}
}
echo("]</p>");
echo("</div>")
echo("</div>");
echo("<div class=\"page-content\"><p>");
echo(posts[currId].ParsedContent);
echo("</p></div>");
echo("<div class=\"post-links\">");
if(currId > 0){
	echo("<p>Previous post: <a href=\"../../../../" + posts[currId-1].Location +"\">" + posts[currId-1].Meta.Title + "</a></p>")
}
if(currId < posts.length - 1){
	echo("<p>Next post: <a href=\"../../../../" + posts[currId+1].Location +"\">" + posts[currId+1].Meta.Title + "</a></p>")
}
echo("</div>");
echo("</body>");
echo("</html>");