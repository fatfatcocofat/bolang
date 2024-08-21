grammar Bo;

program
    : statement* EOF
    ;

statement
    : requireStatement
    | variableDeclaration
    | functionCall
    ;

expression
    : INT | FLOAT | STRING | BOOL | ID
    ;

functionParameters
    : LPAREN (expression (COMMA expression)*)? RPAREN
    ;

functionCall
    : ID functionParameters // foo(1, 2, 3);
    | expression PERIOD ID functionParameters // foo.bar(1, 2, 3); | "foo".bar(1, 2, 3);
    ;

variableDeclaration
    : typeSpec ID ASSIGN expression // int a = 1;
    ;

typeSpec
    : 'int'
    | 'float'
    | 'string'
    | 'bool'
    ;

requireStatement
    : REQUIRE importPath
    ;

importPath
    : LT ID ('/' ID)* GT
    | STRING
    ;

LT              : '<';
GT              : '>';
ASSIGN          : '=';

LPAREN          : '(';
RPAREN          : ')';
LBRACE          : '{';
RBRACE          : '}';
PERIOD          : '.';
COMMA           : ',';
SEMICOLON       : ';';

REQUIRE         : 'require';

INT             : [0-9]+;
FLOAT           : [0-9]+ '.' [0-9]*;
BOOL            : 'true' | 'false';
STRING          : '"' (ESC | ~["\\])* '"'
                | '\'' (ESC | ~['\\])* '\''
                ;

ID              : [a-zA-Z_][a-zA-Z0-9_]*;
WS              : [ \t\r\n]+ -> skip;

// Comments
S_COMMENT       : '//' ~[\r\n]* '\r'? '\n' -> channel(HIDDEN);
M_COMMENT       : '/*' .*? '*/' -> channel(HIDDEN);

fragment ESC    : '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX    : [0-9a-fA-F];
fragment DIGIT  : INT | FLOAT;