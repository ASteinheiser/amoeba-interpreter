#!/bin/bash
Purple='\033[0;35m'
BPurple='\033[1;35m'
Cyan='\033[0;36m' 
BCyan='\033[1;36m'
BlackBG='\033[40m'
NoColor='\033[0m'

echo -e "${BlackBG}                                      ${NoColor}"
echo -e "${BlackBG}${BCyan}    Running the Amoeba test suite!    ${NoColor}"
echo -e "${BlackBG}                                      ${NoColor}"
echo -e "${BlackBG}${BPurple}              ¯\_(ツ)_/¯              ${NoColor}"
echo -e "${BlackBG}                                      ${NoColor}"
echo ""

echo -e "${BlackBG}${BCyan}AST Test Results:${NoColor}"
go test ./ast/
echo ""

echo -e "${BlackBG}${BCyan}Lexer Test Results:${NoColor}"
go test ./lexer/
echo ""

echo -e "${BlackBG}${BCyan}Parser Test Results:${NoColor}"
go test ./parser/
echo ""

echo -e "${BlackBG}${BCyan}Evaluator Test Results:${NoColor}"
go test ./evaluator/
echo ""
