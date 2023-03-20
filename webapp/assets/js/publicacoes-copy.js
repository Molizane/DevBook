$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao);
$('.descurtir-publicacao').on('click', descurtirPublicacao);
$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

function atualizarEngajamento(elementoClicado, contador, tipo, data) {

    if (tipo == 'curtidas') {
        contador.text(data.curtidas);

        if (data.curtiu == 0) {
            elementoClicado.addClass('fa-regular');
            elementoClicado.removeClass('fa-solid');
            return
        }

        elementoClicado.removeClass('fa-regular');
        elementoClicado.addClass('fa-solid');

        return;
    }

    contador.text(data.descurtidas);

    if (data.descurtiu == 0) {
        elementoClicado.addClass('fa-regular');
        elementoClicado.removeClass('fa-solid');
        return
    }

    elementoClicado.removeClass('fa-regular');
    elementoClicado.addClass('fa-solid');
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
    }).fail(function () {
        Swal.fire("ERRO!", "Erro ao criar a Publicação.", "error");
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
        const contadorDeCurtidas = elementoClicado.next('span');
        atualizarEngajamento(elementoClicado, contadorDeCurtidas, 'curtidas', data);
    }).fail(function (erro) {
        Swal.fire("ERRO!", "Erro ao curtir Publicação.", "error");
    }).always(function () {
        elementoClicado.prop('disabled', false);
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
        console.log(data);
        const contadorDeDescurtidas = elementoClicado.next('span');
        atualizarEngajamento(elementoClicado, contadorDeDescurtidas, 'descurtidas', data);
    }).fail(function () {
        Swal.fire("ERRO!", "Erro ao descurtir Publicação.", "error");
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
    }).fail(function () {
        Swal.fire("ERRO!", "Erro ao editar Publicação.", "error");
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
        }).fail(function () {
            Swal.fire("ERRO!", "Erro ao apagar Publicação.", "error");
        })
    });

}
