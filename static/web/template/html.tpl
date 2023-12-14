{{ define "html" }}
<html lang="zh">
<head>
    <title>{{.Title}}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=gb2312">
    <meta http-equiv="content-type" content="no-cache">
    <link href="/img/p2p.css" type="text/css" rel="stylesheet">
    <script type="text/javascript" src="/img/jq.js"></script>
    <script type="text/javascript" src="/html/assert/layui.all.js"></script>
</head>
<body>
{{ template "content" . }}
</body>
</html>
{{ end }}
