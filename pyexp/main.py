import deck
import onitama
import random
from tqdm import tqdm

onitama_deck = deck.Deck()
# print(onitama_deck)

cards = onitama_deck.draw()
# print(cards)
game = onitama.Onitama(cards)

posmap = {
    "a": 0, "b": 1, "c": 2, "d": 3, "e": 4,
    '1': 0, '2': 1, '3': 2, '4': 3, '5': 4
}

t = tqdm()
while True:
    board = onitama.Onitama(cards)
    while not board.is_terminal():
        try:
            board = board.take_action(*random.choice(board.get_possible_actions()))
            t.update()
        except:
            pass
