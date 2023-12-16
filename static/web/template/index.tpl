{{ define "content" }}
<div>Current status: {{ .DDNSGO.Status }}</div>
<button onclick="onClickStartInstance()">Start DDNS-GO</button>
<button onclick="onClickStopInstance()">Stop DDNS-GO</button>
<button onclick="onClickRefresh()">Refresh</button>
<a href="http://192.168.0.200:9876" target="_blank">
    <button>Jump to DDNS-GO</button>
</a>
<div>Username: {{ .DDNSGO.Username }}</div>
<div>Password: {{ .DDNSGO.Password }}</div>
<script>
    function onClickStartInstance() {
        const api = '/cgi-bin/App/ddns-go/api?action=StartInstance';
        fetch(api)
            .then(response => response.json())
            .then(data => console.log(data))
            .catch(error => console.error('Error:', error));
    }

    function onClickStopInstance() {
        const api = '/cgi-bin/App/ddns-go/api?action=StopInstance';
        fetch(api)
            .then(response => response.json())
            .then(data => console.log(data))
            .catch(error => console.error('Error:', error));
    }

    function onClickRefresh() {
        location.reload();
    }
</script>
{{ end }}
