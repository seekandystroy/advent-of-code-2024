# frozen_string_literal: true

require_relative('day10/day10')

case ARGV[0]
when '1001'
  Day10.part1
when '1002'
  Day10.part2
else
  puts('Please choose a day and part to run, in the format DDPP.')
end
