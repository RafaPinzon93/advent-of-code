import argparse
import os.path

import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), "input.txt")


def compute(s: str) -> int:
    left_list = []
    right_list = []
    for line in s.splitlines():
        left, right = map(int, line.split())
        left_list.append(left)
        right_list.append(right)

    total = 0
    left_list.sort()
    right_list.sort()
    for left, right in zip(left_list, right_list):
        total += abs(left - right)

    return total


INPUT_S = """\
3   4
4   3
2   5
1   3
3   9
3   3
"""
EXPECTED = 11


@pytest.mark.parametrize(
    ("input_s", "expected"),
    ((INPUT_S, EXPECTED),),
)
def test(input_s: str, expected: int) -> None:
    assert compute(input_s) == expected


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("data_file", nargs="?", default=INPUT_TXT)
    args = parser.parse_args()

    with open(args.data_file) as f:
        print(compute(f.read()))

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
