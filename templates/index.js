echo("<!doctype HTML>")
echo("<html><head><meta charset=\"utf-8\"/><title>Noah's Blog</title></head>");
echo("<body>");
echo("<h1>Noah's Blog</h1>");
echo("<p>Recent posts:</p>");
echo("<ul>");
for(var i = posts.length-1; i > posts.length-11 && i >= 0; i --){
	echo("<li><a href=\"" + posts[i].Location + "\">" + posts[i].Meta.Title + "</a></li>");
}
echo("</ul>")
echo("</body>");
echo("</html>");