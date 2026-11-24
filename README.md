# Flood Fill

> Striver SDE Sheet 151 · Day 22 Mixed — DFS recoloring with one guard that matters: a no-op color must return before recursing. Go, stdlib only, tested.

## Problem

The paint-bucket tool: starting from a pixel, recolor every 4-connected pixel sharing its original color.

## Approach

Depth-first spread: recolor the pixel, recurse into the four neighbors that still carry the original color. Recoloring *is* the visited-marking — a repainted pixel fails the `== old` test, so no separate visited set exists. The one famous landmine: when the new color equals the old one, nothing fails that test, and the recursion ping-pongs forever — hence the up-front `old == color` return, which the test suite pins explicitly. Grids are graphs; this is DFS wearing pixels, and the sheet places it as the on-ramp to the graph section.

**Complexity:** O(pixels) time, O(region) recursion depth worst case.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
