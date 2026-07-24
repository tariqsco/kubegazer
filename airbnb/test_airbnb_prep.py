"""Tests for airbnb_prep.py."""

import unittest

from airbnb_prep import MapSum, find_routes, highlight_keywords, parse_query


class TestParseQuery(unittest.TestCase):
    def test_basic_url(self):
        url = "https://airbnb.com/search?city=San+Francisco&guests=2"
        self.assertEqual(
            parse_query(url),
            {"city": "San Francisco", "guests": "2"},
        )

    def test_duplicates_encoding_and_bare_key(self):
        query = "amenities=wifi&amenities=pool&location=New%20York&instant_book"
        self.assertEqual(
            parse_query(query),
            {
                "amenities": ["wifi", "pool"],
                "location": "New York",
                "instant_book": "",
            },
        )


class TestFindRoutes(unittest.TestCase):
    def test_valid_layover(self):
        flights = [
            ("SFO", "ORD", 100, 300),
            ("ORD", "JFK", 360, 500),
            ("ORD", "JFK", 600, 740),
        ]
        expected = [[("SFO", "ORD", 100, 300), ("ORD", "JFK", 360, 500)]]
        self.assertEqual(find_routes(flights, "SFO", "JFK", 30, 120), expected)

    def test_layover_too_short(self):
        flights = [
            ("SFO", "ORD", 100, 300),
            ("ORD", "JFK", 310, 450),
        ]
        self.assertEqual(find_routes(flights, "SFO", "JFK", 30, 120), [])


class TestHighlightKeywords(unittest.TestCase):
    def test_overlapping_keywords(self):
        text = "Cozy airbnb listing with superhost status"
        keywords = ["air", "airbnb", "superhost"]
        expected = "Cozy <b>airbnb</b> listing with <b>superhost</b> status"
        self.assertEqual(highlight_keywords(text, keywords), expected)

    def test_adjacent_keywords(self):
        text = "abcdef"
        keywords = ["abc", "def"]
        self.assertEqual(highlight_keywords(text, keywords), "<b>abcdef</b>")


class TestMapSum(unittest.TestCase):
    def test_insert_update_and_prefix_sum(self):
        kv = MapSum()
        kv.insert("apple", 3)
        self.assertEqual(kv.sum("ap"), 3)

        kv.insert("app", 2)
        self.assertEqual(kv.sum("ap"), 5)

        kv.insert("apple", 5)
        self.assertEqual(kv.sum("ap"), 7)


if __name__ == "__main__":
    unittest.main()
