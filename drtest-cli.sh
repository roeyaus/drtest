#!/bin/bash
NOCACHE="false"
TRIPS=0
CLEAR_CACHE=0
MEDALLIONS=
for i in "$@"
do
case $i in
    trips)
        TRIPS=1
        ;;
    clearcache)
        CLEAR_CACHE=1
        ;;
    -m=*|--medallions=*)
        MEDALLIONS="${i#*=}"
        ;;
    -n|--nocache)
        NOCACHE="true"
        ;;
    *)
        echo "Please run as follows:"
        echo "drtest-cli.sh (trips/clearcache) (-m/--medallions)=<comma separated list of medallion strings for trips> (-n/--nocache - don't use cache for trips)"
        exit 1
esac
done

if [ $TRIPS == 1 ]; then
    if [ "$MEDALLIONS" == "" ]; then
            echo "must input list of medallions"
            exit
    fi
    echo "getting trips..."
    wget --method="POST" --header="Accept: application/json" --body-data=$MEDALLIONS -q -O - http://localhost:8080/trips?nocache=$NOCACHE;
    exit 0
elif [ $CLEAR_CACHE == 1 ]; then
    echo "clearing cache..."
    wget --method="POST" --header="Accept: application/json" -q -O - http://localhost:8080/clearcache;
    exit 0
fi
echo "Please run as follows:"
echo "drtest-cli.sh (trips/clearcache) (-m/--medallions)=<comma separated list of medallion strings for trips> (-n/--nocache - don't use cache for trips)"
exit 1

