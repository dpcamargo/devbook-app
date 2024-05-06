$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senhas não coincidem");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function(){
        alert("Usuario cadastrado com sucesso");
    }).fail(function(erro) {
        console.log(erro);
        alert("Erro ao cadastrar usuário");

    });
}