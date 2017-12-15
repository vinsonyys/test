<html>
<head>
<title>My Blog</title> <style>
#menu {
width: 200px; float: right;
} </style>
</head> 
<body>
  	<h1>Edit {{.Post.Title}}</h1>
	<h1>New Blog Post</h1>
	<form action="" method="post">
		标题:<input type="text" name="title" value="{{.Post.Title}}"><br> 
		内容:<textarea name="content" colspan="3" rowspan="10">{{.Post.Content}}</textarea>
		<input type="hidden" name="id" value="{{.Post.Id}}">
		<input type="submit">
	</form>
</body>
</html>