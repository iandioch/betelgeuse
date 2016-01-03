echo("<!doctype HTML>")
echo("<html><head><meta charset=\"utf-8\"/><title>Noah's Blog</title><link rel=\"stylesheet\" type=\"text/css\" href=\"../styles/stylesheet.css\"></head>");
echo("<body>");
echo("<div class=\"page-content\"><ul>");
for(var i = posts.length-1; i > posts.length-11 && i >= 0; i --){
	echo("<li><a href=\"../" + posts[i].Location + "\">" + posts[i].Meta.Title + "</a></li>");
}
echo("</ul></div>")
echo("</body>");
echo("</html>");