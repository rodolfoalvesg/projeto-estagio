## Aplicação em go

# Funcões e testes
1 - A aplicação recebe duas listas: 1 multdimensão com a LISTA DE PRODUTOS e 1 unidimensional com a LISTA DE EMAILS.<br />
2 - Na função principal existe a chamada da função com uma VALIDAÇÃO CONDICIONAL para caso a lista de email esteja VAZIA a função NÃO seja executada.<br />
3 - Caso a condição dois, seja satisfeita, a função é executada e novamente é executada uma VALIDAÇÃO CONDICIONAL verificando se existe email REPETIDO. Caso
haja, a função é encerrada retornando o MAP VAZIO com a mensangem de ERRO.<br />
4 - Caso a condição anterior seja FALSE, um loop é executado na LISTA DE PRODUTOS passando parametros para função CONVERT, que por sua vez realiza uma
SUB MULTIPLICAÇÃO da QTD x VALOR e retorna o valor para ITERAR a soma final.<br />
5 - Uma vez finalizada a ação anterior, é feito um novo LOOP baseado no tamanho da lista de email, adicionando os valores de CHAVES e seus respectivos valores.
6 -  E por fim é retornado o MAP de valores.
