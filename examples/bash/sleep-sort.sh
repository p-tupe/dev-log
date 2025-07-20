#!/usr/bin/env bash

arr=(1 5 9 7 2 8 4 3 6)

for i in ${arr[*]}; do
  sleep $i && echo $i &
done

wait
echo "Voila!"
