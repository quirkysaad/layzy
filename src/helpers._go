package main

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

// View engine helpers

// --- Overlay Compositor Function ---
// Composites the `fg` (modal) directly on top of `bg` (main view) at coordinates (x, y)
func placeOverlay(x, y int, fg, bg string) string {
	fgLines := strings.Split(fg, "\n")
	bgLines := strings.Split(bg, "\n")

	for i, fgLine := range fgLines {
		bgIdx := y + i
		if bgIdx < 0 || bgIdx >= len(bgLines) {
			continue
		}

		bgLine := bgLines[bgIdx]
		fgWidth := ansi.StringWidth(fgLine)
		bgWidth := ansi.StringWidth(bgLine)

		// 1. Slice the visual left part of the background before the modal
		var left string
		if x > 0 {
			if x > bgWidth {
				left = bgLine + strings.Repeat(" ", x-bgWidth)
			} else {
				left = ansi.Truncate(bgLine, x, "")
			}
		}

		// 2. Slice the visual right part of the background after the modal
		var right string
		rightStart := x + fgWidth
		if rightStart < bgWidth {
			right = ansi.TruncateLeft(bgLine, rightStart, "")
		}

		// 3. Stitch left background + modal line + right background together
		bgLines[bgIdx] = left + fgLine + right
	}

	return strings.Join(bgLines, "\n")
}

// --- Backdrop Dimming Helper ---
// Strips all bright ANSI colors from the background and paints it a uniform dark gray
// to simulate a dark translucent backdrop over the inactive UI.
func dimBackground(bg string) string {
	dimStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")). // Muted gray text/borders
		Faint(true)                            // Terminal faint/dim attribute

	lines := strings.Split(bg, "\n")
	for i, line := range lines {
		// Strip existing bright colors and apply the dimmed style
		cleanLine := ansi.Strip(line)
		lines[i] = dimStyle.Render(cleanLine)
	}
	return strings.Join(lines, "\n")
}
