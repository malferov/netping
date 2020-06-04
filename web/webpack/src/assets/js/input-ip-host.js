let pattern = /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\-]*[A-Za-z0-9])$/,
    $input = $('input.main-input'),
    $btn = $('button.main-btn');

$input.keypress(function (e) {
    if (e.which === 13) {
        return false;
    }
}).keyup(function () {
    let this1 = $(this);
    if (!pattern.test(this1.val())) {
        $input.removeClass('input-active');
        $btn.removeClass('btn-active');
        while (this1.val().indexOf("..") !== -1) {
            this1.val(this1.val().replace('..', '.'));
        }
    } else {
        let lastChar = this1.val().substr(this1.val().length - 1);
        if (lastChar == '.') {
            this1.val(this1.val().slice(0, -1));
        }
        let ip = this1.val().split('.');
        if (ip.length == 4 || pattern.test(this1.val())) {
            $input.addClass('input-active');
            $btn.addClass('btn-active');
        }
    }
});

$btn.on('click', function H(e) {
    e.preventDefault();
    $('.js-output-content').css('display', 'block');
});