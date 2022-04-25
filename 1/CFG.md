```
//Context-free Grammer 上下文无关文法 CFG
//语法文法
//EBNF 格式（扩展巴科斯范式）

prog = (functionDecl | functionCall)* ;
functionDecl: "function" Identifier "(" ")"  functionBody; 
functionBody : '{' functionCall* '}' ;
functionCall : Identifier '(' parameterList? ')' ;
parameterList : StringLiteral (',' StringLiteral)* ;

```