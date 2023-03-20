$('#parar-de-seguir').on('click', pararDeSeguir)
$('#seguir').on('click', seguir)
$('#editar-usuario').on('submit', editar)
$('#atualizar-senha').on('submit', atualizarSenha)
$('#deletar-usuario').on('click', deletarUsuario)

function pararDeSeguir() {
    console.log('pararDeSeguir()')
    const usuarioId = $(this).data('usuario-id');

    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done(function () {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function () {
        Swal.fire("ERRO!", "Erro ao parar de seguir o usuário.", "error");
        $('#parar-de-seguir').this.prop('disabled', false);
    })
}

function seguir() {
    console.log('seguir()')
    const usuarioId = $(this).data('usuario-id');

    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST"
    }).done(function () {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function () {
        Swal.fire("ERRO!", "Erro ao seguir o usuário.", "error");
        $('#seguir').this.prop('disabled', false);
    })
}

function editar(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            nome: $('#nome').val().trim(),
            email: $('#email').val().trim(),
            nick: $('#nick').val().trim(),
        }
    }).done(function () {
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso.", "success")
            .then(function () {
                window.location = "/perfil"
            });
    }).fail(function () {
        Swal.fire("ERRO!", "Erro ao atualizar o usuário.", "error")
    });
}

function atualizarSenha(evento) {
    evento.preventDefault();

    if ($('#nova-senha').val().trim() != $('#confirmar-senha').val().trim()) {
        Swal.fire("ERRO!", "As senhas não coincidem", "warning");

        $('#nova-senha').focus();
        return;
    }

    $.ajax({
        url: "/atualizar-senha",
        method: "POST",
        data: {
            atual: $('#senha-atual').val().trim(),
            nova: $('#nova-senha').val().trim()
        }
    }).done(function () {
        Swal.fire("Sucesso!", "A senha foi atualizada com sucesso.", "success")
            .then(function () {
                window.locate = "/perfil";
            });
    }).fail(function () {
        Swal.fire("ERRO!", "Erro ao alterar senha.", "error");
    })
}

function deletarUsuario(evento) {
    Swal.fire({
        title: 'Atenção!',
        html: 'Tem certeza que deseja apagar a sua conta?<br /><br /><h4><b>Essa é uma ação irreversível.</b></h4>',
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning",
    }).then(function (confirmacao) {
        if (confirmacao.value) {
            $.ajax({
                url: '/deletar-usuario',
                method: "DELETE"
            }).done(function () {
                Swal.fire('Sucesso!', 'Sua conta foi excluída com sucesso!', 'success')
                    .then(function () {
                        window.location = '/logout';
                    })
            }).fail(function () {
                Swal.fire("ERRO!", "ocorreu um erro ao excluir a conta", "error")
            });
        };
    });
}
