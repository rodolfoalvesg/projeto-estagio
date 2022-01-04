# Versão 2 do desafio stone

## Pacotes
 - Lista de Valores: `list/`
 - Validação de entradas: `validate/`
 - Cálculo e distribuição de valores da conta: `calculator/`


### Lista contendo as entradas do sistema: `list/`
- O pacote `list/` contém uma (1) estrutura portando todos os itens da compra sendo motandos em:

```go
type ListItens struct {
	Item, Spec    string
	Amount, Price uint
}
```
e acessíveis a partir de: `BaseItens`.

- O pacote ainda contém uma lista de emails no slice: 
    `BaseEmails`


### Validador de entradas do sistema: `validate/`
- O pacote `validate/` recebe como entrada através do seu metódo a lista de emails, para analisar e verificar a existencia de emails duplicado ou vazio.
- Através do tipo:
```go 
type Emails []string
```
é invocado o metódo `Emails.Validate()` que executa a verificação da duplicidade de emails e ausência, retornando um valor boleano: `bool`

### Calculadora e distribuidora de contas: `calculator/`
- O pacote `calculator/` contém uma função chamada: `ListCalculator()`, que recebe como parâmetros: `BaseItens` e `BaseEmails` e retorna dois valores do tipo `int`.

    1 - `valuePeron`: Valor total inteiro sem o resto.
    
    2 - `remaining`: O resto da divisão.

- Após a realização dos cálculos da conta e do resto, é possível fazer a distribuíção dos valores para cada e-mail a instanciação do metódo `AddedValuesToUser()` de interface `ValueDistribuitor ` a partir das estruturas: `RemainingForMost` e `RemainingForOne`, onde:
    - Exemplo:
    ```go
    toMost := calculator.RemainingForMost{ValuePerPerson: valuePerPerson, Remaining: remaining, Email: list.BaseEmails}
    toOne := calculator.RemainingForOne{ValuePerPerson: valuePerPerson, Remaining: remaining, Email: list.BaseEmails}
    ```

    1 - `toMost.AddedValuesToUser()`: Responsável pela distribuição do resto da divisão para o maior número de e-mails possíveis e retorna um `map`
    
    2 - `toOne.AddedValuesToUser()` : Responsável pela distribuíção do resto da divisão para o último e-mail da lista e retorna um `map`

### Pacote principal: `main`
- Aqui importamos todos os demais pacotes criados e necessários para o funcionamento do programa.
- Fazemos a chamadas das validações necessárias.
- Após validado, criamos as instâncias necessárias para obtenção das listas.
- Em seguida é feito o mapeamento dos valores obtidos, para listar em tela.
