.CODE
ORG 00
LDA CONST_3
ADD CONST_4
STA TMP0
LDA CONST_2
NOT
ADD CONST_01
STA TMP2
LDA TMP0
ADD TMP2
STA TMP1
LDA TMP1
STA A
LDA A
ADD A
ADD A
STA TMP3
LDA TMP3
STA Y
HLT
.DATA
ORG 20
CONST_3 DB 3
CONST_4 DB 4
TMP0 DB 00
CONST_2 DB 2
TMP1 DB 00
TMP2 DB 00
CONST_01 DB 01
TMP3 DB 00
A DB 00
Y DB 00