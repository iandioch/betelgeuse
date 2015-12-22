echo("<!doctype HTML>")
echo("<html><head><meta charset=\"utf-8\"/><title>" + posts[currId].Meta.Title + "</title></head>");
echo("<body>");
echo("<h1>");
echo(posts[currId].Meta.Title);
echo("</h1>");
echo("<h3>" + posts[currId].Location + "</h3>");
echo("<p>");
echo(posts[currId].ParsedContent);
echo("</p>")
if(currId > 0){
	echo("<p>Previous post: <a href=\"../../../../" + posts[currId-1].Location +"\">" + posts[currId-1].Meta.Title + "</a></p>")
}
if(currId < posts.length - 1){
	echo("<p>Next post: <a href=\"../../../../" + posts[currId+1].Location +"\">" + posts[currId+1].Meta.Title + "</a></p>")
}
if(posts[currId].Meta.Categories.length > 0){
	echo("<p>This entry was posted on " + posts[currId].Date.Day + "/" + posts[currId].Date.Month +"/" + posts[currId].Date.Year + " in the following categories:");
	echo("<ul>");
	for (var i = 0; i < posts[currId].Meta.Categories.length; i ++) {
		echo("<li>" + posts[currId].Meta.Categories[i] + "</li>");
	}
	echo("</ul>");
	echo("</p>");
}
echo("</body>");
echo("</html>");