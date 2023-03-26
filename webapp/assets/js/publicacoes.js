$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao);
$('.descurtir-publicacao').on('click', descurtirPublicacao);
$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

function atualizarEngajamento(data) {
    $('#conta-curtidas-' + data.id).text(data.curtidas);
    $('#conta-descurtidas-' + data.id).text(data.descurtidas);

    if (data.curtiu == 0) {
        $('#curtir-' + data.id).addClass('fa-regular');
        $('#curtir-' + data.id).removeClass('fa-solid');
    } else {
        $('#curtir-' + data.id).removeClass('fa-regular');
        $('#curtir-' + data.id).addClass('fa-solid');
    }

    if (data.descurtiu == 0) {
        $('#descurtir-' + data.id).addClass('fa-regular');
        $('#descurtir-' + data.id).removeClass('fa-solid');
    } else {
        $('#descurtir-' + data.id).removeClass('fa-regular');
        $('#descurtir-' + data.id).addClass('fa-solid');
    }
}

function criarPublicacao(evento) {
    evento.preventDefault();

    $.ajax({
        url: '/publicacoes',
        method: 'POST',
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function () {
        Swal.fire('Sucesso!', 'Publicação atualizada com sucesso!', 'success')
            .then(function () {
                window.location = "/home";
            })
    }).fail(function (erro) {
        Swal.fire("ERRO!", `Erro ao fazer Publicação:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
    })
}

function curtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/alternar-curtir`,
        method: 'POST'
    }).done(function (data) {
        atualizarEngajamento(data);
    }).fail(function (erro) {
        Swal.fire("ERRO!", "Erro ao curtir Publicação.", "error");
    }).always(function (erro) {
        elementoClicado.prop('disabled', false);
        Swal.fire("ERRO!", `Erro ao curtir Publicação:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
    });
}

function descurtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id')

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/alternar-descurtir`,
        method: 'POST'
    }).done(function (data) {
        atualizarEngajamento(data);
    }).fail(function (erro) {
        Swal.fire("ERRO!", `Erro ao descurtir Publicação:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });
}

function atualizarPublicacao(evento) {
    evento.preventDefault();

    $(this).prop('disabled', false)

    const elementoClicado = $(evento.target);
    const publicacaoId = $(this).data('publicacao-id')

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: 'PUT',
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function () {
        Swal.fire('Sucesso!', 'Publicação atualizada com sucesso!', 'success')
            .then(function () {
                window.location = "/home";
            })
    }).fail(function (erro) {
        Swal.fire("ERRO!", `Erro ao editar Publicação:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
    }).always(function () {
        $('#atualizar-publicacao').prop('disabled', false);
    })
}

function deletarPublicacao(evento) {
    evento.preventDefault();

    Swal.fire({
        title: "<b>Atenção!</b>",
        text: "Tem certeza que deseja excluir esta publicação? Esta ação é irreversível.",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function (confirmacao) {
        if (!confirmacao.value) {
            return;
        }

        const elementoClicado = $(evento.target);
        const publicacao = elementoClicado.closest('div');
        const publicacaoId = publicacao.data('publicacao-id');

        elementoClicado.prop('disabled', true);

        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: 'DELETE',
            data: {
                titulo: $('#titulo').val(),
                conteudo: $('#conteudo').val()
            }
        }).done(function () {
            publicacao.fadeOut("slow", function () {
                $(this).remove();
            })
        }).fail(function (erro) {
            Swal.fire("ERRO!", `Erro ao excluir Publicação:<br/><br/><b>${erro.responseJSON.erro}</b>`, "error");
        })
    });
}
