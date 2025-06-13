#MAIN
VAR $SOME_VALUE (int(32))

IF (rand() > 0.5) #MORE #LESS

#MORE
OUT ("you're WIN!")

#LESS
VAR $OUT_RESULT (sum("you're lose ", $SOME_VALUE, " points"))

OUT ($OUT_RESULT)