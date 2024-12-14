# frozen_string_literal: true

# Day 14
module Day14
  @rows = 103
  @cols = 101

  def self.part1
    robots_quadrant1 = 0
    robots_quadrant2 = 0
    robots_quadrant3 = 0
    robots_quadrant4 = 0

    File.foreach('day14/input.txt').each do |line|
      matches = line.match(/p=(\d+),(\d+) v=(-?\d+),(-?\d+)/)
      x = matches[1].to_i
      y = matches[2].to_i
      vx = matches[3].to_i
      vy = matches[4].to_i

      x_after_100_seconds = (x + 100 * vx) % @cols
      y_after_100_seconds = (y + 100 * vy) % @rows

      robots_quadrant1 += 1 if x_after_100_seconds < @cols / 2 && y_after_100_seconds < @rows / 2
      robots_quadrant2 += 1 if x_after_100_seconds < @cols / 2 && y_after_100_seconds > @rows / 2
      robots_quadrant3 += 1 if x_after_100_seconds > @cols / 2 && y_after_100_seconds < @rows / 2
      robots_quadrant4 += 1 if x_after_100_seconds > @cols / 2 && y_after_100_seconds > @rows / 2
    end

    puts robots_quadrant1 * robots_quadrant2 * robots_quadrant3 * robots_quadrant4
  end
end
