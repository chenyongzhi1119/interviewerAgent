package extract

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode"

	"github.com/ledongthuc/pdf"
)

// PDF extracts plain text from a base64-encoded PDF file.
func PDF(base64Data string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		raw, err = base64.RawStdEncoding.DecodeString(base64Data)
		if err != nil {
			return "", fmt.Errorf("base64 decode: %w", err)
		}
	}

	r := bytes.NewReader(raw)
	pdfReader, err := pdf.NewReader(r, int64(len(raw)))
	if err != nil {
		return "", fmt.Errorf("parse pdf: %w", err)
	}

	var pages []string
	for i := 1; i <= pdfReader.NumPage(); i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}
		text, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}
		pages = append(pages, text)
	}

	joined := strings.Join(pages, "\n\n")
	// Phase 1: join lines that were split by the PDF renderer's column width
	phase1 := rejoinBrokenLines(joined)
	// Phase 2: re-insert structural breaks (section headers, bullets, numbered items)
	result := strings.TrimSpace(restoreStructure(phase1))
	if result == "" {
		return "", fmt.Errorf("PDF 未提取到文字（可能是扫描版图片 PDF，请改用图片上传）")
	}
	return result, nil
}

// ─── Phase 1: join soft-wrapped lines ────────────────────────────────────────

// rejoinBrokenLines merges lines that were split only by the PDF column width.
// Strategy: only blank lines and explicit numbered list items cause hard breaks;
// everything else is joined to recover acronyms, short words, and mid-sentence wraps.
func rejoinBrokenLines(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	lines := strings.Split(text, "\n")

	var out []string
	var cur strings.Builder

	flush := func() {
		if s := strings.TrimSpace(cur.String()); s != "" {
			out = append(out, s)
		}
		cur.Reset()
	}

	for _, raw := range lines {
		trimmed := strings.TrimSpace(raw)
		if trimmed == "" {
			flush()
			out = append(out, "")
			continue
		}
		if cur.Len() == 0 {
			cur.WriteString(trimmed)
			continue
		}
		prev := cur.String()
		if shouldBreak(prev, trimmed) {
			flush()
			cur.WriteString(trimmed)
		} else {
			if needsSpace(prev, trimmed) {
				cur.WriteByte(' ')
			}
			cur.WriteString(trimmed)
		}
	}
	flush()

	var sb strings.Builder
	prevBlank := false
	for _, ln := range out {
		if ln == "" {
			if !prevBlank {
				sb.WriteByte('\n')
			}
			prevBlank = true
		} else {
			sb.WriteString(ln)
			sb.WriteByte('\n')
			prevBlank = false
		}
	}
	return strings.TrimSpace(sb.String())
}

// shouldBreak: only two hard signals trigger a new paragraph.
func shouldBreak(prev, next string) bool {
	if prev == "" || next == "" {
		return true
	}
	if isNumberedListItem(next) {
		return true
	}
	if isLongAllCapsHeading(prev) || isLongAllCapsHeading(next) {
		return true
	}
	return false
}

// isNumberedListItem: "1. text", "2、text", "① text" — needs ≥3 chars.
func isNumberedListItem(s string) bool {
	runes := []rune(s)
	if len(runes) < 3 {
		return false
	}
	first := runes[0]
	if first >= '①' && first <= '⑳' {
		return true
	}
	if first >= '1' && first <= '9' {
		second := runes[1]
		if second == '.' || second == '、' || second == ')' || second == '）' {
			return true
		}
	}
	if len(runes) >= 4 && (first == '(' || first == '（') {
		if runes[1] >= '1' && runes[1] <= '9' {
			return true
		}
	}
	return false
}

// isLongAllCapsHeading: "EDUCATION", "EXPERIENCE" — ≥6 letters to avoid acronyms.
func isLongAllCapsHeading(s string) bool {
	runes := []rune(strings.TrimSpace(s))
	letterCount := 0
	for _, r := range runes {
		if unicode.IsLetter(r) {
			if !unicode.IsUpper(r) {
				return false
			}
			letterCount++
		} else if !unicode.IsSpace(r) {
			return false
		}
	}
	return letterCount >= 6
}

func needsSpace(prev, next string) bool {
	if prev == "" || next == "" {
		return false
	}
	prevRunes := []rune(prev)
	last := prevRunes[len(prevRunes)-1]
	first := []rune(next)[0]
	if last == '.' && isListNumberDotPrefix(prev) {
		return true
	}
	return isASCIIWord(last) && isASCIIWord(first)
}

