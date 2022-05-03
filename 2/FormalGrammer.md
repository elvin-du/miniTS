```
//Fromal Grammer 正则文法
//词法文法

DecimalLiteral: IntegerLiteral '.' [0-9]* 
              | '.' [0-9]+
              | IntegerLiteral 
              ;
IntegerLiteral: '0' | [1-9] [0-9]* ;
StringLiteral: '"' * '"' ; 

Identifier: [a-zA-Z_][a-zA-Z0-9_]* ;

Keyword: func
        | return
        | int
        | string
        | float
        | if
        | else
        | for
        ;

//分隔符
Seperator: (
        | )
        | {
        | }
        | ;
        | ,
        ;

//运算符
Operator: +
        | +=
        | -
        | -=
        | *
        | *=
        | /
        | /=
        | >
        | >=
        | <
        | <=
        | ==
        ; 
        
        
```