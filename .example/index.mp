#MAIN
VAR $SOME_VALUE (32)

IF (rand() > 0.5) #MORE #LESS

OUT (" Try Next?")

#MORE
OUT ("you're WIN!")

#LESS
VAR $OUT_RESULT (sum("you're lose ", $SOME_VALUE, " points!"))

OUT ($OUT_RESULT)