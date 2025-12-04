# frozen_string_literal: true

# Day 15
module Day15
  def self.part1
    file = File.new('day15/input.txt')
    warehouse = warehouse_from_file(file)
    moves = file.gets.chomp.chars
    row, col = starting_pos(warehouse)

    # move method mutates warehouse and position
    moves.each { |move| row, col = move(warehouse, row, col, move) }

    score_boxes(warehouse)
  end

  def self.warehouse_from_file(file)
    warehouse = []

    line = file.gets.chomp
    while line != ''
      warehouse << line.chars
      line = file.gets.chomp
    end

    warehouse
  end

  def self.starting_pos(warehouse)
    row = 0
    col = 0

    while row < warehouse.size
      col = 0
      while col < warehouse[row].size
        return [row, col] if warehouse[row][col] == '@'

        col += 1
      end
      row += 1
    end

    [row, col]
  end

  def self.move(warehouse, row, col, move)
    case move
    when '>'
      move_right(warehouse, row, col)
    when '<'
      move_left(warehouse, row, col)
    when '^'
      move_up(warehouse, row, col)
    when 'v'
      move_down(warehouse, row, col)
    end
  end

  def self.move_right(warehouse, row, col)
    edge = warehouse[row].size - 1
    return [row, col] if col == edge || warehouse[row][col + 1] == '#'

    move_boxes_right(warehouse, row, col) if warehouse[row][col + 1] == 'O'

    if warehouse[row][col + 1] == '.'
      warehouse[row][col] = '.'
      warehouse[row][col + 1] = '@'

      [row, col + 1]
    else
      [row, col]
    end
  end

  def self.move_boxes_right(warehouse, row, col)
    edge = warehouse[row].size - 1
    right_elems = warehouse[row].slice((col + 1)..edge)

    empty_idx = right_elems.find_index('.') || edge
    wall_idx = right_elems.find_index('#')
    return unless empty_idx < wall_idx

    warehouse[row][col + 1 + empty_idx] = 'O'
    warehouse[row][col + 1] = '.'
  end

  def self.move_left(warehouse, row, col)
    return [row, col] if col == 0 || warehouse[row][col - 1] == '#'

    move_boxes_left(warehouse, row, col) if warehouse[row][col - 1] == 'O'

    if warehouse[row][col - 1] == '.'
      warehouse[row][col] = '.'
      warehouse[row][col - 1] = '@'

      [row, col - 1]
    else
      [row, col]
    end
  end

  def self.move_boxes_left(warehouse, row, col)
    left_elems = warehouse[row].slice(0..(col - 1)).reverse
    lidx = left_elems.size - 1

    empty_idx = lidx - (left_elems.find_index('.') || (left_elems.size - 1))
    wall_idx = lidx - left_elems.find_index('#')
    return unless empty_idx > wall_idx

    warehouse[row][empty_idx] = 'O'
    warehouse[row][col - 1] = '.'
  end

  def self.move_up(warehouse, row, col)
    return [row, col] if row == 0 || warehouse[row - 1][col] == '#'

    move_boxes_up(warehouse, row, col) if warehouse[row - 1][col] == 'O'

    if warehouse[row - 1][col] == '.'
      warehouse[row][col] = '.'
      warehouse[row - 1][col] = '@'

      [row - 1, col]
    else
      [row, col]
    end
  end

  def self.move_boxes_up(warehouse, row, col)
    up_elems = warehouse.slice(0..(row - 1)).flat_map { |row_arr| row_arr[col] }.reverse
    uidx = up_elems.size - 1

    empty_idx = uidx - (up_elems.find_index('.') || uidx)
    wall_idx = uidx - up_elems.find_index('#')
    return unless empty_idx > wall_idx

    warehouse[empty_idx][col] = 'O'
    warehouse[row - 1][col] = '.'
  end

  def self.move_down(warehouse, row, col)
    edge = warehouse.size - 1
    return [row, col] if row == edge || warehouse[row + 1][col] == '#'

    move_boxes_down(warehouse, row, col) if warehouse[row + 1][col] == 'O'

    if warehouse[row + 1][col] == '.'
      warehouse[row][col] = '.'
      warehouse[row + 1][col] = '@'

      [row + 1, col]
    else
      [row, col]
    end
  end

  def self.move_boxes_down(warehouse, row, col)
    edge = warehouse.size - 1
    up_elems = warehouse.slice((row + 1)..edge).flat_map { |row_arr| row_arr[col] }

    empty_idx = up_elems.find_index('.') || edge
    wall_idx = up_elems.find_index('#')
    return unless empty_idx < wall_idx

    warehouse[row + 1 + empty_idx][col] = 'O'
    warehouse[row + 1][col] = '.'
  end

  def self.score_boxes(warehouse)
    row = 0
    col = 0
    score = 0

    while row < warehouse.size
      col = 0
      while col < warehouse[row].size
        score += (100 * row) + col if warehouse[row][col] == 'O'

        col += 1
      end
      row += 1
    end

    score
  end

  private_class_method :warehouse_from_file
  private_class_method :starting_pos
  private_class_method :move
  private_class_method :score_boxes
end
