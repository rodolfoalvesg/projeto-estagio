# projeto-estagio
Aplicação em go - Projeto de aprendizado durante estágio na Stone.

# Desafio Técnico - Estágio/Junior Go(lang)

## Requisitos Técnicos

- O código do desafio *deve* estar na linguagem [Go](https://golang.org/)
- Pode-se utilizar qualquer package ou framework, mas lembre-se: stdlib > packages > frameworks (usar standard library do Go é melhor que usar packages, que é melhor que usar frameworks...)

_Indicação de material de estudo no final_

## O que será avaliado?

- Aprendizado do basico da linguagem Go, seguindo os padrões da linguagem (https://go.dev/doc/effective_go)
- Lógica utilizada para resolver o problema
- Resolução atende os casos de uso? e os corner cases?
- Foram criados testes para assegurar os casos de uso?

## Como devo entregar o teste?

Você tem um prazo de **2 semanas** para realizar o teste. Quando finalizar, envie um e-mail com o arquivo da sua solução ou o link para o repositório. Você receberá um e-mail com a confirmação de recebimento!

Caso sua solução não caiba em um único arquivo, não se preocupe. Você também pode gerar um arquivo .zip (comprimido)!

## Vamos ao teste!

Imagine uma lista de compras. Ela possui:

- Itens
- Quantidade de cada item
- Preço por unidade/peso/pacote de cada item

Desenvolva uma função (ou método) que irá receber uma lista de compras (como a detalhada acima) e uma lista de e-mails. Aqui, cada e-mail representa uma pessoa. 

A função deve:

- Calcular a soma dos valores, ou seja, multiplicar o preço de cada item por sua quantidade e somar todos os itens
- Dividir o valor de forma igual entre a quantidade de e-mails
- Retornar um mapa/dicionário onde a chave será o e-mail e o valor será quanto ele deve pagar nessa conta

**⚠️ IMPORTANTE**

- Quando fizer a divisão, é **importante que não falte nenhum centavo!** Cuidado quando tiver, por exemplo, um valor total de R$ 1,00 para ser dividido entre 3 e-mails. Isso daria uma dízima infinita. No entanto, o correto aqui é que duas pessoas fiquem com o valor 0,33 e a terceira fique com 0,34.
- Para facilitar e como boa prática, **não trabalhe com valores com decimais**. Ou seja, ponto flutuante ou float. Use inteiros para representar os valores! Como a menor unidade na nossa moeda é o centavo, use valores inteiros em centavos. Assim, um real será representado por 100 (cem centavos é igual a um real).

**DICAS**

- Para a avaliação, iremos gerar listas de itens de acordo com a função/método definida por você e listas de e-mails aleatórias. Portanto, defina bem o comportamento para, por exemplo, uma lista vazia de itens ou uma lista vazia de e-mails.
- Não vamos trabalhar com valores nulos. Você pode assumir que os valores de entrada **SEMPRE** estarão preenchidos, porém, quando são listas, podem estar vazias.
- Não é necessário escrever o desafio em inglês, mas é algo que você pode fazer caso se sinta confortável com seu inglês técnico.
- Dedique-se ao desafio com calma. Arrume um espaço e um horário que você possa se concentrar e não hesite em procurar dicas na internet, mas tome cuidado com copiar respostas prontas para problemas. Entenda o que está fazendo e verifique muitas vezes com dados diferentes a solução. Estresse bem o modelo com listas grandes de dados e coisas do tipo para ganhar confiança na sua solução.
- Caso ache necessário, você pode incluir um arquivo de documentação a parte do código. Um "README.md" por exemplo onde você explica como chegou nessa solução, como executá-la e etc.

## Mensagem final

Divirta-se no processo! Aproveite esse pequeno teste como uma prática divertida de problemas de programação. Esse pode ser um pontapé inicial para acontecimentos maiores na sua carreira! ;D

### Material de Estudo

#### Go
- [Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes/)
- [Curso Aprenda Go (@ellenkorbes) ](https://www.youtube.com/channel/UCxD5EE0H7qOhRr0tIVsOZPQ)
- [Learn Go](https://learn.go.dev)
- [Gophercise](https://gophercises.com/)
- [Error Handling](https://rauljordan.com/2020/07/06/why-go-error-handling-is-awesome.html)

#### Boas praticas
- [Boas Práticas na Stone](https://github.com/stone-payments/stoneco-best-practices/blob/master/README_pt.md)
- [Uber-go Guide](https://github.com/uber-go/guide/blob/master/style.md)

#### Outros
- [Just for Func](https://www.youtube.com/playlist?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ)
- [Grupo de Estudos de Go (pt-br)](https://www.youtube.com/channel/UCxRoRvJi7NbC2boKAV70t_g)
- [Go Bootcamp from Gopherguides.tv](https://www.youtube.com/watch?v=22R1PqXvtws)
- [Go WEB Examples](https://gowebexamples.com/)
- [Dave Cheney Blog](https://dave.cheney.net/practical-go)
- [Ardan Labs Blog](https://www.ardanlabs.com/blog)

#### Comunidade Go

- [Slack Gopher](https://invite.slack.golangbridge.org/)
- [Telegram](https://t.me/go_br)
