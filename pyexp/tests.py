import unittest
import deck
import onitama
from onitama import RED, BLUE, RED_KING, BLUE_KING, EMPTY

class TestMoveGeneration(unittest.TestCase):
    def test_blue_crane_tiger(self):
        test_deck = deck.Deck()
        test_cards = ["rat", "crane", "tiger", "crab", "rabbit"]
        cards = [test_deck.get_card(name) for name in test_cards]
        board = onitama.Onitama(cards)
        self.assertEqual(
            set(board.get_possible_actions()),
            {
                ("crane", (0, 0), (1, 0)), ("tiger", (0, 0), (2, 0)),
                ("crane", (0, 1), (1, 1)), ("tiger", (0, 1), (2, 1)),
                ("crane", (0, 2), (1, 2)), ("tiger", (0, 2), (2, 2)),
                ("crane", (0, 3), (1, 3)), ("tiger", (0, 3), (2, 3)),
                ("crane", (0, 4), (1, 4)), ("tiger", (0, 4), (2, 4)),
            }
        )
    
    def test_red_crab_rabbit(self):
        test_deck = deck.Deck()
        test_cards = ["rat", "crane", "tiger", "crab", "rabbit"]
        cards = [test_deck.get_card(name) for name in test_cards]
        board = onitama.Onitama(cards, turn=onitama.RED)
        self.assertEqual(
            set(board.get_possible_actions()),
            {
                ("crab", (4, 0), (3, 0)), ("rabbit", (4, 0), (3, 1)),
                ("crab", (4, 1), (3, 1)), ("rabbit", (4, 1), (3, 2)),
                ("crab", (4, 2), (3, 2)), ("rabbit", (4, 2), (3, 3)),
                ("crab", (4, 3), (3, 3)), ("rabbit", (4, 3), (3, 4)),
                ("crab", (4, 4), (3, 4))
            }
        )
    
    def test_blue_sable_viper(self):
        test_deck = deck.Deck()
        test_cards = ["dog", "sable", "viper", "elephant", "goat"]
        cards = [test_deck.get_card(name) for name in test_cards]
        board = onitama.Onitama(cards, board=[
            [EMPTY for _ in range(5)],
            [EMPTY for _ in range(5)],
            [EMPTY, EMPTY, BLUE_KING, EMPTY, EMPTY],
            [EMPTY for _ in range(5)],
            [EMPTY for _ in range(4)] + [RED_KING]
        ])
        self.assertEqual(
            set(board.get_possible_actions()),
            {
                ("sable", (2, 2), (3, 1)), ("sable", (2, 2), (2, 4)),
                ("sable", (2, 2), (1, 3)), ("viper", (2, 2), (1, 1)),
                ("viper", (2, 2), (3, 2)), ("viper", (2, 2), (2, 4))
            }
        )
    
    def test_red_elephant_goat(self):
        test_deck = deck.Deck()
        test_cards = ["dog", "sable", "viper", "elephant", "goat"]
        cards = [test_deck.get_card(name) for name in test_cards]
        board = onitama.Onitama(cards, turn=RED, board=[
            [EMPTY for _ in range(5)],
            [EMPTY for _ in range(5)],
            [EMPTY, EMPTY, RED_KING, EMPTY, EMPTY],
            [EMPTY for _ in range(5)],
            [EMPTY for _ in range(4)] + [BLUE_KING]
        ])
        self.assertEqual(
            set(board.get_possible_actions()),
            {
                ("elephant", (2, 2), (2, 3)), ("elephant", (2, 2), (2, 1)),
                ("elephant", (2, 2), (1, 3)), ("elephant", (2, 2), (1, 1)),
                ("goat", (2, 2), (2, 1)), ("goat", (2, 2), (3, 2)),
                ("goat", (2, 2), (1, 3))
            }
        )

    def test_realgame_blue(self):
        test_deck = deck.Deck()
        test_cards = ["frog", "tanuki", "horse", "sheep", "boar"]
        cards = [test_deck.get_card(name) for name in test_cards]
        board = onitama.Onitama(cards, turn=BLUE, board=[
            [BLUE, EMPTY, EMPTY, BLUE, BLUE],
            [BLUE_KING, EMPTY, EMPTY, EMPTY, EMPTY],
            [EMPTY, RED, BLUE, EMPTY, EMPTY],
            [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
            [RED, RED, RED_KING, EMPTY, RED]
        ])
        self.assertEqual(
            set(board.get_possible_actions()),
            {
                ("horse", (0, 0), (0, 1)),
                ("horse", (0, 3), (1, 3)),
                ("horse", (0, 4), (1, 4)),
                ("horse", (1, 0), (2, 0)),
                ("horse", (1, 0), (1, 1)),
                ("horse", (2, 2), (1, 2)),
                ("horse", (2, 2), (3, 2)),
                ("horse", (2, 2), (2, 3)),
                ('tanuki', (0, 4), (1, 4)),
                ('tanuki', (0, 3), (1, 3)),
                ('tanuki', (1, 0), (0, 1)),
                ('tanuki', (1, 0), (2, 0)),
                ('tanuki', (2, 2), (1, 3)),
                ('tanuki', (0, 3), (1, 1)),
                ('tanuki', (0, 4), (1, 2)),
                ('tanuki', (2, 2), (3, 0)),
                ('tanuki', (2, 2), (3, 2)),
            }
        )
