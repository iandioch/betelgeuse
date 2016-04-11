echo("<?xml version=\"1.0\"?>\n")
echo("<rss version=\"2.0\">\n");
echo("\t<channel>\n");
echo("\t\t<title>mycode.doesnot.run</title>\n");
echo("\t\t<link>http://mycode.doesnot.run</link>\n");
echo("\t\t<description>Noah's blog &lt;3</description>\n");

for(var i = posts.length-1; i > posts.length-11 && i >= 0; i --){
	echo("\t\t<item>\n");
	echo("\t\t\t<title>" + encodeURIComponent(posts[i].Meta.Title) + "</title>\n");
	echo("\t\t\t<link>http://mycode.doesnot.run/" + posts[i].Location + "</link>\n");
	echo("\t\t</item>\n");
}
echo("\t</channel>\n");
echo("</rss>");
