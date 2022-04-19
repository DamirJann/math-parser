### Running
```
 $ go run . --expr "34+5+3/5-4*(4+545)/3"
2022/04/19 16:28:30.584051 tokens len=17: [0xc000004090 0xc0000040a8 0xc0000040c0 0xc0000040d8 0xc0000040f0 0xc000004108 0xc000004120 0xc000004138 0xc000004150 0xc000004168 0xc
000004180 0xc000004198 0xc0000041b0 0xc0000041c8 0xc0000041e0 0xc0000041f8 0xc000004210]
2022/04/19 16:28:30.603228 FACTOR -> number(34) parsed
2022/04/19 16:28:30.603311 TERM1 -> EPSILON parsed%!(EXTRA int=34)
2022/04/19 16:28:30.603311 TERM -> FACTOR TERM1 parsed
2022/04/19 16:28:30.603917 FACTOR -> number(5) parsed
2022/04/19 16:28:30.603917 TERM1 -> EPSILON parsed%!(EXTRA int=5)
2022/04/19 16:28:30.603917 TERM -> FACTOR TERM1 parsed
2022/04/19 16:28:30.604450 FACTOR -> number(3) parsed
2022/04/19 16:28:30.604450 FACTOR -> number(5) parsed
2022/04/19 16:28:30.604450 TERM1 -> EPSILON parsed%!(EXTRA int=5)
2022/04/19 16:28:30.604979 TERM1 -> / TERM TERM1 parsed%!(EXTRA int=5)
2022/04/19 16:28:30.604979 TERM -> FACTOR TERM1 parsed
2022/04/19 16:28:30.604979 FACTOR -> number(4) parsed
2022/04/19 16:28:30.605555 FACTOR -> number(4) parsed
2022/04/19 16:28:30.605555 TERM1 -> EPSILON parsed%!(EXTRA int=4)
2022/04/19 16:28:30.605555 TERM -> FACTOR TERM1 parsed
2022/04/19 16:28:30.605555 FACTOR -> number(545) parsed
2022/04/19 16:28:30.606203 TERM1 -> EPSILON parsed%!(EXTRA int=545)
2022/04/19 16:28:30.606203 TERM -> FACTOR TERM1 parsed
2022/04/19 16:28:30.606203 EXPRESSION1 -> EPSILON parsed%!(EXTRA int=545)
2022/04/19 16:28:30.606203 EXPR1 -> + TERM EXPR1 parsed%!(EXTRA int=545)
2022/04/19 16:28:30.606203 EXP -> TERM EXPR1 parsed
2022/04/19 16:28:30.606203 FACTOR -> ( EXPR ) parsed
2022/04/19 16:28:30.606203 FACTOR -> number(3) parsed
2022/04/19 16:28:30.606710 TERM1 -> EPSILON parsed%!(EXTRA int=3)
2022/04/19 16:28:30.606733 TERM1 -> / TERM TERM1 parsed%!(EXTRA int=3)
2022/04/19 16:28:30.606733 TERM1 -> * TERM TERM1 parsed%!(EXTRA int=3)
2022/04/19 16:28:30.606733 TERM -> FACTOR TERM1 parsed
2022/04/19 16:28:30.606733 EXPRESSION1 -> EPSILON parsed%!(EXTRA int=3)
2022/04/19 16:28:30.607240 EXPR1 -> - TERM EXPR1 parsed%!(EXTRA int=3)
2022/04/19 16:28:30.607283 EXPR1 -> + TERM EXPR1 parsed%!(EXTRA int=3)
2022/04/19 16:28:30.607283 EXPR1 -> + TERM EXPR1 parsed%!(EXTRA int=3)
2022/04/19 16:28:30.607283 EXP -> TERM EXPR1 parsed
2022/04/19 16:28:30.607886 computed ast:
EXPR
├───TERM
│       ├───FACTOR
│       │       └───34
│       └───TERM1
│               └───eps
└───EXPR1
        ├───+
        ├───TERM
        │       ├───FACTOR
        │       │       └───5
        │       └───TERM1
        │               └───eps
        └───EXPR1
                ├───+
                ├───TERM
                │       ├───FACTOR
                │       │       └───3
                │       └───TERM1
                │               ├───/
                │               ├───FACTOR
                │               │       └───5
                │               └───TERM1
                │                       └───eps
                └───EXPR1
                        ├───-
                        ├───TERM
                        │       ├───FACTOR
                        │       │       └───4
                        │       └───TERM1
                        │               ├───*
                        │               ├───FACTOR
                        │               │       ├───(
                        │               │       ├───EXPR
                        │               │       │       ├───TERM
                        │               │       │       │       ├───FACTOR
                        │               │       │       │       │       └───4
                        │               │       │       │       └───TERM1
                        │               │       │       │               └───eps
                        │               │       │       └───EXPR1
                        │               │       │               ├───+
                        │               │       │               ├───TERM
                        │               │       │               │       ├───FACTOR
                        │               │       │               │       │       └───545
                        │               │       │               │       └───TERM1
                        │               │       │               │               └───eps
                        │               │       │               └───EXPR1
                        │               │       │                       └───eps
                        │               │       └───)
                        │               └───TERM1
                        │                       ├───/
                        │                       ├───FACTOR
                        │                       │       └───3
                        │                       └───TERM1
                        │                               └───eps
                        └───EXPR1
                                └───eps
2022/04/19 16:28:30.609670 simplified ast:
2022/04/19 16:28:30.610242 postfix notation: %v [{1 34} {1 5} {2 +} {1 3} {1 5} {5 /} {2 +} {1 4} {1 4} {1 545} {2 +} {4 *} {1 3} {5 /} {3 -}]
Evaluation of `34+5+3/5-4*(4+545)/3` is equal to -693

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

