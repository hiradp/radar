package radar

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Host struct {
	Name      string
	Endpoints []Endpoint
}

func (h *Host) HTML() string {
	tmpl := template.Must(template.ParseFiles("assets/default.tmpl"))

	var b bytes.Buffer
	if err := tmpl.Execute(&b, h); err != nil {
		log.Fatalln("ERROR - failed to render template", err)
	}

	return b.String()
}

func (h *Host) String() string {
	data := make([][]string, len(h.Endpoints))

	for i, e := range h.Endpoints {
		data[i] = []string{
			e.IPAddress,
			e.ServerName,
			e.Grade,
			e.Cert.String(),
		}
	}

	b := &strings.Builder{}
	table := tablewriter.NewWriter(b)

	table.SetHeader([]string{"IP Address", "Server Name", "Grade", "SSL Cert"})
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()

	return fmt.Sprintf("Host: %s\n%s", h.Name, b.String())
}
