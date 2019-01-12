#!/usr/bin/env bash

function rand(){
    min=$1
    max=$(($2-$min+1))
    num=$(($RANDOM+100000000)) #增加一个10位的数再求余
    echo $(($num%$max+$min))
}

rnd=$(rand 40000 50000)
