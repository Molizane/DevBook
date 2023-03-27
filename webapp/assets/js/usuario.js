$('#parar-de-seguir').on('click', pararDeSeguir)
$('#seguir').on('click', seguir)
$('#editar-usuario').on('submit', editar)
$('#atualizar-senha').on('submit', atualizarSenha)
$('#deletar-usuario').on('click', deletarUsuario)
$('#bloquear-usuario').on('click', bloquearUsuario)
$('#desbloquear-usuario').on('click', desbloquearUsuario)

function pararDeSeguir() {
    const usuarioId = $(this).data('usuario-id');

    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done(function () {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function (erro) {
        Swal.fire("ERRO!", `Erro ao parar de seguir o usuário:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
        $('#parar-de-seguir').this.prop('disabled', false);
    });
}

function seguir() {
    const usuarioId = $(this).data('usuario-id');

    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST"
    }).done(function () {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function (erro) {
        Swal.fire("ERRO!", `Erro ao seguir o usuário:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
        $('#seguir').this.prop('disabled', false);
    });
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
    }).fail(function (erro) {
        Swal.fire("ERRO!", `Erro ao atualizar o usuário:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
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
    }).fail(function (erro) {
        Swal.fire("ERRO!", `Erro ao alterar senha:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
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
            }).fail(function (erro) {
                Swal.fire("ERRO!", `Erro ao excluir a conta:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
            });
        };
    });
}

function bloquearUsuario(evento) {
    const usuarioId = $(this).data('usuario-id');

    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/bloquear`,
        method: "POST"
    }).done(function () {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function (erro) {
        Swal.fire("ERRO!", `Erro ao bloquear o usuário:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
        $('#parar-de-seguir').this.prop('disabled', false);
    });
}

function desbloquearUsuario(evento) {
    const usuarioId = $(this).data('usuario-id');

    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/desbloquear`,
        method: "POST"
    }).done(function () {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function (erro) {
        Swal.fire("ERRO!", `Erro ao desbloquear o usuário:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
        $('#parar-de-seguir').this.prop('disabled', false);
    });
}
