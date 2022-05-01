#!/bin/bash
#set -x
twid=${1?usage: ./today.sh userid}

today=$(date +%Y%m%d)

file=$today-twlist.log
 
if ! ./twlist -name=$twid > $file; then
    exit
fi

befile=$(ls  20*-twlist.log | tail -2 | head -1)
if [ "$befile" = "$file" ]; then
    ls -l "$file"
    exit
fi


# show diff
if [ "$DIFF" = no ]; then
    :
elif [ "$DIFF" = diff ]; then
    diff "$befile" "$file"
elif [ "$DIFF" = open ]; then
    # XCode Filemerge commandline interface
    opendiff "$befile" "$file" &
#elif [ "$DIFF" = meld ]; then
else
    meld "$befile" "$file" &
fi

