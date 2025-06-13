#MAIN
VAR int(32) > $SOME_VALUE

IF (rand() > 0.5) #MORE #LESS

#MORE
OUT "you're WIN!"

#LESS
VAR (sum("you're lose ", $SOME_VALUE, " points")) > $OUT_RESULT

OUT $OUT_RESULT&