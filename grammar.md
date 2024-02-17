## Syntax Grammar § 1.1

The syntactic grammar is used to parse the linear sequence of tokens into the \
nested syntax tree structure. It starts with the first rule that matches an entire \
Monkey program (or a single REPL entry).

````
program     -> declaraction* <EOF> ;
````

### Declarations § 1.1.1

A program is a series of declaration, which are the statements that bind new \
identifiers or any of the other statements types.

````
declaration -> letDecl
             | statement ;

letDecl     -> <LET> <IDENT> <EQUAL> expression <SEMICOLON>? ;
````

### Statements § 1.1.2

The remaining statement rules produce side effects, but do not introduce \
bindings.

````
statement   -> returnStmt
             | exprStmt ;

returnStmt  -> <RETURN> expression <SEMICOLON>? ;
exprStmt    -> expression <SEMICOLON>? ;
````

### Expression § 1.1.3

Expression produce values. Monkey has a number of unary and binary operators \
with different levels of precedence. Some grammars for languages do not \
directly encode the precedence relationship and specify that elsewhere. Here, \
we use a separate rule for each precedence level to make it explicit.

````
expression  -> equality ;

equality    -> comparison ( ( <EQUAL_EQUAL> | <BANG_EQUAL> ) comparison* )* ;
comparison  -> term ( ( <LESS> | <MORE> ) term )* ;
term        -> factor ( ( <PLUS> | <MINUS> ) factor )* ;
factor      -> unary ( ( <STAR> | <SLASH> ) unary )* ;

unary       -> ( <BANG> | <MINUS> ) unary | call ;
call        -> index ( <LPAREN> expressions? <RPAREN> )* ;
index       -> primary <LBRACKET> expression <RBRACKET> ;
primary     -> <IDENT>
             | <NUMBER>
             | <LPAREN> expression <RPAREN>
             | "true" 
             | "false" 
             | if
             | function 
             | <STRING>
             | array
             | hash
             | macro ;
````

### Utility rules § 1.1.4

In order to keep the above rules a littler cleaner, some of the grammar is split out \
into a few reused helper rules.

````
block       -> <LBRACE> declaration* <RBRACE> ;

if          -> <IF> <LPAREN> expression <RPAREN> block ( <ELSE> block )? ;
function    -> <FN> <LPAREN> parameters? <RPAREN> block ;
array       -> <LBRACKET> expressions* <RBRACKET> ;
hash        -> <LBRACE> (expression <COLON> expression ( <COMMA> expression <COLON> expression )* )* <RBRACE> ;
macro       -> <MACRO> <LPAREN> parameters? <RPAREN> block ;

parameters  -> <IDENT> ( <COMMA> <IDENT> )* ;
expressions -> expression ( <COMMA> expression )* ;
````

## Lexical Grammar § 1.2

The lexical grammar is used by the lexer to group characters into tokens. \
Where the syntax is [context free](https://en.wikipedia.org/wiki/Context-free_grammar), the lexical grammar is [regular](https://en.wikipedia.org/wiki/Regular_grammar) - note that \
there are no recursive rules.

````
EQUAL       -> "=" ;
EQUAL_EQUAL -> "==" ;
BANG        -> "!" ;
BANG_EQUAL  -> "!=" ;

PLUS        -> "+" ;
MINUS       -> "-" ;
STAR        -> "*" ;
SLASH       -> "/" ;

LESS        -> "<" ;
MORE        -> ">" ;

COMMA       -> "," ;
COLON       -> ":" ;
SEMICOLON   -> ";" ;

LPAREN       -> "(" ;
LBRACE       -> "{" ;
LBRACKET     -> "[" ;
RPAREN       -> ")" ;
RBRACE       -> "}" ;
RBRACKET     -> "]" ;

STRING      -> "\"" <CHARACTER>* "\"" ;
IDENT       -> <ALPHA>+ ;
NUMBER      -> <DIGIT>+ ;

FN          -> "fn" ;
LET         -> "let" ;
TRUE        -> "true" ;
FALSE       -> "false" ;
IF          -> "if" ;
ELSE        -> "else" ;
RETURN      -> "return" ;
MACRO       -> "macro" ;

WHITESPACE  -> " " | "\t" | "\n" | "\r" ;
ALPHA       -> "a" ... "z" | "A" ... "Z" | "_" ;
DIGIT       -> "0" ... "9" ;
CHARACTER   -> "\t" | " " ... "~" ;

EOF         -> "" ;
````