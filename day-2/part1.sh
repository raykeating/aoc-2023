#!/bin/bash

redLimit=12;
greenLimit=13;
blueLimit=14;

sum=0

mapfile -t arr < data.txt
for i in "${!arr[@]}"; do # for each line

    # Parse the line for easier access
    # get just the color part of the string
    colors=$(echo ${arr[i]} | cut -d ":" -f 2)
    # replace the semicolon with a comma
    colors=${colors//;/,}
    # replace the comma and space with just a comma
    colors=${colors//, /,}
    # turn colors back into a string
    colors=$(echo ${colors[0]})

    redMax=0
    greenMax=0
    blueMax=0

    # use a comma to delimit between each color+count statement in the line
    IFS=","
    echo "|$colors|"
    for colorAndCount in $colors;
    do
        # use space to separate color+count statement into count and color
        IFS=" "

        array=( $colorAndCount )

        count=${array[0]}
        color=${array[1]}

        if [ "$color" = "red" ] && [[ $count -gt $redMax ]]; then
            redMax=$count
        fi

        if [ "$color" = "green" ] && [[ $count -gt $greenMax ]]; then
            greenMax=$count
        fi

        if [ "$color" = "blue" ] && [[ $count -gt $blueMax ]]; then
            blueMax=$count
        fi
    done

    echo $redMax
    echo $greenMax
    echo $blueMax

    if [[ $blueMax -le $blueLimit ]] && [[ $redMax -le $redLimit ]] && [[ $greenMax -le $greenLimit ]]; then
        sum=$(($sum + $i + 1))
    fi

    echo "------"

    # reset color counts for next line
    redMax=0
    greenMax=0
    blueMax=0

done < data.txt;

echo $sum