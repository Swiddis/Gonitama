from itertools import product
from termcolor import colored
import mcts
import random

RED_KING = -2
RED = -1
BLUE_KING = 2
BLUE = 1
EMPTY = 0
DEFAULT_BOARD = [
    [BLUE, BLUE, BLUE_KING, BLUE, BLUE],
    [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
    [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
    [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
    [RED, RED, RED_KING, RED, RED]
]

class Onitama():
    def __init__(
        self,
        cards,
        turn=BLUE,
        board=DEFAULT_BOARD,

    ):
        self.board = board
        self.cards = cards
        self.turn = turn


    def get_current_player(self):
        return self.turn
    
    def _get_current_pieces(self):
        return {
            (i, j)
            for i, row in enumerate(self.board)
            for j, col in enumerate(row)
            if self.board[i][j] == self.turn
            or self.board[i][j] == 2 * self.turn
        }

    def get_possible_actions(self):
        moves = []
        cards = self.cards[1:3] if self.turn == BLUE else self.cards[3:5]
        pieces = self._get_current_pieces()
        for card, piece in product(cards, pieces):
            for move in card["moves"]:
                mx = -self.turn * move["dx"] + piece[1]
                if mx < 0 or mx >= 5:
                    continue
                my = self.turn * move["dy"] + piece[0]
                if my < 0 or my >= 5:
                    continue
                if (my, mx) in pieces:
                    continue
                moves.append((card["name"], piece, (my, mx)))
        return moves

    def take_action(self, card, piece, destination):
        card_idx = 0
        for i, c in enumerate(self.cards):
            if c["name"] == card:
                card_idx = i
                break
        cards = [card for card in self.cards]
        cards[0], cards[card_idx] = cards[card_idx], cards[0]
        board = [[i for i in row] for row in self.board]
        board[destination[0]][destination[1]] = board[piece[0]][piece[1]]
        board[piece[0]][piece[1]] = EMPTY
        return Onitama(cards, -self.turn, board)

    def is_terminal(self):
        if self.board[0][2] == RED_KING or self.board[4][2] == BLUE_KING:
            return True
        if not any(RED_KING in row for row in self.board):
            return True
        if not any(BLUE_KING in row for row in self.board):
            return True
        return False

    def get_reward(self):
        if self.board[0][2] == RED_KING:
            return RED
        if self.board[4][2] == BLUE_KING:
            return BLUE
        if not any(BLUE_KING in row for row in self.board):
            return RED
        if not any(RED_KING in row for row in self.board):
            return BLUE
        return 0
    
    def __hash__(self):
        return hash((self.turn, *(card["name"] for card in self.cards), *(tuple(row) for row in self.board)))

    def __str__(self):
        result = ""
        result += f"{colored('holding', 'white')}: {self.cards[0]['name']}\n"
        result += f"{colored('blue', 'blue')}: {self.cards[1]['name']}, {self.cards[2]['name']}\n"
        result += " ABCDE\n"
        for i, row in enumerate(self.board, 1):
            result += f"{i}"
            for col in row:
                if col == RED:
                    result += colored("p", "red")
                elif col == BLUE:
                    result += colored("P", "blue")
                elif col == RED_KING:
                    result += colored("k", "red")
                elif col == BLUE_KING:
                    result += colored("K", "blue")
                else:
                    result += colored(".", "grey")
            result += f"\n"
        result += f"{colored('red', 'red')}: {self.cards[3]['name']}, {self.cards[4]['name']}"
        return result
    
    def __eq__(self, state):
        return self.turn == state.turn and self.board == state.board and self.cards == state.cards

class OnitamaMCTS(mcts.Node):
    def __init__(self, node, target):
        self.node = node
        self.target = target

    def find_children(self):
        return [
            OnitamaMCTS(self.node.take_action(*move), self.target)
            for move in self.node.get_possible_actions()
        ]
    
    def find_random_child(self):
        try:
            return random.choice(self.find_children())
        except:
            return None
    
    def is_terminal(self):
        return self.node.is_terminal()
    
    def reward(self):
        return 1 if self.node.get_reward() == self.target else 0
    
    def __hash__(self):
        return hash(self.node)
    
    def __eq__(self, node):
        self.node == node.node
    
    def __str__(self):
        return str(self.node)
