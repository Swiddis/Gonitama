import json
import random

class Deck():
    def __init__(self, filename="cards.json"):
        with open(filename, 'r') as card_file:
            self.cards = json.loads(card_file.read())
    
    def draw(self, count=5):
        sample = random.sample(self.cards, count)
        random.shuffle(sample)
        return sample
    
    def compress(self, cards):
        return [tuple((m["dx"], m["dy"]) for m in c["moves"]) for c in cards]
    
    def get_card(self, name):
        for card in self.cards:
            if card["name"] == name:
                return card

    def __str__(self):
        return f"<Deck of {len(self.cards)} cards>"
