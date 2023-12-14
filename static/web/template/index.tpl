{{ define "content" }}
<button onclick="clickHandler()">Launch DDNS-GO</button>
<script>
    function clickHandler() {
        window.location.href = "/cgi-bin/App/ddns-go-manager/api?action=StartInstance";
    }
</script>
{{ end }}
