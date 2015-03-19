package main
import (
"fmt"
"log"
"os"
"text/template"
)


func main() {
// Define a template.
const letter = `
Dear {{.Honorific}}{{.LastName}},
{{if .Attended}}
It was a pleasure to see you at the fundraiser.{{else}}
It is a shame you could not attend the fundraiser.{{end}}
{{if .Donated}}
Thank you for your donation.{{end}}
{{range .UpcommingEvents}}
Remember to come to the next event: {{.}}
{{end}}
Best wishes,
Jimmy
`

// Prepare some data to insert into the template.
type Recipient struct {
Honorific, LastName string
Attended, Donated bool
UpcommingEvents []string
}
var upcommingEvents = []string{"movie watching",
"Dinner",
"baby shower"}
var recipients = []Recipient{
{"Mr", " Leong", true, true, upcommingEvents},
{"Mrs", " Valdivia", false, true, upcommingEvents},
{"Mr", " Soo", true, false, upcommingEvents},
{"Ms", " Wilson", false, false, upcommingEvents},
}
// Create a new template and parse the letter into it.
t := template.Must(template.New("letter").Parse(letter))
// Execute the template for each recipient.
for _, r := range recipients {
err := t.Execute(os.Stdout, r)
if err != nil {
log.Println("executing template:", err)
}
}
fmt.Println("")
