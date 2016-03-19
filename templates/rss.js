echo("<rss version=\"2.0\">");
echo("<channel>");
echo("<title>mycode.doesnot.run</title>");
echo("<link>http://mycode.doesnot.run</link>");
echo("<description>The blog of Noah Donnelly</description>");

for(var i = posts.length-1; i > posts.length-11 && i >= 0; i --){
	echo("<item>");
	echo("<title>" + posts[i].Meta.Title + "</title>");
	echo("<link>http://mycode.doesnot.run/" + posts[i].Location + "</link>");
	echo("</item>");
}
echo("</channel>");
echo("</rss>");
