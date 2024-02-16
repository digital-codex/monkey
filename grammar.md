## Syntax Grammar § 1.1

The syntactic grammar is used to parse the linear sequence of tokens into the \
nested syntax tree structure. It starts with the first rule that matches an entire \
Monkey program (or a single REPL entry).

````
program     -> declaraction* EOF ;
````

### Declarations § 1.1.1

A program is a series of declaration, which are the statements that bind new \
identifiers or any of the other statements types.

````
declaration -> letDecl
             | statement ;

letDecl     -> "let" <IDENTIFIER> "=" expression ";"? ;
````

### Statements § 1.1.2

The remaining statement rules produce side effects, but do not introduce \
bindings.

````
statement   -> exprStmt
             | returnStmt ;

exprStmt    -> expression ";"? ;
returnStmt  -> "return" expression ";"? ;
````

### Expression § 1.1.3

Expression produce values. Monkey has a number of unary and binary operators \
with different levels of precedence. Some grammars for languages do not \
directly encode the precedence relationship and specify that elsewhere. Here, \
we use a separate rule for each precedence level to make it explicit.

````
expression  -> equality ;

equality    -> comparison ( ( "!=" | "==" ) comparison* )* ;
comparison  -> term ( ( ">" | "<" ) term )* ;
term        -> factor ( ( "-" | "+" ) factor )* ;
factor      -> unary ( ( "/" | "*" ) unary )* ;

unary       -> ( "!" | "-" ) unary | call ;
call        -> index ( "(" arguments? ")" )* ;
index       -> primary "[" expression "]" ;
primary     -> "true" 
             | "false" 
             | "null" 
             | NUMBER 
             | STRING 
             | IDENTIFIER
             | "(" expression ")"
             | if
             | function 
             | array
             | hash
             | macro ;
````

### Utility rules § 1.1.4

In order to keep the above rules a littler cleaner, some of the grammar is split out \
into a few reused helper rules.

````
block       -> "{" declaration* "}" ;
if          -> "if" "(" expression ")" block ( "else" block )? ;
function    -> "fn" "(" parameters? ")" block ;
array       -> "[" elements* "]" ;
hash        -> "{" pair* "}" ;
macro       -> "macro" "(" parameters? ")" block ;

parameters  -> IDENTIFIER ( "," IDENTIFIER )* ;
arguments   -> expression ( "," expression )* ;
elements    -> expression ( "," expression )* ;
pair        -> expression ":" expression ( "," expression ":" expression )* ;
````

## Lexical Grammar § 1.2

The lexical grammar is used by the lexer to group characters into tokens. \
Where the syntax is [context free](https://en.wikipedia.org/wiki/Context-free_grammar), the lexical grammar is [regular](https://en.wikipedia.org/wiki/Regular_grammar) - note that \
there are no recursive rules.

````
NUMBER      -> <DIGIT>+ ;
STRING      -> "\"" <CHARACTER>* "\"" ;
IDENTIFIER  -> <ALPHA>+ ;
ALPHA       -> "a" ... "z" | "A" ... "Z" | "_" ;
DIGIT       -> "0" ... "9" ;
CHARACTER   -> "\t" | "\n" | " " ... "~" ;
````