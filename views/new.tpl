<html>
<head>
<title>My Blog</title> <style>
#menu {
width: 200px; float: right;
} </style>
</head> 
<body>
  	<h1>New Blog Post</h1>
	<form action="/add" method="post">
		标题:<input type="text" name="title"><br>
		内容:<textarea name="content" colspan="3" rowspan="10"></textarea> 
		<input type="submit">
	</form>
</body>
</html>