# frozen_string_literal: true

require 'matrix'

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

  def self.part2
    robots = []
    File.foreach('day14/input.txt').each do |line|
      matches = line.match(/p=(\d+),(\d+) v=(-?\d+),(-?\d+)/)
      x = matches[1].to_i
      y = matches[2].to_i
      vx = matches[3].to_i
      vy = matches[4].to_i

      robots.append([x, y, vx, vy])
    end

    File.open('day14/total_out.txt', 'w') do |file|
      (1001..9999).each do |i|
        file.puts "Second #{i}"
        matrix = Matrix.build(@rows, @cols) { 0 }
        robots.each do |robot|
          x = robot[0]
          y = robot[1]
          vx = robot[2]
          vy = robot[3]

          x_after_i_seconds = (x + i * vx) % @cols
          y_after_i_seconds = (y + i * vy) % @rows

          matrix[y_after_i_seconds, x_after_i_seconds] += 1
        end

        matrix.to_a.each do |row|
          file.puts row.map { |num| num.zero? ? ' ' : num }.join
        end
      end
    end
  end
end
