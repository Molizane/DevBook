$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val().trim() != $('#confirmar-senha').val().trim()) {
        Swal.fire("ERRO!", "As senhas não coincidem", "error");
    } else {
        $.ajax({
            url: "/usuarios",
            method: "POST",
            data: {
                nome: $('#nome').val().trim(),
                email: $('#email').val().trim(),
                nick: $('#nick').val().trim(),
                senha: $('#senha').val().trim()
            }
        }).done(function () { // 200, 201, 204
            Swal.fire("Sucesso!", "Usuário cadastrado com sucesso.", "succes")
                .then(function () {
                    $.ajax({
                        url: "/login",
                        method: "POST",
                        data: {
                            email: $('#email').val(),
                            senha: $('#senha').val()
                        }
                    }).done(function () {
                        window.location = '/home'
                    })
                        .fail(function (erro) { // 400 401 403 4094 500
                            Swal.fire("ERRO!", "Erro ao autenticar o usuário.", "error");
                        })
                })
        }).fail(function (erro) { // 400 401 403 4094 500
            Swal.fire("ERRO!", "Erro ao cadastrar o usuário.", "error");
        });
    }
}