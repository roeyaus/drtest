#!/bin/bash

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
    -n=*|--nocache=*)
    NOCACHE="${i#*=}"
    ;;
    *)
            # unknown option
    ;;
esac
done

if [ "TRIPS" == "1"]