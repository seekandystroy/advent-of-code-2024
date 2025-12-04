# frozen_string_literal: true

require_relative('day10/day10')
require_relative('day11/day11')
require_relative('day12/day12')
require_relative('day13/day13')
require_relative('day14/day14')

puts case ARGV[0]
     when '1001'
       Day10.part1
     when '1002'
       Day10.part2
     when '1101'
       Day11.part1
     when '1102'
       Day11.part2
     when '1201'
       Day12.part1
     when '1202'
       Day12.part2
     when '1301'
       Day13.part1
     when '1302'
       Day13.part2
     when '1401'
       Day14.part1
     when '1402'
       Day14.part2
     else
       'Please choose a day and part to run, in the format DDPP.'
     end
