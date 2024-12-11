# frozen_string_literal: true

require_relative('day10/day10')
require_relative('day11/day11')

case ARGV[0]
when '1001'
  Day10.part1
when '1002'
  Day10.part2
when '1101'
  Day11.part1
else
  puts('Please choose a day and part to run, in the format DDPP.')
end
