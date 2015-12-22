echo("<!doctype HTML>")
echo("<html><head><meta charset=\"utf-8\"/><title>" + posts[currId].Meta.Title + "</title></head>");
echo("<body>");
echo("<h1>");
echo(posts[currId].Meta.Title);
echo("</h1>");
echo("<p>");
echo(posts[currId].ParsedContent);
echo("</p>")
echo("</body>");
echo("</html>");