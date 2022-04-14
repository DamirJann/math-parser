## Syntaxis

* `digit → 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9`
* `factor → digit | (expr)`
* `term → term * factor | term / factor | factor`
* `expr → expr + term | expr - term | term`

## Automata for lexical analyzer

![Lexical Analyzer Automata](https://github.com/DamirJann/math-parser/blob/master/img/automata_for_lexical_analyzer.drawio.png)