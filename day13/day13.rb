# frozen_string_literal: true

# Day 13
module Day13
  def self.part1
    file = File.new('day13/input.txt')

    ractors = []
    until file.eof?
      ractors << Ractor.new(read_prize_lines(file)) do |vars|
        puts "Ractor start for #{vars.inspect}"
        start_time = Time.now
        prize_x, prize_y, a_x, a_y, b_x, b_y = vars
        a_presses, b_presses = Day13.amount_of_presses(prize_x, prize_y, a_x, a_y, b_x, b_y)

        elapsed_time = Time.now - start_time
        puts "Ractor end for #{vars.inspect}. Got to #{a_presses}, #{b_presses}, in #{elapsed_time}"
        a_presses * 3 + b_presses
      end
    end

    puts ractors.map(&:take).sum
  end

  def self.read_prize_lines(file)
    button_a = file.readline(chomp: true)
    a_matches = button_a.match(/Button A: X\+(\d+), Y\+(\d+)/)
    a_x = a_matches[1].to_i
    a_y = a_matches[2].to_i

    button_b = file.readline(chomp: true)
    b_matches = button_b.match(/Button B: X\+(\d+), Y\+(\d+)/)
    b_x = b_matches[1].to_i
    b_y = b_matches[2].to_i

    prize = file.readline(chomp: true)
    prize_matches = prize.match(/Prize: X=(\d+), Y=(\d+)/)
    prize_x = prize_matches[1].to_i + 10_000_000_000_000
    prize_y = prize_matches[2].to_i + 10_000_000_000_000

    file.readline unless file.eof?

    [prize_x, prize_y, a_x, a_y, b_x, b_y]
  end

  def self.amount_of_presses(prize_x, prize_y, a_x, a_y, b_x, b_y)
    (0..100_000_000_000).each do |a_presses|
      break if a_presses * a_x > prize_x || a_presses * a_y > prize_y

      b_presses = (prize_x - a_presses * a_x) / b_x
      if a_presses * a_x + b_presses * b_x == prize_x && a_presses * a_y + b_presses * b_y == prize_y
        return [a_presses, b_presses]
      end
    end

    [0, 0]
  end

  private_class_method(:read_prize_lines)
end
