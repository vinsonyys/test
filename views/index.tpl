<html>
<head>
<title>My Blog</title> <style>
#menu {
width: 200px; float: right;
} </style>
</head> 
<body>
    <h1>Blog posts</h1>
    <h2>{{.Website}}</h2>
    <h3>{{.Email}}</h3>
    <ul>
    {{range .blogs}}
    <li>
    <a href="/view/{{.Id}}">{{.Title}}</a> from {{.Created}}
    <a href="/edit/{{.Id}}">Edit</a>
    <a href="/delete/{{.Id}}">Delete</a>
    </li> {{end}}
    </ul>
</body>
</html>