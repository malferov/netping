let pattern = /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\-]*[A-Za-z0-9])$/,
    $input = document.querySelector('input.main-input'),
    $btn = document.querySelector('button.main-btn');

$input.addEventListener('keydown', function (e) {
    if (e.which === 13) {
        return false;
    }
});

$input.addEventListener('keyup', function () {
    let this1 = this;
    if (!pattern.test(this1.value)) {
        $input.classList.remove('input-active');
        $btn.classList.remove('btn-active');
        while (this1.value.indexOf("..") !== -1) {
            this1.value = this1.value.replace('..', '.');
        }
    } else {
        let lastChar = this1.value.substr(this1.value.length - 1);
        if (lastChar == '.') {
            this1.value = this1.value.slice(0, -1);
        }
        let ip = this1.value.split('.');
        if (ip.length == 4 || pattern.test(this1.value)) {
            $input.classList.add('input-active');
            $btn.classList.add('btn-active');
        }
    }
});

$btn.onclick = function H(e) {
    e.preventDefault();
    document.querySelector('.js-output-content').style.display = "block";
};
