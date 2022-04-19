### Running
```
 go run . --expr "34+5+3/5-4*(4+545)/3"
```


### Grammar

This is LL(1) grammar
* `FACTOR ⟶ number | ( EXPR )`
* `TERM ⟶ FACTOR TERM1`
* `TERM1 ⟶ * FACTOR TERM1 | / FACTOR TERM1 | epsilon`
* `EXPR ⟶ TERM EXPR1`
* `EXPR1 ⟶ + TERM EXPR1 | - TERM EXPR1 | epsilon`


###  First and Follow
* `FIRST(FACTOR) = { number, ( }`
* `FIRST(TERM) = { number, ( }`
* `FIRST(EXPR) = { number, ( }`
* `FIRST(EXPR1) = { +, -, epsilon }`
* `FIRST(TERM1) = { *, /, epsilon }` 


* `FOLLOW(FACTOR) = { ), *, /, +, - }`
* `FOLLOW(TERM) = { ), +, - }`
* `FOLLOW(EXPR) = { ) }`
* `FOLLOW(EXPR1) = { ) }`
* `FOLLOW(TERM1) = { ), +, - }`



Computed by https://mikedevice.github.io/first-follow/

## Automata for lexical analyzer

![Lexical Analyzer Automata](https://github.com/DamirJann/math-parser/blob/master/img/automata_for_lexical_analyzer.drawio.png)

## Supported 
* Operations:
  * MULTIPLICATION - `*`
  * DIVISION - `/`
  * PLUS - `+`
  * MINUS - `-`
* Numbers:
  * Integer

