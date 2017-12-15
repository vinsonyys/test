<html>
<head>
<title>My Blog</title> <style>
#menu {
width: 200px; float: right;
} </style>
</head> 
<body>
  	<h1>{{.Post.Title}}</h1>
	{{.Post.Created}}<br/>
	{{.Post.Content}}
</body>
</html>