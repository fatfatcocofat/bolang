#!/bin/sh

old_pwd=$(pwd)

grammar="$old_pwd/grammar"
output="$old_pwd/parser"
antlr4="java -jar $old_pwd/tools/antlr-4.13.1-complete.jar"

cd $grammar

$antlr4 -Dlanguage=Go -o $output -visitor -package parser -no-listener *.g4