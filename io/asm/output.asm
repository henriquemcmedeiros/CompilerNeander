.CODE
ORG 00
LDA CONST_1
ADD CONST_2
STA TMP0
LDA TMP0
ADD TMP0
ADD TMP0
ADD TMP0
STA TMP1
LDA TMP1
STA y
HLT
.DATA
ORG 20
CONST_1 DB 1
CONST_2 DB 2
TMP0 DB 00
CONST_4 DB 4
TMP1 DB 00
y DB 00