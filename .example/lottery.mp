:META name  LOTTERY
:META input float

#MAIN
VAR (32) = $SOME_VALUE

IF (rand() > $INPUT[0]) > (#MORE #LESS)

PRINT (" Try Next?")

#MORE
PRINT ("you're WIN!")

#LESS
VAR (sum("you're lose ", $SOME_VALUE, " points!")) = $OUT_RESULT

PRINT ($OUT_RESULT)