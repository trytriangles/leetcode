class Position:
    def __init__(self, row: int, col: int):
        self.row = row
        self.col = col

    def __repr__(self) -> str:
        return f"{self.row}.{self.col}"

    def is_valid(self, board_size: int) -> bool:
        if self.row < 0:
            return False
        if self.row >= board_size:
            return False
        if self.col < 0:
            return False
        if self.col >= board_size:
            return False
        return True

    def copy(self):
        return Position(self.row, self.col)

    def move_northeast(self):
        self.row -= 1
        self.col += 1

    def move_southeast(self):
        self.row += 1
        self.col += 1

    def move_southwest(self):
        self.row += 1
        self.col -= 1

    def move_northwest(self):
        self.row -= 1
        self.col -= 1

    def move_north(self):
        self.row -= 1

    def move_south(self):
        self.row += 1

    def move_west(self):
        self.col -= 1

    def move_east(self):
        self.col += 1


class Board:
    def __init__(self, size: int):
        self.size = size
        self.board: list[str] = [[0 for _ in range(size)] for _ in range(size)]
        self.queens_placed: list[Position] = []
        self.next_row = 0
        self.next_col = 0
        self.solutions: list[str] = []

    def threatened_spaces(self) -> set[Position]:
        result: set[Position] = set()
        for row in range(self.size):
            for col in range(self.size):
                if self.board[row][col] == 1:
                    self.board[row][col] = 0
                elif self.board[row][col] == 2:
                    attackables = self.find_attackable_positions(Position(row, col))
                    for a in attackables:
                        result.add(a)
        return result

    def find_attackable_positions(self, p: Position) -> list[Position]:
        result: list[Position] = []
        for row in range(self.size):
            for col in range(self.size):
                if row == p.row and col == p.col:
                    continue
                if col == p.col:
                    result.append(Position(row, col))
                elif row == p.row:
                    result.append(Position(row, col))
        # NW
        x = Position(p.row - 1, p.col - 1)
        while x.is_valid(self.size):
            result.append(x.copy())
            x.move_northwest()
        # SE
        x.row = p.row + 1
        x.col = p.col + 1
        while x.is_valid(self.size):
            result.append(x.copy())
            x.move_southeast()
        # NE
        x.row = p.row - 1
        x.col = p.col + 1
        while x.is_valid(self.size):
            result.append(x.copy())
            x.move_northeast()
        # SW
        x.row = p.row + 1
        x.col = p.col - 1
        while x.is_valid(self.size):
            result.append(x.copy())
            x.move_southwest()
        return result

    def highlight_position(self, p: Position, num: int = 1) -> None:
        self.board[p.row][p.col] = num

    def highlight_attackable_positions(self, pos: Position):
        self.highlight_position(pos, 2)
        positions = self.find_attackable_positions(pos)
        for p in positions:
            self.highlight_position(p)

    def highlight_threatened_spaces(self):
        for p in self.threatened_spaces():
            if self.board[p.row][p.col] == 0:
                self.board[p.row][p.col] = 1

    def is_occupied(self, p: Position) -> bool:
        return self.board[p.row][p.col] != 0

    def space_is_safe(self, p: Position) -> bool:
        attackables = self.find_attackable_positions(p)
        for a in attackables:
            if self.is_occupied(a):
                return False
        return True

    @property
    def string_form(self) -> list[str]:
        result: list[str] = []
        for row in self.board:
            result.append("".join("Q" if x == 2 else "." for x in row))
        return result

    def backtrack(self) -> bool:
        if len(self.queens_placed) == 0:
            return False
        last_queen = self.queens_placed.pop()
        self.board[last_queen.row][last_queen.col] = 0
        self.highlight_threatened_spaces()
        self.next_row = last_queen.row
        self.next_col = last_queen.col + 1
        if self.next_col == self.size:
            self.next_col = 0
            self.next_row += 1
            if self.next_row >= len(self.queens_placed):
                self.backtrack()
            if self.next_row == self.size:
                self.backtrack()
        return True

    @property
    def current_space(self) -> int:
        return self.board[self.next_row][self.next_col]

    @current_space.setter
    def current_space(self, value: int):
        self.board[self.next_row][self.next_col] = value

    def solve(self) -> list[str]:
        if self.size == 2 or self.size == 3:
            return []
        while self.next_row < self.size:
            if self.next_col == self.size:
                self.backtrack()
            elif self.current_space == 0:
                self.current_space = 2
                self.highlight_threatened_spaces()
                self.queens_placed.append(Position(self.next_row, self.next_col))
                if len(self.queens_placed) == self.size:
                    self.solutions.append(self.string_form)
                    if self.backtrack():
                        pass
                else:
                    self.next_row += 1
                    self.next_col = 0
            else:
                self.next_col += 1
        return self.solutions


class Solution:
    def solveNQueens(self, n: int) -> list[list[str]]:
        return Board(n).solve()
