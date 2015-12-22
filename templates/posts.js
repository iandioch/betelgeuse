echo("<!doctype HTML>")
echo("<html><head><meta charset=\"utf-8\"/><title>" + posts[currId].Meta.Title + "</title></head>");
echo("<body>");
echo("<h1>");
echo(posts[currId].Meta.Title);
echo("</h1>");
echo("<p>");
echo(posts[currId].ParsedContent);
echo("</p>")
if(posts[currId].Meta.Categories.length > 0){
	echo("<p>This entry was posted under:");
	echo("<ul>");
	for (var i = 0; i < posts[currId].Meta.Categories.length; i ++) {
		echo("<li>" + posts[currId].Meta.Categories[i] + "</li>");
	}
	echo("</ul>");
	echo("</p>");
}
echo("</body>");
echo("</html>");