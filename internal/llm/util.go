package llm

import (
	"fmt"
	"io"
	"net/http"
)

// sseEscape ensures newlines in LLM output don't break SSE framing.
func sseEscape(s string) string {
	out := ""
	for _, r := range s {
		if r == '\n' {
			out += "\ndata: "
		} else {
			out += string(r)
		}
	}
	return out
}

func flush(w io.Writer) {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

// WriteError sends an SSE error line to the client.
func WriteError(w io.Writer, err error) {
	fmt.Fprintf(w, "data: [ERROR] %s\n\n", err.Error())
}
