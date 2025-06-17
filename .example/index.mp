:META name  LOTTERY
:META input float

#MAIN
VAR $SOME_VALUE (32)

IF (rand() > 0.5) #MORE #LESS

PRINT (" Try Next?")

#MORE
PRINT ("you're WIN!")

#LESS
VAR $OUT_RESULT (sum("you're lose ", $SOME_VALUE, " points!"))

PRINT ($OUT_RESULT)