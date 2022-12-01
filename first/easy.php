<?php


$input = <<<INPUT
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
INPUT;

$elves = explode("\n\n", $input);

$cal = [];
foreach ($elves as $key => $value) {
    $cal[] = array_sum(explode("\n", $value));
}
echo max($cal);
