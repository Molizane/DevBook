$('#login').on('submit', fazerLogin)

function fazerLogin(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val().trim(),
            senha: $('#senha').val().trim()
        }
    }).done(function () { // 200, 201, 204
        window.location = '/home';
    }).fail(function (erro) { // 400 401 403 4094 500
        Swal.fire("ERRO!", "Erro ao fazer login.", "error");
    });
}
