{{ define "content" }}
<button onclick="clickHandler()">Launch DDNS-GO</button>
<script>
    function clickHandler() {
        const api = '/cgi-bin/App/ddns-go/api?action=StartInstance';
        fetch(api)
            .then(response => response.json())
            .then(data => console.log(data))
            .catch(error => console.error('Error:', error));
    }
</script>
{{ end }}
