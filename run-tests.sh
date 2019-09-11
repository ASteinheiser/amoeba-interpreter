#!/bin/bash
Purple='\033[0;35m'
BPurple='\033[1;35m'
Cyan='\033[0;36m' 
BCyan='\033[1;36m'
NoColor='\033[0m'

echo -e "${BCyan}Running the Amoeba test suite!${NoColor}"
echo ""
echo -e "${BPurple}¯\_(ツ)_/¯${NoColor}"
echo ""

echo -e "${Cyan}Lexer Test Results:${NoColor}"
go test ./lexer/
echo ""

echo -e "${Cyan}Parser Test Results:${NoColor}"
go test ./parser/
echo ""
