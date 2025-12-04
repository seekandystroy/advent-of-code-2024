# frozen_string_literal: true

# Day 12
module Day12
  def self.part1
    matrix = File.foreach('day12/input.txt').map { |line| line.strip.chars.prepend(-1).append(-1) }
    matrix.prepend(Array.new(matrix[0].length, -1))
    matrix.append(Array.new(matrix[0].length, -1))

    visited = matrix.map { |row| row.map { |_| false } }
    row = 1
    col = 1
    sum = 0

    while row < matrix.length - 1
      perimeter, area, new_row, new_col = traverse_for_perimeter(matrix, row, col, 0, 0, visited)

      # puts("Perimeter #{perimeter}; Area #{area}; New: #{new_row}, #{new_col}")

      sum += perimeter * area
      row = new_row
      col = new_col
    end

    sum
  end

  def self.part2
    matrix = File.foreach('day12/input.txt').map { |line| line.strip.chars.prepend(-1).append(-1) }
    matrix.prepend(Array.new(matrix[0].length, -1))
    matrix.append(Array.new(matrix[0].length, -1))

    visited = matrix.map { |row| row.map { |_| false } }
    row = 1
    col = 1
    sum = 0

    while row < matrix.length - 1
      sides, area, new_row, new_col = traverse_for_sides(matrix, row, col, 0, 0, visited)

      # puts("Area * Sides: #{area} * #{sides}, #{matrix[row][col]}")

      sum += sides * area
      row = new_row
      col = new_col
    end

    sum
  end

  def self.traverse_for_perimeter(matrix, row, col, perimeter, area, visited)
    symbol = matrix[row][col]
    visited[row][col] = true # visited
    # puts("Visiting #{row}, #{col}")
    area += 1

    if matrix[row][col + 1] == symbol && !visited[row][col + 1]
      perimeter, area = traverse_for_perimeter(matrix, row, col + 1, perimeter, area, visited)
    elsif matrix[row][col + 1] != symbol
      perimeter += 1
    end

    if matrix[row + 1][col] == symbol && !visited[row + 1][col]
      perimeter, area = traverse_for_perimeter(matrix, row + 1, col, perimeter, area, visited)
    elsif matrix[row + 1][col] != symbol
      perimeter += 1
    end

    if matrix[row][col - 1] == symbol && !visited[row][col - 1]
      perimeter, area = traverse_for_perimeter(matrix, row, col - 1, perimeter, area, visited)
    elsif matrix[row][col - 1] != symbol
      perimeter += 1
    end

    if matrix[row - 1][col] == symbol && !visited[row - 1][col]
      perimeter, area = traverse_for_perimeter(matrix, row - 1, col, perimeter, area, visited)
    elsif matrix[row - 1][col] != symbol
      perimeter += 1
    end

    new_row, new_col = find_next(visited, row)

    [perimeter, area, new_row, new_col]
  end

  def self.traverse_for_sides(matrix, row, col, sides, area, visited)
    symbol = matrix[row][col]
    visited[row][col] = true # visited
    # puts("Visiting #{row}, #{col}")
    area += 1
    sides += sides(matrix, row, col, visited)

    if matrix[row][col + 1] == symbol && !visited[row][col + 1]
      sides, area = traverse_for_sides(matrix, row, col + 1, sides, area, visited)
    end

    if matrix[row + 1][col] == symbol && !visited[row + 1][col]
      sides, area = traverse_for_sides(matrix, row + 1, col, sides, area, visited)
    end

    if matrix[row][col - 1] == symbol && !visited[row][col - 1]
      sides, area = traverse_for_sides(matrix, row, col - 1, sides, area, visited)
    end

    if matrix[row - 1][col] == symbol && !visited[row - 1][col]
      sides, area = traverse_for_sides(matrix, row - 1, col, sides, area, visited)
    end

    new_row, new_col = find_next(visited, row)

    [sides, area, new_row, new_col]
  end

  def self.find_next(visited, row)
    (row..visited.length - 2).each do |r|
      (1..visited.length - 2).each do |c|
        return [r, c] unless visited[r][c]
      end
    end

    [visited.length - 1, visited.length - 1]
  end

  def self.sides(matrix, row, col, visited)
    symbol = matrix[row][col]
    sides = 0
    sides += 1 if matrix[row][col + 1] != symbol
    sides += 1 if matrix[row + 1][col] != symbol
    sides += 1 if matrix[row][col - 1] != symbol
    sides += 1 if matrix[row - 1][col] != symbol

    # remove top side
    if matrix[row - 1][col] != symbol &&
       ((matrix[row][col - 1] == symbol && visited[row][col - 1] && matrix[row - 1][col - 1] != symbol) ||
       (matrix[row][col + 1] == symbol && visited[row][col + 1] && matrix[row - 1][col + 1] != symbol))
      sides -= 1
    end
    # remove bottom side
    if matrix[row + 1][col] != symbol &&
       ((matrix[row][col - 1] == symbol && visited[row][col - 1] && matrix[row + 1][col - 1] != symbol) ||
        (matrix[row][col + 1] == symbol && visited[row][col + 1] && matrix[row + 1][col + 1] != symbol))
      sides -= 1
    end
    # remove left side
    if matrix[row][col - 1] != symbol &&
       ((matrix[row - 1][col] == symbol && visited[row - 1][col] && matrix[row - 1][col - 1] != symbol) ||
        (matrix[row + 1][col] == symbol && visited[row + 1][col] && matrix[row + 1][col - 1] != symbol))
      sides -= 1
    end
    # remove right side
    if matrix[row][col + 1] != symbol &&
       ((matrix[row - 1][col] == symbol && visited[row - 1][col] && matrix[row - 1][col + 1] != symbol) ||
        (matrix[row + 1][col] == symbol && visited[row + 1][col] && matrix[row + 1][col + 1] != symbol))
      sides -= 1
    end

    # at the end of traversing a straight line, we'll have an extra side due to starting that side from the opposite direction
    # when that line meets the initial point, we remove that extra side
    if matrix[row - 1][col] != symbol &&
       ((matrix[row][col - 1] == symbol && visited[row][col - 1] && matrix[row - 1][col - 1] != symbol) &&
        (matrix[row][col + 1] == symbol && visited[row][col + 1] && matrix[row - 1][col + 1] != symbol))
      sides -= 1
    end
    if matrix[row + 1][col] != symbol &&
       ((matrix[row][col - 1] == symbol && visited[row][col - 1] && matrix[row + 1][col - 1] != symbol) &&
        (matrix[row][col + 1] == symbol && visited[row][col + 1] && matrix[row + 1][col + 1] != symbol))
      sides -= 1
    end
    if matrix[row][col - 1] != symbol &&
       ((matrix[row - 1][col] == symbol && visited[row - 1][col] && matrix[row - 1][col - 1] != symbol) &&
        (matrix[row + 1][col] == symbol && visited[row + 1][col] && matrix[row + 1][col - 1] != symbol))
      sides -= 1
    end
    if matrix[row][col + 1] != symbol &&
       ((matrix[row - 1][col] == symbol && visited[row - 1][col] && matrix[row - 1][col + 1] != symbol) &&
        (matrix[row + 1][col] == symbol && visited[row + 1][col] && matrix[row + 1][col + 1] != symbol))
      sides -= 1
    end

    sides
  end

  private_class_method(:traverse_for_perimeter)
  private_class_method(:traverse_for_sides)
  private_class_method(:find_next)
end
