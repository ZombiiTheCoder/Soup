#include "Includes.h"

int main(int argc, char **argv)
{
    assertExpr(argc == 2);
    char *src = fileReadText(argv[1]);
    printf("src:\n%s\n", src);

    Token *tokens = tokenize(src);
    tokenPrintAll(tokens);
    tokenFreeAll(tokens);
    return 0;
}

