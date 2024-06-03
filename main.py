import cProfile
import pstats
import io

N = 16

def print_board(board):
    for i in range(N):
        for j in range(N):
            if board[i][j] == 1:
                print("Q ", end="")
            else:
                print(". ", end="")
        print()
    print("-" * (2 * N))

def is_safe(board, row, col):
    for i in range(col):
        if board[row][i] == 1:
            return False

    i, j = row, col
    while i >= 0 and j >= 0:
        if board[i][j] == 1:
            return False
        i -= 1
        j -= 1

    i, j = row, col
    while i < N and j >= 0:
        if board[i][j] == 1:
            return False
        i += 1
        j -= 1

    return True

def solve_nq_util(board, col):
    if col >= N:
        return True

    for i in range(N):
        if is_safe(board, i, col):
            board[i][col] = 1
            if solve_nq_util(board, col + 1):
                return True
            board[i][col] = 0

    return False

def solve_nq():
    board = [[0] * N for _ in range(N)]
    if not solve_nq_util(board, 0):
        print("Solution does not exist")
        return

# Profile the solve_nq function
pr = cProfile.Profile()
pr.enable()
solve_nq()
pr.disable()

s = io.StringIO()
sortby = 'cumulative'
ps = pstats.Stats(pr, stream=s).sort_stats(sortby)
ps.print_stats()
print(s.getvalue())
