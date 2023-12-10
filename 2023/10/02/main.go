package main

import (
	"fmt"
	"strings"
)

var input0 = `.....
.S-7.
.|.|.
.L-J.
.....`

var input1 =
	`..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

var input2 = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

var input3 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

var input = `7-J-..FF.FLJJ7-7-L|7FFF7F7FF..|7-77.F77L---|7L-77-FF.LF-J77F-7L7L7|--.F-|.FFFF-|7-|.F--|FF.F.F77-L7-|7FF.FJ7FLL7-J7-|7|-L7-|7.7FF|-7-F.7777|
|-|--FJ|F|..LLFL7FJJF-LJ-J-F|-L7-J-FJFJ|FF--J.L|J-|L-7LJJF-JJFLL-|J.J-LF7--7|..|L|L--.FLLJFJ-FJF-L|7|--J--..--|J7.F.LF..||.LJ.F-F7-|LJ-|F--7
|--.L|L|.L--7.L.|7|.F-J7F-7|JFF-7LF-7LL7LLL7LFLJ7LL7.F7L--.|.F7J.L7-|.7J|J|L|F--7LL.FF7||FF|L|7|-FF|J7F||LFLJ7|F7-FF.LJFF7-7FFJ.L-F77.|F||.-
|.LL7L.|7FJ||7JL-LF-|77|J||JF7L|J.|7|..|7L.L-|JFF-J-FLL7||.--.|FJFL7L.L.JF7JLFJ||7L77-L-F-F7-LF--JFJL|JL-7|JFF|J.7-|7F|LJF7F7JL7.F|L.F-|J77.
|-L-F|FF-L7J.L7J|LJ..|7JLL7-F7J|LFF77|.LL-|L-|FF|||.F7.-7-7J.FLL.-.L.LJFL|J7-|.J7FFLJ7FFJJFJFF|.L-L-FL-7L|JFJ-7J.F-FL-F7.L7JJ7||L---F7JL..F7
-7.L.FFFJLJ7.LF.JF7J-L7-JLLL||FF-LJL-L.7F7JFLLJJLJ-||F.F|.|.F|7.FJ7LFL--7.L7.|-L77JLJ|F7.|.-|7.L.|LF77|L7JF7-7J.LF7JL.L|FFJJJ.F|J|7F-F7F-FF-
|F-.FL7|.F7--J|-L-FJF|LL77JFJ|-|-|77L|-|L-F77|.|LJJ||-7JJFJ-7L|7|LJ-L-FJ|FFF7F7JL7-L---.7.FJ|7|L7-7J|F77LF7|7.77LJ.|F|-F7-J.|..F-77F7|J7.|FJ
7J.|7J|J-LJJ|LJ-FLJLL-7.|FFL7|JF7LF|.F77L7|L7-77JJ-FJ-J.FFJJL-FLL.JFFLJ7||J.-J-7F||.|.|.F7|.LJJ-|-JL|777-|L7J-FJ|JFFLJL7J7L7FF-L7L7.F7FJ--||
|LFJ7F|JJ.L7|7|F77.7.|-LFF-7||F||7-.F||7.FJFJF7F---7--7-JFLF7.FJLF.FL..FL-7|J.L-L-J7|FFL|---J.77||F7F7F--JFJ|7L|-7F7-||.LLLF7|.---77F7J.FFJ7
J-FJ|F|JLF.LL---7J7F|||.|L7||L7||J7.F||F7L7L-J||F--J7FF.F77F|7LF-F7.L7-J-|||--JJ|L|7-77||.LFF.LF-7|||||F-7|F7--.LF7LF-7|7FFJLLF-|J|L-J-F-|--
|.JJL7|.F|7J|LL.|LJFL77-|7||L7||L-7F7|LJL7L7F-J||F7F7-F7|L7J.7.LF||-F|7-F-L77|J.L.LL7F77|FFF-7FJFJ|LJ|LJJ|LJL7JFFJL77FF7--LJF-|||LLJ-L--|LFJ
7JF-L|7.-7L-77-F7-F--L7.|FJL-JLJF-J|||F7FJFJL--JLJ||L7|||FJ.F7-JF|L-7F-7|7JLJ7.7.-7-FJL7|-FL7|L7L7L-7L7F-JF--JF7L7FJF-J|JJ.L.||F-JL7.LFJF-L7
7JF7|.|FJJ.|J-F77-L.FJF-LL-7F-7FJ|FJ|||LJ|L7F--7F7|L7LJ||L-7FL7F7|F-JL7|F7FJ||.F77|FL-7L7-F-JL7|FJF-JFJL-7L-7FJ|FJL7|F-J7|FLFL7J|F7J|7LF-7L-
L.||--|||F7J|F|.|-7F7-|7|LF||FJL7FL7LJL7F--JL-7LJLJ7|F7LJF-JFLFJ||L7F-JLJ|J7F7-|L7F-7|L7|7L-7FJ|L7|F-J|F7L7FJ|FJ|F7LJL7J-|7FF7LFJ7L||7F||J.|
L|7.|.FJFFFJF7-LLF-|JF|FJFLLJ|F7L7FL-7FJL7F---JFF--7LJL-7L-7F7L7||FJL-7F-J|-|F7L7LJFJ-FJL7F-J|7L7LJL7-FJL7|L7|||LJ|F--J.|L--L|J|F|.-J--L|7F7
|||FF7FJ-L.-L-J|.|-L-7F|-LLF-J|L-JF--JL7FJL--7F7L7FJF7F-JF-J|L7||||F7FJL--7-F7F7|F7L-7L7FJL-7L-7L7F7L-JF-J|FJ|L7F-JL-7F7-F|L7|-LFJ7J|.J.|F-J
F7--J-J|.FL.|JF77.7-FLF7.L|L-7|F7-L7F-7||F7F-J||FJL7||L-7|F7L7||LJ|||L7F--JFJ|||LJL-7|FJ|F-7|F-JFJ||F-7L7FJL7|FJL7F--J|L-7JL--F.|-J.FF-F---|
77.|F|-|FLJF77|J7-|FJF||F7F7FJLJL7FLJF|LJ|||F-J||F7||L7FJLJL7|||F-J|L7|L7F7L7||L7F-7|LJFJL7||L-7L-J||FJFJ|F-J|||FJL7F7|F-JJ-L.JJJJ7L7J7|J..-
LF7.-|.F|JLF7-.L7..LJ-|LJ||||F---JF7FFJF-JLJL-7LJ|LJL7LJF---J|LJL7FJFJ|FJ||L|||FJ|FJL-7L-7|||F-JF-7LJ|FJFJL7-|L-JF-J|||L7F7J.-7J|FJFJFLJ.-||
FL-J||-|.|.LJ7L7F-7|7FL-7||LJL7F7FJL7L7||F7F-7|F7L7F7L-7L7JF7L7F-J|FJFJL7||FJLJL7||FF7L-7LJLJL--JFJ7FJL7|F7L7|F-7L7FJ|L7L-7F7J7--J|-F|-F7-|L
FL7|F-7FF--7L7.FL7|F-7F7||L--7LJ|L-7|FJ|FJ|L7|LJL7LJL7L|FJFJL-JL-7|L7L7FJ||L7F--J|L7|L7FJF-7F7F--JF-JF7|LJ|FJLJFJFJL7|7|F-J-F--..L|-J|F7L7|J
J-|J|LLJ-7-JJL7|L|LJFJ||||F7FL-7|F7|||FJ|FJ-|L7F7|F--JFJ|JL7F----JL7|FJ|FJ||||FF7L7|L7||FJFLJLJF-7L-7|LJLFJL-7FJFJF7||FJL----77-L-J7J7.FJ7J.
.F.FJ7.J7J-|.JL|7|F7|FJ||LJL7F7|LJLJ||L7|L-7|FJ|LJ|F7FJFJF7||F7F7F7||L7LJFJFJ|FJL7||FJLJ|F----7L7|F7|L7F7L7F-JL7|||LJLJF-----J77.FJL7LLJ.7JL
.7-J7FJ.F7.L7JF7FLJ|||FJL--7LJ||F--7||FJ|F-J|L7L-7|||L7L7|||||LJ||||||L7FJ|L7|L7FJ||L--7||F---JFJ||LJFJ||FJL7F7||FJF---J-F7F77.-7J.F|-L7-J.|
F-JFL|L-J77.FF|L---JLJL7|F7L7FJLJF-JLJL7||F-JFJF-JLJL7|FJ|||||F-J|LJ|F7|L--7||FJ|L||7F7||||F--7L7|L-7L7||L7FJ|LJ||FJF7F--JLJL77..LF7JF-JJF-|
|7-7J|7.L|7-JJL--7F--7FJFJL-J|F7J|F7F-7||||F-JFL---7FJ||FJ||||L-7|F7||||F--J|||FJFJ|FJ||||||F-JFJ|F-JFJ||F||FJF-J||.|LJF--7F7L-77FF---7-LJFJ
L|-F7|FL7L-J-FJF7||F-J|FJF--7||L7LJ||FJLJLJL7F7JF7L||FJ|L7||||F-J||LJ|LJL7F7|||L7L7LJFJ|LJLJL7FJFJL-7|FJL7|LJFJF7|L-JF7L-7LJ|F-J-LL-|.L.LFJ.
F7F|.-J.F||.FJ.|LJ||F7|L-JF7|||FJF7LJL--7F--J||FJ|FJ|L7L7|LJLJL-7||F-JF--J||||L7L7|F7|FJF--7FJL7|F7FJLJF-JL-7L-JLJF7FJL--JF7|L-7J.FFFF7F-LJ|
F-J7-LLL.FF77-F|F-JLJ||F7FJLJLJL7||F77F7||FF7||L7|L7|FJFJL7F---7LJ||F7|F7FJLJ|FJFJLJ|||FJF7LJF-J||||F--JF7F7L7F---JLJF7F--J|L7FJ7-FFFJ7-L|-F
L-J|..|.FFJL7F7LJF---J|||L-----7|||||FJ|||FJLJL7||FJLJFJ7F|L7F7L--JLJ||||L--7LJFJF7FJLJ|FJ|F7L-7||||L-7FJ||L7||.F----JLJF7FJ7LJFF7|L-7L7-J-7
7JFJ7-F--L-7LJL-7L---7|||-F-7F7||||||L7LJLJF7F7|||L--7L7FFJFJ||F7F7LFJ||L7F-JF-JFJ|L--7||FJ||F7|||||F-J|FJL7||L-JF-----7|LJF-7FF-7L7F77J|LL.
|.J7|.|.L|-L-7F7L7F--JLJL7L7LJLJ|||||F|F---JLJLJ||F7FL7L7L7L-J||LJL7L7||FJL-7|F7|FJF7J|||L7|||||LJ||L7L||.FJLJF-7L---7LLJF-JFJ-|FJ7L-FF7|LFJ
77|L77|F--7F7LJ|FJL----7FJ-L7F-7LJLJ|FJL----7F7|||||F7|FJ-L-7FJL7F-JFJLJ|F77||||||FJL7||L7||LJ|L-7||FJFJL-JF7FJLL-7F7L--7|F7|.FJL77|J||||77|
L.|7L77L|L-|L7J||LF7F--JL---JL7L---7|L7F----J||FJ||||LJ|F7F-J|F7||F7L--7|||FJ||LJ|L7FJ|L7|||F-JF7|LJL7L-7F-J|L7F7FLJL---J||LJFJF7L-7.-LJ|7L-
L-L--L|-JFJL7L-JL7||L-------7FJ.F--J|L||F77F7|||FJ||L-7|||L7FJ||||||F7F|||||FJL-7|7||||FJ||||F7||L7F-JF7||F7L-J|L7-F--7.FJL7FJFJL7FJ--J|L-J7
LF7.FFJ.F77|L---7||L-7F--7F-J|F-JF-7L7||||FJLJ||L7|L7FJ|||FJ|FJ||||||L7|LJ||L7F7|L7||FJL7|||LJ|||FJL7FJ|||||F7.L7L-JF7L7|F-JL7L-7LJ7JFLJF|F|
.LJFF7L-F-7F----J|L-7|L-7|L-7|L--J7L7||||||F--J||LJ-|L7||||FJL7||||||FJL-7|L7LJ||FJ|LJF-J||L7FJ|||F7||JLJLJLJL7F|F--JL-J||.F-JF7L7J-.L.LF-J7
|LL7JL-7L7LJF---7L7FJL--J|F7LJF7F---J||||LJ|F-7L7F--JFJLJLJL7FJ||||||L7F7||FJF-J|L7L-7L-7|L7|L-J|||LJ|F-7F----JFJL----77|L-JF7|L-J||7-7..LJ|
LFF--7LJ|L-7|F--JFJL----7|||F7||L---7|||L-7LJLL7|L7F7L----7FJL7||||||FJ|||||FL-7|FJF-JF-J|FJL-7FJ||F7|L7LJF7F7-|F-----JFJF-7||L7F7F77-F-7-7-
.-7.|FFF|7LLJL--7L7FF---JLJLJLJL----JLJ|F-J7F--JL7||L7F-7FJL7FJLJ||LJL7||||L--7LJL7L7FJF7||F7FJ|FJLJLJFJF7|LJL-JL-----7|FJ-LJL7||LJL77L.7--J
.FJFL-F7L77F7LF7L7L7L--7F7F-7F--------7|L--7L--7FJ||FJL7|L7FJ|.FFJ|F7FJ|||L7F-JF--JFJ|FJ||||||L|L-7F--JFJLJF7F-7F-----J||F7F-7LJ|F--JJ7-L||.
LLJ7.L-7-|-||FJL7L7L7F7LJLJ.LJJF--7F--J|F7FJF--JL7||L7FJL-JL7L7FJFJ|LJFJ||FJL7.L--7|FJL7|||||L7L--JL---JJF7|LJJ|L7F--7FJ||LJFJ-FJ|J|J.F7-F77
L||J7.LL7||||L-7L7L7||||F--7|F7L-7LJJF7||LJ.L-7F7|LJFJL-7F--JFJ|FJLL-7L7||L7FJF--7|||F7|||||L7L7F------7FJLJLF7L-J|F7|L7||F-JF7|FJ77F7F7.|LJ
F7JJFJ-LLF-JL-7L7|FJLJL-JF7L-JL-7L--7|||L7F-7|LJ|L7FJF-7|L--7L7||FF--JFJLJFJL-JF-J||LJ||||LJFJFJL----7FJL----JL-7FJ|||FJLJL7-|||L-7F--J|7J|7
LL||L|JL-L---7L-J|L-7F7F-JL-7F-7L-7FJ||L-JL7L-7LL7|L7|FJL7.FJFJ||FJF-7L7-FJF7F7L-7|L7-||||F-JFJF-7F7FJL-7F------JL7|LJ|F--7L7|LJF7LJF--J|FF7
.F-L7F77.F7F-JF-7L--J|||F---JL7L-7|L-JL-7F7|F-JF7LJFJ||F7L7L-JLLJ|FJFJFJFJFJ||L7FJ|FJFJ|LJL7FJ7L7LJLJF-7LJF7F7F7F7|L--JL-7L-J|F-JL7FJ-J||7L|
|.F|7|L--J|L-7||L7F7FJLJL7F7F-JF-JL--7F-J|LJL7FJL-7L-J||L7L7F7|F-J|FJFJ7L7L7||FJL7|L7L7|F7FJL--7L--7FJ-L--JLJLJ||LJF7F7F7L-7J|L--7LJJLL7FJJJ
|7JJFJF7F7L--J|F7|||L-7F7LJ|L77L7F-7|LJF7|F--J|F--J-F7LJ-L7LJL7|F7||FJF--JFJLJL7FJL-JFJ||LJF---JF--J|F7F-------JL7FJLJLJL-7L7|F--JF7J7|FJ|7|
||FFL-JLJ|F7F7||LJ||F-J|L--JFJF-J|FJF7L|LJL7F-JL7F7FJL-7F7L7F-J||LJ||7L-7FJF-7FJL--7.L-JL-7|F7LFJF--J||L---7F7|F-JL------7L7||L7F7||7|-LL|-J
-L7|JF---J|||||L7FJ||F7L---7|-L7FJL-JL-JF-7LJF--J|LJF7FJ|L7|L7FJL7LLJF--JL-JFJ|F7F7L7F7F--J||L-JFJF--JL7FF7LJL-JF-7F7F7F7L7LJL7LJLJL---77||L
|.LJJL----J||LJJ|L7|||L--7FJL7-LJF-----7L7|F7|FF7|F-JLJFL7LJFJ|F-JF--JF--7F-J-LJ||L7LJ|L--7|L7F7|FJF7F7L-JL----7|-LJLJLJ|FJF7||F----7F-J7-|J
--L.7.F----JL7F-JFJ||L--7LJF7|F-7L----7|FJ|||L-JLJL-----7|F7L7|L7-L-7FJLFJL-7F7FJ|LL7FJ7F7LJJLJLJL-JLJL-------7|L---7LF7||FJ|FJL7F-7LJLJ.FL7
.7...FL7F7F-7||F7L7||LF-JF7|LJL7|F7F7FJLJFLJ|F7F7F7F7F7FJ|||FJ|FJF--JL7FJF-7LJ|L7|F7LJF-JL-------7F7JF----7F7FJ|F7F7L7||LJL7LJF7LJFJ-L7|-J.|
FLF7-|.LJ||FJ||||FJLJFJF7|||F7FJLJLJLJF----7||LJ||LJLJ||FJ||L7|L7L-7F7||FJFL7FJFJLJL-7L7F7F7F7F-7LJL7L---7LJ|L7LJLJL7LJL---JF7|L-7L-7LF7.|F|
FF|--FFF-J|L-JLJLJ7F7L-JLJLJ||L-----7F|F--7||L-7|L-7F7|||FJ|FJL7|F-J|||||F--J|FJF-7F7|FJ|LJLJLJLL--7|F---JF7L7L--7F7|F7F7F7FJ||F-JF7L-JL7F7J
-FJ|-|-L-7|F--7F7F-JL7F---7FJL------JFJL-7|||F-J|F7LJ|LJLJ7||.L||L7FJLJ||L-7FJ|FJFJ|||L7|F--7F-----JLJF7F-JL7L7F7LJ|||||||LJFJ||F7|L-7F7LJ|J
|.F|-|F|L||L-7|||L--7|L--7LJF--7F----JF7FJLJ|L-7||L--JF---7LJF-J|FJL7JFJ|F-J|-|L7L7|LJ.LJL-7|L7F7F----JLJF--JFLJ|F-J|||||L-7L-JLJ||F7LJL--J7
F7.LL-77JLJF7|LJL---JL7F-JF7L7FJ|F7F-7||L7F7|F7||L----JF-7L7-|F-JSF7L7L7|L7FJFJFJFJ|F7F7F7FJL7LJLJF------JF----7|L-7||LJL--JF7JF-JLJL------7
7|FF7.F7F7FJ|L-------7|L--JL-JL7|||L7LJL-J||LJLJL7F7F7FJ-L-JLLJJ7LJL7|F|L7||FJFJ7L7||LJLJLJF7L7F7FJF7F-7F-JF7F7||F7|LJF----7|L7|F7F7F7F----J
-J-FJF|LJ|L7|F-7F-7F-JL7F7F7F--J||L-J-F---JL-7F7FJ||||L---7..J|JF-F-JL7L7||||FJ-F7LJL----7FJL7LJLJFJLJFJ|F-JLJLJLJLJF7L7F-7LJFJLJ|||||L----7
|J|..-L-7L-J|L7||FJL--7|||||L7F7|L-7F7|F----7LJ|L7||||F--7|7.F|.|J|F7FJ|LJ||LJF-J||F7F---J|F-JF--7L7F7L-JL---7F7F---JL-JL7L7FJF7FJ|LJ|F7F7FJ
|JLJFLL-L7F7L-JLJ|F---JLJLJ|-LJLJF7LJLJL---7L-7||LJ||||F7LJ7-LJ-F-LJ|L77.|LJ7FJF7L-J|L-7F7|L--JF-JFJ||-F7F---J||L7F7F7F-7L7LJFJ|L-JLLLJLJLJ.
LF-|7|.FJLJL----7|L-----7F7L7F-7FJL-7F----7L-7|L--7LJ|||L7JJFLJFLLFLL7L7-J|LFJFJ|F-7L--J|||F-7FJF7|FJL-JLJF7F-JL7||LJLJLL7L--JFJJ7|7||L7LL-J
FJFF7FF7F7F7F7F7||F7F7F7LJL7|L7|L--7LJF--7|F-JL7F7L7FJ||FJJ|JLF|-FLJLL-J.LJ.L7|FJ||L7F-7|LJL7LJFJLJ|F7F---JLJF--J||F-7F-7|F-7FJ|L|-JJ|J|JFJ.
LJ|JFFJLJLJLJLJLJLJ||LJ|F--JL-JL--7L-7|F-J|L--7||L7|L7LJL7L7.F-JJL|F|JLLF7JFL||L-JF-JL7LJF--JF7|F7-LJLJF7F---JF77|||FJ|FJLJ-LJ-L7LLJ.L-F7|.7
FF--F|F-7F-7F7F-7F7|L7FJL7F7F-7F-7L--J||F7L---JLJ.LJL|F-7L7J-LLJF-7-|7-.-|.7.LJLF7L7F-JF7L---JLJ||F7F7J||L7F--JL7|LJL-JL--7.|LJ.7.L.7J.L7F.7
L||J.LJFJ|FJ|LJFJ||L-JL--J|LJFJ|FJF7.FJLJL7F7F7F7-F7LLJ|L-J|--7.L|.LL|J|JL-|F7-FJL7LJF7||F7F7F7FJ||LJL-JL-J|F---J|F----7F-JF77LF|7.FJ-FLLJ-L
L7FF--7L-JL-JF-JFJL7F7F7F7L7FL7|L-JL-JF7F7LJLJLJL-JL7.F77F7---JJJ|F7-F-|7--FJL7L-7L--JLJLJLJLJ|L7|L-7F-7F7F|L----JL---7|L--JL7-F-7-F|7-JL|77
FJFL-7|F--7F7L7FJF7LJ||LJL7|F-J|F-----J||L7F7F---7F-J7|L-J|J7..LF--JF|F|L-JL-7L7FJF--7F-7F---7L-J|F-J|FLJL-JF-7F7F----J|F----JF7|77FF7L7-LJ|
J7F--JLJF7|||FJ|FJL--J|F7F||L-7|L--7F7J|L7||LJFF7LJF-7|F--J777-.F.L-FL|JL|.F-JFJL-JF7LJFJL7F-JF-7LJF7|F-----J7LJ|L---77LJF-7F7||7.7JL|-|J.F|
F-|F7F7FJ||||L-JL----7LJL7LJF-J|7F7LJL7L-J|L7F7|L7-L7LJL7JJJLLF-LJ7.|L|.FF-JF7L-7FFJL7-L--JL-7|7|F-JLJL-----7F7FL---7|F--JFJ||||L7LJ.LF|LLL7
J|||LJLJ.|LJL---7JF7FJF--JF7|F7|FJL-7FJF7LL-J||L7L7FJF7FJJ.F|FLJ.F-7F.|FJL--J|F7L7L-7|F7F7F7FJ|FJ|F---------J|L7F7F7LJ|F--JFJLJL-7-7F7.J.L-|
.FLJJFF7FJF7F7F7L-JLJFJF--J|LJLJL--7LJFJL7|F-JL7|FJ|FJLJJJ7FL77F-J---7FF-JL|.LJL7L--JLJ|||||L-JL-JL--7F------JFJ|||L--JL---JF-7F-J.F7L|.|--7
---L-FJLJFJLJLJL-7F7FJ|L--7|F7F7F7LL--JF-JFJF-7LJ|FJL7|F7-L77.L-7-FJ|7L-7F7.LLF-JF7F--7|||||F7|F--7F-J|F------JFJLJF--7F----JFLJJLFJ|7|.F7LJ
.|7|FL--7L7.F----J|||JF--7|LJLJLJL---7FJF7|FJ||F7LJF-JFJ|7F|-F|J|L|JL|7-77L-7LL-7|LJF-J|||||||FJF-JL--JL--7F-7FJF-7L-7|L---7F7FF7LL7|F|L|7.|
F|-F||F-JFJFJF-7F7||L7L-7|L7F7F7F7F7LLJFJLJL-7LJL-7L--JFJ--7.JJFJL77JL7.LJ-LF7F-JL7|L-7LJLJ|||L7L77F7F7F--J|FJ|FJJ|F7||F7F7LJL-JL7FJ|F-J||.7
|7F|JF|F7L7L7L7||||L-JF-JL-J|LJLJLJL7|FJF-7F7L7F-7L7F-7L7J-LJL|-J7|.7.|F7FFJLFJF7FJF--JF7F7LJL7L7|FJLJ|L---JL-JL-7||LJLJLJL------J|FJ7.FF7J7
|LLL-7LJL-J||FJ||LJF--JF-7F7|F------JFJFJJLJL-J|FJ.|L7L-JJ..JFLL-F77J7-LLLJ7-L-JLJFJF7FJ||L7F7L-J|L--7|F------7F-JLJF----7JF7F--7-|L-7F7||LF
J-F|7|JJFF--J|FJL7JL-7FJFJ|LJL-7F----JFJLF7F7F7||F7L-J7F7-7|FJLL7L|L7JFFLLL|.|F-7FJFJLJ-LJ|LJL--7|F77||L---7F7|L-7JFJF---JFJ||F-JFJF-J||||.F
|.7L|LJ.FL-7FJL7FJF--J|LL-JF--7LJF---7|F7|LJLJLJLJL7F-7|||FL||7.F-7LF77FLL|LLFL7LJFJF--7F7F7F7F-JLJL-JL7LF7LJ||F7|FJFJF--7L7LJL7||FJF7|LJ|F7
-JJ---L|JFFJ|F-JL7L7F7|F--7L-7L--JF7L|LJLJF7F------J|FJ|L7JFF---7JJ7FJF77JF-JLFJF7|JL-7||||||LJF--7F7F7L-J|F7|LJLJL7L-JF-JFJF-7|FJL7|||F-J||
FJ7J||.|LFL-JL7F7L7LJLJL7FJF7L----JL7|F7F7|||F---7F7|L7|FJF7L7F-JJ|FL-FJ|.|F7|L-JLJF--JLJLJLJF7|.FJ|||L7F7LJLJF-7F7|F--JF7L7||LJ|F-J||||F7||
F-JJF7-|JL.|F7LJL-JF-7F-J|.|||F7-F--JLJLJLJ|LJF-7LJLJFJ||FJ|FJ|F7LF7J||||-F77F-7LF7L---7F-7F-J|L7L-J||FJ||F7F7|FJ|LJL7F7|L-JL7JF||F-JLJLJ|||
77L-L7.L7FFFJ||F7F7|FJL-7|FJL-J|FJF7F7F--7FJF7|-L----JFJLJFJL7LJ|FJ|7--|J||L7|FJFJ|F---JL7LJF-JFJF--J|L-JLJ||||L7L7F-J|||F--7L7FJLJF7F7F7LJ|
-LJ|F-JFL-|L7L-JLJLJL--7||L7F-7|L-JLJ||F-JL-J|L-7F7F7-|F7FJ.LL-7||FJJ.FJ7FL7LJL-JFJL-7F7FJF-JF7L7L---JF7JF7LJLJFJL||FFJLJL-7L-JL7F-JLJ|||F7|
..|FJ7L-..FFJF--7F7F7F7||L7LJFJL7F7F7LJL---7JL--J|LJL7||LJ7FF77|LJ|J|-FLF-7L----7L7F7|||L7L--JL7|JF7F-JL-JL7F-7L-7||FJF7F--JF7FFJ|F7F7LJLJ||
L-LJ|J7LF.LL-JF-J|LJLJ||L7L--JF7LJLJL--7F7FJF7-F7|F--J|L--7FJL7L7FJF7F7-L7|F7F7FJFJ|||||FJF---7LJFJ|L-----7||FL-7LJLJFJ||F--JL7L7LJLJL---7||
F|L7|F---7J7F|L--JLF7FJL7L-7F7|L7F----7LJLJFJL-JLJL---JF--JL7FJFJL7|LJ|F7|LJLJ|L7|-||LJLJFL--7L-7L7|LF7|F-J|L--7|F-7FJ|LJ|F---JFJF-7F----J||
F|||J|L-7.LF-F.F7F-JLJF7L--J|||FJ|F---J7F--JF7F-------7|F7F7|L7L7FJL7FJ||L---7L7|L7|L7F7F7F7FJF-JFJL-JL7L7LL-7FJ|L7|L----JL----JFJJ|L-7F-7||
.J-F-L7.F.F7F7FJ|L----JL-7F-JLJL-JL--7F-JF7FJ|L-----7.|LJLJ|L7|FJ|F7|L7||F7F7L7|L7||FJ|LJLJ|L7|F7L----7|LL7F7LJFJFJ|F--7F7F--7F7L-7|F7LJFJLJ
F.F|.J7FJF|LJ|L7|F-------J|F7F77F7.F7||F-JLJFJF--7F-JFJF7F-JFJ|L7||||FJ|LJLJ|FJL7|LJL7L---7||||||F----JL--J|L-7|FJJLJF7|||L-7||L-7|LJL7FJ..J
L-|JF-FF77L7FJFJ|L--------J||||FJL-JLJ|L---7L7L-7LJF7L-JLJF7L7L7|||LJL7L---7|L-7||F-7L7F-7|L-J|||L-7F------JF-JLJF7F-J|LJL7FJLJJFJL7F7||7F77
|7L-|JJ||F-JL7L7L7F--7F7F7FJLJLJF--7F7|F---J-L--JFFJL7F-7FJL7|FJ|LJF--JF7F7||F-JLJL7L-JL7|L-7FJ|L-7LJF7F7F7FJ|F7FJLJF-JF7FJL---7L--J|LJL-7-7
L7F-L7-||L-7FJFJFJL-7LJLJLJF7F7FJF-J|||L------7F7FJF-JL7|L-7LJ|FJF-JF77|||LJ|L--7F7L-7F7||F-JL7L7FJ7FJLJLJLJF7|LJF7FJF7||L7F---JF7F7L-7F7|J|
|L7FF7-||JFJ||L7|F7FJF7F-7FJLJ||FL--JLJF---7F7|||L7|F--J|F7L-7|L7L--JL-J|L-7L7F-J|L--J|LJ||F-7L-JL7FJF------JLJF7||L-JLJL7||F7F7|LJ|F7LJLJFL
-JLL|L-JL7L7|F-J||||FJ||FJ|F--J|F7F--7FJ|F7LJLJ||FJ|L7F-J||7FJ|JL7F-----JF7|FJL7FJF7F7|F7|LJFJF-7FJL7||F---7F--JLJL7F7F7FJLJ||||L7FJ|L7F7-LJ
|.|JL-7F7L7||L-7|||LJ|LJL-JL---J||L7L||F7||F7FFJ||FJFJL7FJL7L7|F7|||F7F-7||||F7||FJ|||LJ||F7L-JFJ|F7LJFJF7-LJF7F7JFJ|LJLJF7F|LJL-J|FJFJ||..|
|7||.LLJL7LJL--JLJ|F7F7F7F7F-7F7|||L7|LJLJ||L7L7|||||F-JL-7|-|||||L-J||FJ|LJ||||LJFJ||F7||||.F7L7LJL7|L-JL---JLJL7L7|.F--JL-JF7F7FJL7L-JL-77
|LJ-7|L7JL7F-7F7F7|||||||||L7||LJL--JL7F--J|FJF||||FJL-7F-J|FJLJ||F--J||LL-7LJLJF7|FJLJ||LJ|FJ|.L-7FJF-------7F-7L7LJFJF7F--7|LJ||JFJF-7F-JJ
7LL-FJF|.FJ|||||||LJLJLJLJL-JLJF7F---7||F7FJL-7||||L7F-JL7FJ|F-7|||-F7|L7F7L--7FJLJL--7||F-JL7|F7FJL7L----7F-JL7L7L--JFJLJF7LJF-JL7L7|JLJ7||
L7JF---.LL7|FJ|LJL7F-7F7F7F7F-7|LJF--JLJ|||F--J||||FJ|F7FJ|LLJJ|LJL-J||FJ||F--J|F7F7F-JLJL7F-J|||L-7|7F---JL--7|-L7F--JF--JL-7L--7L-JL--7F7-
LF-JJJ-FLFJ|L-J.F-JL7|||||||||LJF7L---7FJLJ|F7FJ|||L7LJ||FJF-7FJF-7F7||L7||L--7||LJ|L---7FJL-7LJ|F-JL7|F-----7LJF7LJJF7L--7F-JF7-L7F7F--J||J
7.7JFJ-7JL7|F--7L---J||LJLJLJF--J|F---JL-7FJ||L7||L7|F-J||JL7|L7|JLJLJL7|||LF7|LJF-JF7F7|L-77|F-JL-7FJLJ7F---JF-JL7F-JL7F-JL--JL-7|||L7F-J|.
J-JL|7.|-LLJL-7||F7F7LJF-7F-7L--7|L-7F7F7|L7|L7|LJFJ||F7||F7||FJL7F7F7FJ||L7||L-7L7FJ||LJF-JFJL7F-7||F7F-JF7F7L7F7LJF-7LJF-7F--7FJLJL7LJF7|7
|JL7|L.F7|FF--JL-JLJL-7L7|L7L---J|F-J|||||FJL7|L7FJJ||||||||||L7FJ||||L7||FJ||F7|FJ|FJL-7|F7L-7|L7|||||L--JLJL-J|L7FJJ|F-JFJ|F7LJLL|LL7FJ|L7
7L-|JF7-L7-L-----7F7F-JFJL-JF7F-7LJF7|LJ|||F7||FJ|F7||||||||||FJ|FJ|||FJ||L7|||||L7|L7F-J|||F7||FJ||||L7F-7F----JFJ|F-JL-7L7LJL7F7-|FFJ|FL7|
F.|.FLL7JJ7|F----J||L-7L--7FJ||LL7FJ|L-7||||LJ||-||||||LJ||LJ||FJ|F||||7LJJ||||||7LJF|L-7||LJ||||FJ|||FJ|7LJF--77L7|L----J|L7F7LJL--7|FJJ-LJ
LFL-7JJ|.FLFJF7F--JL-7L---J|FJ|F-J|FJF7||||L-7|L7||LJ|L77|L-7||L7L7LJLJF---J|LJ|L7F--JF7|||F7LJ|||FJ||L7L---JF7L7|LJF----7FFJ||F-7F-JLJJ-FLJ
.FL|LFF.FFLL-JLJF----JF--7FJL7||F-JL7|||LJL7FJL7||L-7L7L7|F7|LJJ|FJ|F7FJF7F7L7FJFJL-7FJLJ|LJL7F||||FJ|FJF7F7FJL7L---JF--7L7L-J|L7|L7J7FJF7J7
7-7||F7-F-7LLF--JF7F--JF7||F7||||F--J||L7F7||FFJ|L7FJFL7|LJLJ.F7||F7||L7|LJL7||FJF--JL-7FJF-7L7|||||FJL7||||L7FL7F-7FJF7L7L--7|FJL-J.JJ.|J|F
--|JL|7-|LJ7L|F--J||F--JLJ||LJ|||L--7||FJ|LJL7L7L7|L--7|L---7FJLJLJLJ|FJL-7FJ||L7|F7F7FJL7|JL7|LJLJ|L7FJ||||FJF7LJFJL-JL7|F7FJ|L-7-F7JFFL-|J
L-JJ|J|LJ.L--LJF--J|L--7F-JL--JLJ7F-J||L7L-7FJL|FJ|F7FJL7F-7||F-7F7F-J|F7FJL7|L7||||||L7FJL-7|L-7F-JFJL7||||L-JL-7L--7F7|||||F|F-JF|J-L7JLFJ
77|.|LJ-7.7-FF-JF-7|F--J|-F7F7F7F7|F-JL7|-FJ|7FJL7|||L7FLJJ||||FJ||L7FJ|LJF-J|F||LJ|||FJL7F-JL7FJL7FJ|FJ||||F7F-7L--7||LJ|||L7|L--77JJLLJL|7
L-7F7.F|JFF--JF7L7LJL---JFJLJLJLJLJL--7|L7L7L7L7FJ||L7|F--7||LJL7||FJ|FJF-JF7L7||F-J|LJF-J|FF-JL-7||F7L7|||||||FJF-7||L-7|||FJL-7FJJ.|.|J.7J
F|-JF--|.FL-7FJ|FJ|F----7|F--7F---7F-7|L-J7L7L7|L7LJ.|||F-J|L7F-J||L7||FJF7||FJLJ|F7L77L-7L7L-7F7||LJL7||||||||L7|FJ|L7FJLJ|L7JFLJ||.|.F--L7
LL.L|7|J7F7FJL7LJF7L---7|LJF7LJF--JL7|L-7F--JFJL-J|F-JLJL-7|FJ|F-J|FJ||L7|LJ||F7|||L7L7F7L7|LFJ||||F-7|LJLJLJ|L7||L7L7LJF-7|FJ.JJ.F|-JFJJ7.L
F|L.-JJL-FFJF7|F-JL7F--JL--JL--JF-7FJL7FJL7F7L----7L7F-7F7|LJFJ|F-J|FJ||||F-JLJL7|||L7LJL7|L7L7||||L7|L7F77F7L7||L7L7L7|L7LJL7.LF-|JJLJLL77|
FJ-J|LL7|.L7|||L--7LJF-7F7F-7F-7L7|L-7||FFJ||F7F7FJL||FJ|LJF7L-JL7FJ|FJFJ|L-7F7FJ|L-7L7F7||FJ7||||L7|L7||L-JL7LJ|FJLL7L7-L--7|7-|JJFF7J.|.|J
FJ-LF7|7.FL||LJF--JF-J-LJ||FJL7L7||F-J|L7|FJ||||LJ.FJ|L7L--JL--7FJL7||7L7|F7||LJFL--JLLJ||||F-J|LJ.||L|LJF--7|J|LJF7||FJJJ-L|L7-LJ.LJ|F--JJJ
J7L|F-L.-L7||7|L-7FJFF7F-J|L7FJFJ||L7FJFJLJFJ|||F--JFJFJF7F7F--JL7FJ|L-7LJ|||L-7F7FF-7F-J||||F7L7FFJ|FJF7L-7LJ--7FJ|FJL7.-JLL-J-7-7.LLJ|JL.|
LJL7J|JLL-LLJ-F--J|F-J|L-7|FJL7L7|L-JL7L-7LL-J|||F7FJFJFJ|||L-7F-J|FJF7|F-JLJF7LJL-JFJL-7|LJ||L7L7L-JL7||F7L7F7.FJFJ|F-JJ-77FL|7F..FJ.FL.F77
FL.||L-.FF.LF-L---JL-7L--J|L7FJFJL---7|F7L---7||LJ||FL7|-LJL-7||F7||FJLJL-7F7||F-7F7L--7|L-7|L7L7L7F--J|LJL-J|L7|FJJ||JJLF|-J-|JL7..|F|.FLFJ
F7FFFJ-LFF7FJLFF----7L-7F7|7LJ.L7F7F7|||L-7F7|||F-J|F-JL----7||LJ|||L7F---J|||||FJ||F7FJ|F7|L7|7L7|L7F7L-7F-7|FJ||F7LJJ7.LL|J||L-|7FF-J77.L7
L|FFL-7JL7J-|.FJF--7L-7LJ|L-7F--J|LJLJ|L-7||LJ|||F7|L7F-7F--JLJF-J||FJL-7F7|||||L7|LJ|||||LJ-LJF-JL7LJ|F7LJFJ|L-JLJL7JL-7.|..FL.FFJ7.7JL|7J|
FJLF---7JLFFL-L7L7FJF7L--JF7||F-7L7JF7L7FJ|L7FJ|||||FJ|L|L----7|F7|||F--J|||LJ|L7|L7LLJFJL-7F--JF-7L7FJ||F7L-JF7F7F7L7-.JJF-JFL|J-J|-F-7-|.|
L7.J.L||J.-JF|.L-JL-JL-7F7|LJ||LL7L7||JLJ7L-JL7|LJ||L7|FJF----JLJ||LJL7F7||L-7L-J|FJF7JL7F-JL7F7L7|FJL-JLJL7F-JLJLJL-JLFL7LJLF.|LJF7.|FLJL||
F-J.--JJ.||.F7|.LF-----J||L7|LJF7L7LJL7-F7F7LFJL7FJL7||L7L------7||F--J|LJL-7|F--J|FJL77|L7F-J||FJ|L7F-----JL-------7J.--|7|L|7|LF7LL7L7|J|J
7JL|7LJFJJ-7.LJ7-L-7F7F7|L7L---J|JL7F7L-JLJL7L7FJL7FJLJFJF7F-7F7|||L7F7L---7LJL--7LJF7L7L-J|F7||L7|FJL--7F-7F7F7F-7FJ7-J7F-JLJ-J.--F.--77-|.
|7J.-7JL7.|F-7LFF-LLJ|||L7L7F7F7|F-J||F-7F-7L7LJJF||F-7L7||L7||||||LLJL7F--JF----JF-JL7L--7|||||FJ||7FF-JL7||LJ||LLJJ|||-L-L7.L7FFL--|7J|J|.
--.J-L7L-7L-|-.LJFF--J|L7|FJ||||||F7||L7||JL-JJLF7LJL7L7||L-J||LJLJLLF-JL-7FJF-7F7L7F7L7F7|LJ|||L7||-LL--7|LJF-J|77FFJFJ|JF.|-7J-L7LLL7F|-F|
L|F|-JL7J|J.7F---|L--7|7||L7|LJ||LJ|||FJ|L----7J|L--7L7LJL--7LJLFJJ..L-7F7|L-JFJ||FJ|||LJ||JJLJ|FJ||FLJJLLJJLL--J-7J.-77|FJ7L7FF7F--LJJ7.FL|
F77J||..|.LF||.7F||7LLJ-LJFJ|F-JL7FLJ||.|F7F--J-L--7|FJF-7F7L77-L|JF-J-LJ|||F-JFJ||-||F7|LJJ-F-JL7LJJ7FLL.|-L||F|J|L7JL7LLJ7LLJ|F||LJ.F-7L-|
F-J---F.F7F-J|F|7.F7JLJ77LL7|L7F7|F--JL7LJ||JJ|7FF-JLJFJFJ|L7L-7L--|-||-FLJFJF7|FJL-JLJL--7J7L---JJ.L7-FJFJ7L|--FFJ-|JFLJ|LJ|L|F-|7JLJ.FJJ--
L|L7.F|-J-L-LJ-|.LLJ.|F|7.L||JLJLJL-7F7L-7||JF--LL7F-7L7|FJFJF-J-7L7-|--F--JFJ||L7F7F7F7F7|-7J|-JJ.|J|.7FF.|.|7F7|.|F--J.|J.L-7LJL777.-L.L|J
FJF--J7-L77|J.FL7|JJ|-7JF7-LJLJ.|FJ-||L7FJLJ-J.JJ.||JL7||L7L-JJFFL7|-7|.L-7FJL|L7LJ||LJLJLJ-F7JL|J.J.FF7-|-J-L|7|J.7-7.|FJ-7FL|-77.LJ.LJ--7.
|FL.J.|---7-7-7.F|7-J|LFJLLL-7FF|JL-LJLLJJ-L7L|JLLLJ-FJ||FJFJ.LF7L7JJLLL|FLJ|.L-JF-JL7.FL||LJ...|-FJ.J-F-LF|L|.-JJFL7F-JL|.7-.|JF--7J.F7J|J7
F--L-JJFF-L7J7|-JLF7L-J|-7.|FJL|JJ|L7LFJL|.LJ7|.|FLJL|FJLJ-L-FJF7|F.|F||LJ|LL7LJ.|F7FJ7-F--7J|FL.F7-F-LL-.|JF--J|-|F||FFF--JF7.-F.L|F-JJFJ.-
|F-JF.FL|.7..JJ7|FJF7-FLJJ|-J-||-FJ.F-7-F|FL|.L-77.|FLJLL.|.||.|LJJ-JJLL7||J.F.FFLJ||J.FF7.-7LL--F7FF77LLF7-7FLFLF-777L-7LF7LF-JLFF|L|J-F-F7
L-7J..7.L--J-LLFJF-L7F--|FL|.F77J|LJ77|L|J7.F7J.L7-|JJ|FJ|J--J.JLLF7|J.|L|F-LF.-L7L||L-J|---77L7LFL-LJ|---|FFJ-7-F.|.JF-|-F|7J-7JFFJF77.-77F
|-7-77-FJLJJ.J.J-|J.L.L----L-JJJ.FJ.LLJ-L..LJJ-F-F-JJ.-J7J..7J.L7--|--L-|JLJLF7.L7-LJ-JL--|-|-JL-LJ-LF.L-.L-7JL|L|-L7LLJL|.J7..J-J.L-7-7---J`
type position struct {
	Char    string
	Visited string
}

type vertex struct {
	X, Y int
}

type edge struct {
	From, To vertex
}

type Map struct {
	EdgeList []edge
	Cord [][]position
	StartX, StartY int
	X, Y int
}

func (m Map) FloodFill() {
	m.Cord[m.StartX][m.StartY].Visited = "I"
	m.FloodFillRecur(m.StartX, m.StartY)
}

func (m Map) FloodFillRecur(x int, y int) {
	if x == 0 || y == 0 || x == len(m.Cord)-1 || y == len(m.Cord[0])-1 {
		return
	}
	if m.Cord[x][y].Visited == "-" {
		m.Cord[x][y].Visited = "I"
		m.FloodFillRecur(x+1, y)
		m.FloodFillRecur(x-1, y)
		m.FloodFillRecur(x, y+1)
		m.FloodFillRecur(x, y-1)
	}
}



func (m Map) CheckUnvistedPoints() int {
	inPolygonCount := 0
	for y, row := range m.Cord {
		for x, col := range row {
			if col.Visited == "."  {
				inPolyCount := m.RayCast(x, y)
				inPolygonCount = inPolygonCount + inPolyCount
				if inPolyCount > 0 && inPolyCount %2 != 0{
					m.Cord[y][x].Visited = "I"
				} else {
					m.Cord[y][x].Visited = "-"
				}

			}
		}
	}
	return inPolygonCount
}

//func (m Map) CheckUnvistedPoints() int {
//	inPolygonCount := 0
//	for i, row := range m.Cord {
//		for j, col := range row {
//			if col.Visited == "." || col.Visited == "I" {
//				inPolyCount := m.RayCast(i, j)
//				inPolygonCount = inPolygonCount + inPolyCount
//				if inPolyCount > 0 {
//					m.Cord[i][j].Visited = "I"
//				}
//
//			}
//		}
//	}
//	return inPolygonCount
//}

func (m Map) RayCast(x int, y int) int {
	intersectionCount := 0
	if x ==61 && y == 13 {
		fmt.Println("here")

	}
	edge1 := edge{From:vertex{X:x, Y:y}, To:vertex{X:x, Y: len(m.Cord[0])-1}}

	for _, edge2 := range m.EdgeList {
		if Intersect(edge1.From, edge1.To, edge2.From, edge2.To) {
			intersectionCount++
		}
	}
	

	return intersectionCount
}

func ccw(A,B,C vertex) bool {
	return (C.Y-A.Y) * (B.X-A.X) > (B.Y-A.Y) * (C.X-A.X)
}

// # Return true if line segments AB and CD intersect
func Intersect(A,B,C,D vertex) bool {
	return ccw(A,C,D) != ccw(B,C,D) && ccw(A,B,C) != ccw(A,B,D)
}

func (m Map) Intersect(edge1 edge, edge2 edge) bool {
	a := edge1.From
	b := edge1.To
	c := edge2.From
	d := edge2.To
	//fmt.Println(a, b, c, d)
	if m.Orientation(a, b, c) != m.Orientation(a, b, d) && m.Orientation(c, d, a) != m.Orientation(c, d, b) {
		return true
	}
	return false
}

func (m Map) Orientation( a vertex, b vertex, c vertex) int {
	val := (b.Y-a.Y)*(c.X-b.X) - (b.X-a.X)*(c.Y-b.Y)
	if val == 0 {
		return 0
	}
	if val > 0 {
		return 1
	}
	return 2
}

func (m Map) printMap() {
	fmt.Println(m.StartX, m.StartY)
	for _, row := range m.Cord {
		for _, col := range row {
			switch col.Visited {
				case "7":
					fmt.Print("┐")
			case "J":
				fmt.Print("┘")
			case "L":
				fmt.Print("└")
			case "F":
				fmt.Print("┌")
			case "-":
				fmt.Print("─")
			case "|":
				fmt.Print("│")
			case ".":
				fmt.Print(".")
			case "I":
				fmt.Print("I")

			default:
				if col.Visited == "" {
					fmt.Print("S")
				} else {
					fmt.Print(col.Visited)
				}
			}
		}
		fmt.Println()
	}
}


func(m *Map) FollowPath(x int, y int, px int, py int, length int, curVer *vertex) int {
	//m.Cord[y][x].Char = "*"
	if length>0 && x == m.StartX && y == m.StartY {
		return length
	}


	if x < 0 || y < 0 || x >= len(m.Cord[0]) || y >= len(m.Cord) || m.Cord[y][x].Char == "." {
		return -1
	}





	m.Cord[y][x].Visited = m.Cord[y][x].Char
	fmt.Print(m.Cord[y][x].Char)

	if m.Cord[y][x].Char == "L" || m.Cord[y][x].Char == "F" || m.Cord[y][x].Char == "7" || m.Cord[y][x].Char == "J" {
		if curVer == nil {
			curVer = new(vertex)
			curVer.X = x
			curVer.Y = y
		} else {
			m.EdgeList = append(m.EdgeList, edge{From:*curVer, To:vertex{X:x, Y:y}})
			curVer.X = x
			curVer.Y = y
		}
	}
	switch m.Cord[y][x].Char {
	case "S":
		return length
	case "L":
		if py != y {
			return m.FollowPath(x+1, y, x, y, length+1, curVer)
		} else {
			return m.FollowPath(x, y-1, x, y, length+1, curVer)
		}
	case "|":
		if py < y {
			return m.FollowPath(x, y+1, x, y, length+1, curVer)
		} else {
			return m.FollowPath(x, y-1, x, y, length+1, curVer)
		}
	case "F":
		if px != x {
			return m.FollowPath(x, y+1, x, y, length+1, curVer)
		} else {
			return m.FollowPath(x+1, y, x, y, length+1, curVer)
		}
	case "-":
		if px < x {
			return m.FollowPath(x+1, y, x, y, length+1, curVer)
		} else {
			return m.FollowPath(x-1, y, x, y, length+1, curVer)
		}
	case ".":
		return -1
	case "7":
		if px != x {
			return m.FollowPath(x, y+1, x, y, length+1, curVer)
		} else {
			return m.FollowPath(x-1, y, x, y, length+1, curVer)
		}
	case "J":
		if px != x {
			return m.FollowPath(x, y-1, x, y, length+1, curVer)
		} else {
			return m.FollowPath(x-1, y, x, y, length+1, curVer)
		}


	}
	return -1
}

// main Such a mess this one. (it also only 'mostly' works. I eyeballed a couple of bugs for the count. 
func main() {
	rows := strings.Split(input, "\n")
	length := len(rows)
	width := len(rows[0])
	var rougha, roughb, roughc, roughd Map
	rougha.Cord = make([][]position, length)
	roughb.Cord = make([][]position, length)
	roughc.Cord = make([][]position, length)
	roughd.Cord = make([][]position, length)
	roughd.EdgeList = []edge{}
	for i, val := range rows {
		rougha.Cord[i] = make([]position, width)
		roughb.Cord[i] = make([]position, width)
		roughc.Cord[i] = make([]position, width)
		roughd.Cord[i] = make([]position, width)
		for j, char := range val {
			if string(char) == "S" {
				rougha.StartX, rougha.StartY = j, i
				rougha.X, rougha.Y = j, i
				roughb.StartX, roughb.StartY = j, i
				roughb.X, roughb.Y = j, i
				roughc.StartX, roughc.StartY = j, i
				roughc.X, roughc.Y = j, i
				roughd.StartX, roughd.StartY = j, i
				roughd.X, roughd.Y = j, i
			} else {
				rougha.Cord[i][j] = position{Char: string(char), Visited: "."}
				roughb.Cord[i][j] = position{Char: string(char), Visited: "."}
				roughc.Cord[i][j] = position{Char: string(char), Visited: "."}
				roughd.Cord[i][j] = position{Char: string(char), Visited: "."}
			}
		}
	}

	//rough.printMap()
	//fmt.Println(rough.StartX, rough.StartY, 0)
	fmt.Print("S")
	var a,b,c,d int
	//rightChar := rougha.Cord[rougha.StartY][rougha.StartX+1].Char
	//if  rightChar == "J" || rightChar == "7" || rightChar == "-" {
	//
	//	a = rougha.FollowPath(rougha.StartX+1, rougha.StartY, rougha.StartX, rougha.StartY, 1)
	//}
	//
	//if a > 0 {
	//	fmt.Println()
	//	rougha.printMap()
	//	fmt.Println()
	//}
	//
	//fmt.Println()
	//fmt.Print("S")
	////fmt.Println(rough.StartX, rough.StartY, 0)
	//if roughb.StartX-1 > 0 {
	//	leftChar := roughb.Cord[roughb.StartY][roughb.StartX-1].Char
	//	if leftChar == "F" || leftChar == "L" || leftChar == "-" {
	//
	//		b = roughb.FollowPath(roughb.StartX-1, roughb.StartY, roughb.StartX, roughb.StartY, 1)
	//	}
	//}
	//if b > 0 {
	//	fmt.Println()
	//	roughb.printMap()
	//	fmt.Println()
	//}
	//
	//
	//fmt.Println()
	//fmt.Print("S")
	////fmt.Println(roughc.StartX, roughc.StartY, 0)
	//downChar :=  roughc.Cord[roughc.StartY+1][roughc.StartX].Char
	//if downChar == "J" || downChar == "L" || downChar == "|" {
	//	c = roughc.FollowPath(roughc.StartX, roughc.StartY+1, roughc.StartX, roughc.StartY, 1)
	//}
	//if c > 0 {
	//	fmt.Println()
	//	roughc.printMap()
	//	fmt.Println()
	//}
	fmt.Println()
	fmt.Print("S")
	//fmt.Println(roughd.StartX, roughd.StartY, 0)
	if roughd.StartY-1 > 0 {
		upChar := roughd.Cord[roughd.StartY-1][roughd.StartX].Char
		if upChar == "7" || upChar == "F" || upChar == "|" {

			d = roughd.FollowPath(roughd.StartX, roughd.StartY-1, roughd.StartX, roughd.StartY, 1, nil)
		}
	}
	if d > 0 {
		fmt.Println()
		roughd.printMap()
		fmt.Println()
	}
	fmt.Println()
	fmt.Println(a, b, c, d)
	//roughd.FindVertex()
	fmt.Println(roughd.CheckUnvistedPoints())
	fmt.Println()
	roughd.printMap()
}
