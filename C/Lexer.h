#ifndef LEXER_H
#define LEXER_H

typedef enum {
    T_OPENPAREN,
    T_CLOSEDPAREN,
    T_OPENBRACE,
    T_CLOSEDBRACE,
    T_OPENBRACKET,
    T_CLOSEDBRACKET,
    T_COLON,
    T_SEMICOLON,
    T_COMMA,
    T_DOT,
    T_SLASH,
    T_PERCENT,
    T_QUESTIONMARK,
    T_LESSTHEN,
    T_LESSTHENEQUAL,
    T_GREATERTHEN,
    T_GREATERTHENEQUAL,
    T_OR,
    T_PIPE,
    T_AND,
    T_AMPERSAN,
    T_STAR,
    T_BANG,
    T_CARROT,
    T_NEQUALS,
    T_DEQUALS,
    T_EQUALS,
    T_PLUSEQUALS,
    T_DASHEQUALS,
    T_SLASHEQUALS,
    T_MODULOEQUALS,
    T_STAREQUALS,
    T_PLUS,
    T_DPLUS,
    T_MINUS,
    T_DMINUS,
    T_IF,
    T_VAR,
    T_USE,
    T_ELSE,
    T_CONST,
    T_WHILE,
    T_RETURN,
    T_FUNCTION,
    T_INT,
    T_FLOAT,
    T_BOOLEAN,
    T_STRING,
    T_CHAR,
    T_IDENTIFIER,
    T_TYPE,
    T_EOF,
    T_UNKNOWN,
    T_N
} TokenType;

char *TokenTypeStr[T_N] = {
    "(",
    ")",
    "{",
    "}",
    "[",
    "]",
    ":",
    ";",
    ",",
    ".",
    "/",
    "%",
    "?",
    "<",
    "<=",
    ">",
    ">=",
    "||",
    "|",
    "&&",
    "&",
    "*",
    "!",
    "^",
    "!=",
    "==",
    "=",
    "+=",
    "-=",
    "/=",
    "%=",
    "*=",
    "+",
    "++",
    "-",
    "--",
    "If",
    "Var",
    "Import",
    "Else",
    "Const",
    "While",
    "Return",
    "Function",
    "Int",
    "Float",
    "Boolean",
    "String",
    "Char",
    "Identifier",
    "Type",
    "EOF",
    "T_UNKNOWN"
};

bool strcmpn(char *a, char *b, const size_t len) {
    for(size_t i = 0; i < len; i++){
        if(a[i] == '\0' || b[i] == '\0' || a[i] != b[i])
            return false;
    }
    return true;
}

char* strdup(char *str) {
    const size_t len = strlen(str);
    char *ret = malloc(len+1);
    memcpy(ret, str, len);
    ret[len] = '\0';
    return ret;
}

char* strndup(char *str, const size_t len){
    char *ret = malloc(len + 1);
    memcpy(ret, str, len);
    ret[len] = '\0';
    return ret;
}

TokenType tokenStrToType(char *str) {
    for(size_t i = 0; i < T_N; i++){
        if (strcmpn(str, TokenTypeStr[i], strlen(TokenTypeStr[i])))
            return i;
    }
    return T_UNKNOWN;
}

typedef struct Token{
    TokenType type;
    char *value;
    size_t line;
    size_t column;
    struct Token *next;
}Token;

Token* tokenNew(size_t line, size_t column) {
    Token *token = calloc(1, sizeof(Token));
    token->type = T_UNKNOWN;
    token->value = NULL;
    token->line = line;
    token->column = column;
    return token;
}

Token* tokenAppend(Token *head, Token *tail) {
    if(!head)
        return tail;
    Token *cur = head;
    while(cur->next)
        cur = cur->next;
    cur->next = tail;
    return head;
}

typedef struct{
    size_t line;
    size_t column;
    char *start;
    char *pos;
    Token *tokens;
}Lexer;

void tokenPrint(Token *token) {
    printf("type: %s\n", TokenTypeStr[token->type]);
    printf("value: %s\n", token->value);
    printf("line: %zu\n", token->line);
    printf("column: %zu\n\n", token->column);
}

void tokenPrintAll(Token *tokens) {
    while(tokens){
        tokenPrint(tokens);
        tokens = tokens->next;
    }
}