func isListNumberDotPrefix(s string) bool {
	runes := []rune(s)
	n := len(runes)
	if n < 2 || runes[n-1] != '.' {
		return false
	}
	for _, r := range runes[:n-1] {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func isASCIIWord(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

// ─── Phase 2: restore structural breaks ──────────────────────────────────────

// chineseResumeHeaders are section titles commonly found in Chinese resumes.
// When these appear glued to surrounding content, we insert blank lines around them.
var chineseResumeHeaders = []string{
	"教育背景", "教育经历",
	"专业技能", "技能特长", "核心技能", "技术技能", "技术栈",
	"工作经历", "实习经历", "工作经验",
	"项目经历", "项目经验",
	"个人信息", "基本信息",
	"自我评价", "个人评价",
	"获奖情况", "荣誉奖项", "证书与荣誉",
	"科研成果", "发表论文", "学术成果",
	"社会实践", "课外活动",
}

// restoreStructure re-inserts newlines that semantic structure requires:
//  1. Blank lines around Chinese section headers
//  2. Newlines before bullet points (·)
//  3. Newlines before dash-items ("- Chinese…") that are resume separators
//  4. Newlines before numbered list items embedded in running text
func restoreStructure(text string) string {
	// 1. Section headers
	for _, h := range chineseResumeHeaders {
		text = strings.ReplaceAll(text, h, "\n\n"+h+"\n")
	}

	// 2. Bullet points ·: each must start on its own line
	runes := []rune(text)
	var sb strings.Builder
	for i, r := range runes {
		if r == '·' && i > 0 && runes[i-1] != '\n' {
			sb.WriteByte('\n')
		}
		sb.WriteRune(r)
	}
	text = sb.String()

	// 3. Dash items: "-" followed by CJK (not digits, not space, not another dash)
	// Avoids breaking date ranges like 2023-09 or negative numbers.
	runes = []rune(text)
	sb.Reset()
	for i, r := range runes {
		if r == '-' && i > 0 && runes[i-1] != '\n' && runes[i-1] != '-' && i+1 < len(runes) {
			next := runes[i+1]
			if isCJK(next) {
				sb.WriteByte('\n')
			}
		}
		sb.WriteRune(r)
	}
	text = sb.String()

	// 4. Numbered list items embedded in running text
	// "...句子。1.下一条..." → "...句子。\n1.下一条..."
	text = splitEmbeddedNumberedItems(text)

	// Collapse 3+ blank lines → 2
	for strings.Contains(text, "\n\n\n") {
		text = strings.ReplaceAll(text, "\n\n\n", "\n\n")
	}

	return strings.TrimSpace(text)
}

// splitEmbeddedNumberedItems inserts a newline before numbered items that appear
// in the middle of text (i.e., not already at the start of a line).
// "...段落内容。1.条目内容..." → "...段落内容。\n1.条目内容..."
func splitEmbeddedNumberedItems(text string) string {
	runes := []rune(text)
	var sb strings.Builder
	n := len(runes)

	for i := 0; i < n; i++ {
		r := runes[i]

		// Look for a digit that starts a list item
		if r >= '1' && r <= '9' && i > 0 {
			prev := runes[i-1]
			// prev must not already be a newline, digit, or dash (dates like 2023-09)
			if prev != '\n' && !(prev >= '0' && prev <= '9') && prev != '-' && prev != '~' {
				// Collect consecutive digits
				j := i
				for j < n && runes[j] >= '0' && runes[j] <= '9' {
					j++
				}
				// Must be followed by "." or "、"
				if j < n && (runes[j] == '.' || runes[j] == '、') {
					// After the punct must NOT be a digit (to avoid "3.14")
					afterPunct := j + 1
					if afterPunct >= n || !(runes[afterPunct] >= '0' && runes[afterPunct] <= '9') {
						sb.WriteByte('\n')
					}
				}
			}
		}

		sb.WriteRune(r)
	}
	return sb.String()
}

// isCJK reports whether r is a CJK (Chinese/Japanese/Korean) character.
func isCJK(r rune) bool {
	return (r >= 0x4E00 && r <= 0x9FFF) || // CJK Unified Ideographs
		(r >= 0x3400 && r <= 0x4DBF) || // Extension A
		(r >= 0xF900 && r <= 0xFAFF) // Compatibility
}
