window.onload = () => {
    initButtonEvent()
}

function initButtonEvent() {
    // 切换状态
    this.toggleStatus()

    // 更新证书
    document.getElementById('js-update-certificate').addEventListener('click', () => {
        this.onClickUpdateCertificates()
    })

    // 刷新页面
    document.getElementById('js-refresh').addEventListener('click', () => {
        window.location.reload()
    })
}

function toggleStatus() {
    const toggle = document.getElementById('toggle')
    const status = document.getElementById('status')
    toggle.addEventListener('change', function() {
        toggle.style.opacity = '0'
        if (this.checked) {
            onClickStartInstance(status, toggle)
        } else {
            onClickStopInstance(status, toggle)
        }
    })
}

function onClickStartInstance(status, toggle) {
    const api = '/cgi-bin/App/ddns-go/webmain?action=StartInstance'
    fetch(api)
        .then(response => response.json())
        .then(data => {
            status.textContent = 'Opened'
            status.className = 'text-3xl text-green-500'
            toggle.style.opacity = '1'
            setOrClearTips()
        })
        .catch(error => {
            setOrClearTips(error)
            toggle.checked = false
            toggle.style.opacity = '1'
        })
}

function onClickStopInstance(status, toggle) {
    const api = '/cgi-bin/App/ddns-go/webmain?action=StopInstance'
    fetch(api)
        .then(response => response.json())
        .then(data => {
            status.textContent = 'Closed'
            status.className = 'text-3xl text-red-500'
            toggle.style.opacity = '1'
            setOrClearTips()
        })
        .catch(error => {
            setOrClearTips(error)
            toggle.checked = true
            toggle.style.opacity = '1'
        })
}

function onClickUpdateCertificates() {
    const api = '/cgi-bin/App/ddns-go/webmain?action=UpdateCertificates'
    const loadingIcon = document.getElementById('js-loading-icon')
    loadingIcon.className = 'h-5 w-5 mr-3 block animate-spin'
    fetch(api)
        .then(response => response.json())
        .then(data => {
            // success
            loadingIcon.className = 'hidden'
            messageboxTips('更新成功', 'bg-green-400')
        })
        .catch(error => {
            // error
            loadingIcon.className = 'hidden'
            messageboxTips(error, 'bg-red-400')
        })
}

function setOrClearTips(error) {
    const errorTips = document.getElementById('js-error-tips')
    if (error) {
        errorTips.textContent = error
        errorTips.className = 'text-red-500 mt-2 opacity-100'
    } else {
        errorTips.textContent = ''
        errorTips.className = 'text-red-500 mt-2 opacity-0'
    }
}

function messageboxTips(tips, className = '') {
    const notification = document.getElementById('notification')
    if (className) {
        notification.classList.add(className)
    }
    notification.textContent = tips
    notification.classList.remove('hidden')
    notification.style.animation = 'slideDown 0.5s forwards'

    // 三秒后开始关闭动画
    setTimeout(() => {
        notification.style.animation = 'fadeOut 0.5s forwards'

        // 动画结束后隐藏元素
        notification.addEventListener('animationend', () => {
            notification.classList.add('hidden')
        })
    }, 3000)
}
