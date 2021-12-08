## Aplicação em go

# Funcões e testes
0 - Aplicação dividida em módulos pacotes e importados no main. O inicio de tudo, começa com uma
validação da lista de email verificando se atende os requisitos.<br />
1 - A aplicação recebe duas listas do pacote LIST.<br />
    - items, emails := list.List()<br />
2 - Em seguida é chamado do pacote calculator o metódo: calculator.ListCalculator(items, emails) <br />
3 - Quando passada as listas: items(array de structs) e emails (um slice) a função listCalculator() realiza o cálculo de soma, em seguida passando parâmetros para a função AddedValuesUser() responsável pela distribuíção dos valores no map.

4 - Por fim é retornado o map contendo as chaves e seus respectivos valores.
