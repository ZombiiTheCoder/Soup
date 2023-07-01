#!/bin/sh
clear
set -ex
output="main.out"
flags="-std=c11 -Wall -Wextra -Wpedantic -Werror -g"
gcc main.c $flags -o $output