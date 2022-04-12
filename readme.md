## Syntactic

* `digit → 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9`
* `factor → digit | (expr)`
* `term → term * factor | term / factor | factor`
* `expr → expr + term | expr - term | term`