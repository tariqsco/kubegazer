"""Airbnb-style interview prep: four common coding problems."""

from __future__ import annotations

from urllib.parse import unquote_plus


def parse_query(url_str: str) -> dict:
    """
    Parses a URL or query string into a dictionary.

    - Treats '+' and '%20' as spaces.
    - Collects duplicate keys into a list of values.
    - Handles keys with no value as empty strings.
    """
    query = url_str.split("?", 1)[-1]
    if "#" in query:
        query = query.split("#", 1)[0]
    if not query:
        return {}

    parsed: dict = {}
    for segment in query.split("&"):
        if not segment:
            continue
        if "=" in segment:
            raw_key, _, raw_value = segment.partition("=")
        else:
            raw_key, raw_value = segment, ""

        key = unquote_plus(raw_key)
        value = unquote_plus(raw_value)

        if key not in parsed:
            parsed[key] = value
        else:
            existing = parsed[key]
            if isinstance(existing, list):
                existing.append(value)
            else:
                parsed[key] = [existing, value]

    return parsed


def find_routes(
    flights: list[tuple],
    start: str,
    end: str,
    min_layover: int,
    max_layover: int,
) -> list[list[tuple]]:
    """
    Finds all valid flight itineraries from start to end.

    Each flight is (origin, destination, dep_time, arr_time).
    Valid layover: min_layover <= (next_dep - prev_arr) <= max_layover.
    """
    routes: list[list[tuple]] = []

    def dfs(city: str, path: list[tuple], prev_arr: int | None, used: set[int]) -> None:
        if city == end and path:
            routes.append(path.copy())
            return

        for i, flight in enumerate(flights):
            if i in used:
                continue
            origin, destination, dep, arr = flight
            if origin != city:
                continue
            if path:
                layover = dep - prev_arr
                if layover < min_layover or layover > max_layover:
                    continue

            used.add(i)
            path.append(flight)
            dfs(destination, path, arr, used)
            path.pop()
            used.remove(i)

    dfs(start, [], None, set())
    return routes


def highlight_keywords(text: str, keywords: list[str]) -> str:
    """
    Encloses matching keywords in <b> and </b> tags.

    Overlapping or adjacent matches are merged into a single <b>...</b> block.
    """
    intervals: list[list[int]] = []
    for keyword in keywords:
        if not keyword:
            continue
        start = 0
        while True:
            idx = text.find(keyword, start)
            if idx == -1:
                break
            intervals.append([idx, idx + len(keyword)])
            start = idx + 1

    if not intervals:
        return text

    intervals.sort()
    merged = [intervals[0]]
    for start, end in intervals[1:]:
        last_start, last_end = merged[-1]
        if start <= last_end:
            merged[-1][1] = max(last_end, end)
        else:
            merged.append([start, end])

    parts: list[str] = []
    cursor = 0
    for start, end in merged:
        parts.append(text[cursor:start])
        parts.append("<b>")
        parts.append(text[start:end])
        parts.append("</b>")
        cursor = end
    parts.append(text[cursor:])
    return "".join(parts)


class MapSum:
    """Sum values for all keys sharing a prefix (insert overwrites)."""

    def __init__(self) -> None:
        self._values: dict[str, int] = {}

    def insert(self, key: str, val: int) -> None:
        self._values[key] = val

    def sum(self, prefix: str) -> int:
        return sum(v for k, v in self._values.items() if k.startswith(prefix))