Token* tokenFree(Token *token) {
    if(!token)
        return NULL;
    Token *next = token->next;
    free(token);
    return next;
}

void tokenFreeAll(Token *tokens) {
    while(tokens)
        tokens = tokenFree(tokens);
}

bool canSkip(char *pos){
    return *pos == ' ' || *pos == '\n' || *pos == '\r' || *pos == '\b' || *pos == '\f' || *pos == '\t' || *pos == '\v';
}

Lexer lexAdvance(Lexer lex, char *value){
    if(*lex.pos == '\n'){
        lex.line++;
        lex.column = 1;
        lex.pos++;
        return lex;
    }
    if(value!=NULL) {
        lex.column+=strlen(value);
        lex.pos+=strlen(value);
        return lex;
    }
    lex.pos++;
    lex.column++;
    return lex;
}

size_t digitLen(char *str) {
    size_t len = 0;
    while(isdigit(*str)){
        len++;
        str++;
    }
    return len;
}

Token* tokenInferType(Lexer lex) {
    Token *token = tokenNew(lex.line, lex.column);
	switch (*(lex.pos)) {
	case '(':
        token->type = T_OPENPAREN;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case ')':
        token->type = T_CLOSEDPAREN;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '{':
        token->type = T_OPENBRACE;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '}':
        token->type = T_CLOSEDBRACE ;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '[':
        token->type = T_OPENBRACKET;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case ']':
        token->type = T_CLOSEDBRACKET;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case ':':
        token->type = T_COLON;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case ';':
        token->type = T_SEMICOLON;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '.':
        token->type = T_DOT;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
    case ',':
        token->type = T_COMMA;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '?':
        token->type = T_QUESTIONMARK;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '^':
        token->type = T_CARROT;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '=':
		if (*(lex.pos+1) == '=' ){
			token->type = T_DEQUALS;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_EQUALS;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '/': 
		if (*(lex.pos+1) == '=') {
			token->type = T_SLASHEQUALS;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_SLASH;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '%':
		if (*(lex.pos+1) == '=') {
			token->type = T_MODULOEQUALS;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_PERCENT;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '+':
		if (*(lex.pos+1) == '+') {
			token->type = T_DPLUS;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		} else if (*(lex.pos+1) == '=') {
			token->type = T_PLUSEQUALS;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_PLUS;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '-':
		if (*(lex.pos+1) == '-') {
			token->type = T_DMINUS;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}else if (*(lex.pos+1) == '=') {
			token->type = T_DASHEQUALS;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_MINUS;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '*':
		if (*(lex.pos+1) == '=' ){
			token->type = T_STAREQUALS;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_STAR;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '>':
		if (*(lex.pos+1) == '=') {
			token->type = T_GREATERTHENEQUAL;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_GREATERTHEN;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '<':
		if (*(lex.pos+1) == '=') {
			token->type = T_LESSTHENEQUAL;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_LESSTHEN;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '!':
		if (*(lex.pos+1) == '=') {
			token->type = T_NEQUALS;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_BANG;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	case '&':
		if (*(lex.pos+1) == '&') {
			token->type = T_AND;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_AMPERSAN;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
		
	case '|':
		if (*(lex.pos+1) == '|') {
			token->type = T_OR;
            token->value = strdup(TokenTypeStr[token->type]);
            return token;
		}
		token->type = T_PIPE;
        token->value = strdup(TokenTypeStr[token->type]);
        return token;
	}
    if(isdigit(*(lex.pos))){
        token->type = T_INT;
        token->value = strndup(lex.pos, digitLen(lex.pos));
        return token;
    }
    return token;
}

Token* tokenize(char *text) {
    Lexer lex = {.line = 1, .column = 1, .start = text, .pos = text};
    
    
    while(*(lex.pos) != '\0'){
        while(canSkip(lex.pos)){
            lex = lexAdvance(lex, NULL);
        }
        Token *token = tokenInferType(lex)
        if(token->type != T_UNKNOWN){
            lex.tokens = tokenAppend(lex.tokens, token);
            lex = lexAdvance(lex, token->value);
            
        }
        
    }
    
    return tokenAppend(lex.tokens, tokenNew(T_EOF, NULL, lex.line, lex.column));
}

#endif /* end of include guard: LEXER_H */