{{ define "content" }}
<!doctype html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="../static/output.css" rel="stylesheet">
    <style>
        @keyframes slideDown {
            from {
                transform: translateY(-100%);
                opacity: 0;
            }
            to {
                transform: translateY(0);
                opacity: 1;
            }
        }

        @keyframes fadeOut {
            from {
                opacity: 1;
            }
            to {
                opacity: 0;
            }
        }
    </style>
</head>
<body>
<div class="flex items-center h-12 bg-indigo-500 text-white shadow-md">
    <p class="ml-4">DDNS-GO</p>
    <div class="w-4 ml-auto mr-4 hover:animate-spin" id="js-refresh">
        <svg xmlns="http://www.w3.org/2000/svg"  viewBox="0 0 1024 1024" data-v-ea893728=""><path fill="currentColor" d="M771.776 794.88A384 384 0 0 1 128 512h64a320 320 0 0 0 555.712 216.448H654.72a32 32 0 1 1 0-64h149.056a32 32 0 0 1 32 32v148.928a32 32 0 1 1-64 0v-50.56zM276.288 295.616h92.992a32 32 0 0 1 0 64H220.16a32 32 0 0 1-32-32V178.56a32 32 0 0 1 64 0v50.56A384 384 0 0 1 896.128 512h-64a320 320 0 0 0-555.776-216.384z"></path></svg>
    </div>
</div>
<div class="mx-auto my-10 w-80 shadow-lg rounded">
    <div class="pt-4 pb-2 border-b">
        <div class="pl-4 text-slate-800">插件状态</div>
    </div>
    <div class="px-8 py-14 text-center">
        <div class="flex items-center justify-center">
            <span id="status" class="text-3xl text-red-500">{{ .DDNSGO.Status }}</span>
            <div class="flex items-center ml-4">
                <label for="toggle" class="inline-flex relative items-center cursor-pointer">
                    <input type="checkbox" id="toggle" class="sr-only peer" {{if eq .DDNSGO.Status "Opened"}}checked{{end}}>
                    <div class="w-12 h-6 bg-red-500 rounded-full peer-checked:bg-green-500 peer-focus:ring-4 peer-focus:ring-slate-100 transition duration-300 ease-in-out"></div>
                    <div class="dot absolute left-0.5 top-0.5 bg-white w-5 h-5 rounded-full transition-transform duration-300 ease-in-out peer-checked:translate-x-6"></div>
                </label>
            </div>
        </div>
        <div class="text-orange-500 mt-2 opacity-0" id="js-error-tips"></div>
    </div>
</div>
<div class="text-center">
    <button class="w-80 rounded text-slate-600 hover:border-indigo-400 hover:text-indigo-600 py-2 border-slate-300 border-solid border-2">
        <a href="./settings">进入设置</a>
    </button>
</div>
<div class="mx-auto my-4 flex items-center justify-center w-80 rounded bg-indigo-500 hover:bg-indigo-600 text-white py-2" id="js-update-certificate">
    <div class="hidden" id="js-loading-icon">
        <svg xmlns="http://www.w3.org/2000/svg"  viewBox="0 0 1024 1024" data-v-ea893728=""><path fill="currentColor" d="M771.776 794.88A384 384 0 0 1 128 512h64a320 320 0 0 0 555.712 216.448H654.72a32 32 0 1 1 0-64h149.056a32 32 0 0 1 32 32v148.928a32 32 0 1 1-64 0v-50.56zM276.288 295.616h92.992a32 32 0 0 1 0 64H220.16a32 32 0 0 1-32-32V178.56a32 32 0 0 1 64 0v50.56A384 384 0 0 1 896.128 512h-64a320 320 0 0 0-555.776-216.384z"></path></svg>
    </div>
    更新证书
</div>
<div id="notification" class="hidden fixed top-5 left-1/3 text-white px-4 py-2 rounded shadow-lg"></div>
<script src="../static/index.js"></script>
</body>
</html>
{{ end }}
