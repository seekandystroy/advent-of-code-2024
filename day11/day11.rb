# frozen_string_literal: true

# Day 11
module Day11
  def self.part1
    nums = File.new('day11/input.txt').readline(chomp: true).split(' ').map(&:to_i)

    25.times { nums = blink(nums) }

    puts(nums.length)
  end

  def self.part2
    nums = File.new('day11/input.txt').readline(chomp: true).split(' ').map(&:to_i).sort

    cache = {}
    sum = nums.map { |num| blink_with_cache(num, cache, 1) }.sum

    puts(sum)
  end

  def self.blink(nums)
    nums.flat_map { |num| individual_blink(num) }
  end

  def self.blink_with_cache(num, cache, blink_num, path = Set.new)
    total_blinks = 75
    return cache[num][total_blinks - blink_num] if cache.dig(num, total_blinks - blink_num)

    children = individual_blink(num)
    sum = if blink_num < total_blinks
            children.map do |child|
              blink_with_cache(child, cache, blink_num + 1, path.add(num))
            end.sum
          else
            individual_blink(num).length
          end

    cache[num] = cache[num] || {}
    cache[num][total_blinks - blink_num] = sum

    sum
  end

  def self.individual_blink(num)
    if num == 0
      [1]
    elsif even_digits?(num)
      half_zeroes = 10.pow(num_digits(num) / 2)
      first_half = num / half_zeroes
      second_half = num - first_half * half_zeroes

      [first_half, second_half]
    else
      [num * 2024]
    end
  end

  def self.num_digits(num)
    Math.log10(num).floor + 1
  end

  def self.even_digits?(num)
    num_digits(num).even?
  end

  private_class_method(:blink)
  private_class_method(:blink_with_cache)
  private_class_method(:individual_blink)
  private_class_method(:num_digits)
  private_class_method(:even_digits?)
end
