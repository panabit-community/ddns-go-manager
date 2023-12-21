{{ define "content" }}
<div>Current status: {{ .DDNSGO.Status }}</div>
<button onclick="onClickStartInstance()">Start DDNS-GO</button>
<button onclick="onClickStopInstance()">Stop DDNS-GO</button>
<button onclick="onClickRefresh()">Refresh</button>
<br>
<a href="./settings"><button>Jump to DDNS-GO</button></a>
<br>
<button onclick="onClickUpdateCertificates()">Update Certificates</button>
<div id="response"></div>
<script>
    function onClickStartInstance() {
        const api = '/cgi-bin/App/ddns-go/webmain?action=StartInstance';
        fetch(api)
            .then(response => response.json())
            .then(data => document.getElementById("response").innerHTML = data)
            .catch(error => document.getElementById("response").innerHTML = error);
    }
    function onClickStopInstance() {
        const api = '/cgi-bin/App/ddns-go/webmain?action=StopInstance';
        fetch(api)
            .then(response => response.json())
            .then(data => document.getElementById("response").innerHTML = data)
            .catch(error => document.getElementById("response").innerHTML = error);
    }
    function onClickRefresh() {
        location.reload();
    }
    function onClickUpdateCertificates() {
        const api = '/cgi-bin/App/ddns-go/webmain?action=UpdateCertificates';
        fetch(api)
            .then(response => response.json())
            .then(data => document.getElementById("response").innerHTML = data)
            .catch(error => document.getElementById("response").innerHTML = error);
    }
</script>
{{ end }}
