//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	goprinter "go/printer"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

type docSet struct {
	Types   map[string]*typeDoc
	Funcs   []*funcDoc
	Methods map[string][]*funcDoc
	Consts  map[string][]*constDoc
	Vars    map[string]*varDoc
}

type typeDoc struct {
	Name    string
	Kind    string
	Doc     string
	Fields  []fieldDoc
	File    string
	Line    int
	Aliases string
}

type fieldDoc struct {
	Name string
	Type string
	Doc  string
	JSON string
}

type funcDoc struct {
	Name     string
	Recv     string
	Doc      string
	Params   []fieldDoc
	Results  []fieldDoc
	Sign     string
	File     string
	Line     int
	Category string
}

type constDoc struct {
	Name  string
	Type  string
	Value string
	Doc   string
}

type varDoc struct {
	Name  string
	Value string
	Doc   string
}

type pageSpec struct {
	Title      string
	File       string
	Summary    string
	Types      []string
	Functions  []string
	MethodRecv []string
}

func main() {
	docs, err := collectDocs(".")
	if err != nil {
		fail(err)
	}
	if err := os.MkdirAll("docs/api", 0o755); err != nil {
		fail(err)
	}
	for _, page := range pages() {
		var out bytes.Buffer
		renderPage(&out, docs, page)
		if err := os.WriteFile(filepath.Join("docs/api", page.File), out.Bytes(), 0o644); err != nil {
			fail(err)
		}
	}
}

func collectDocs(root string) (*docSet, error) {
	set := &docSet{
		Types:   map[string]*typeDoc{},
		Methods: map[string][]*funcDoc{},
		Consts:  map[string][]*constDoc{},
		Vars:    map[string]*varDoc{},
	}
	files, err := filepath.Glob(filepath.Join(root, "*.go"))
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	for _, file := range files {
		base := filepath.Base(file)
		if strings.HasSuffix(base, "_test.go") {
			continue
		}
		parsed, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		for _, decl := range parsed.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				switch d.Tok {
				case token.TYPE:
					for _, spec := range d.Specs {
						ts, ok := spec.(*ast.TypeSpec)
						if !ok || !ts.Name.IsExported() {
							continue
						}
						td := &typeDoc{
							Name: ts.Name.Name,
							Doc:  cleanDoc(firstDoc(ts.Doc, d.Doc)),
							File: strings.TrimPrefix(file, "./"),
							Line: fset.Position(ts.Pos()).Line,
						}
						switch t := ts.Type.(type) {
						case *ast.StructType:
							td.Kind = "struct"
							td.Fields = parseFields(t)
						case *ast.Ident:
							td.Kind = "type"
							td.Aliases = t.Name
						default:
							td.Kind = "type"
							td.Aliases = exprString(t)
						}
						set.Types[td.Name] = td
					}
				case token.CONST:
					groupType := ""
					for _, spec := range d.Specs {
						vs, ok := spec.(*ast.ValueSpec)
						if !ok {
							continue
						}
						if vs.Type != nil {
							groupType = exprString(vs.Type)
						}
						for i, name := range vs.Names {
							if !name.IsExported() {
								continue
							}
							value := ""
							if i < len(vs.Values) {
								value = exprString(vs.Values[i])
							}
							set.Consts[groupType] = append(set.Consts[groupType], &constDoc{
								Name:  name.Name,
								Type:  groupType,
								Value: value,
								Doc:   cleanDoc(firstDoc(vs.Doc, d.Doc)),
							})
						}
					}
				case token.VAR:
					for _, spec := range d.Specs {
						vs, ok := spec.(*ast.ValueSpec)
						if !ok {
							continue
						}
						for i, name := range vs.Names {
							if !name.IsExported() {
								continue
							}
							value := ""
							if i < len(vs.Values) {
								value = exprString(vs.Values[i])
							}
							set.Vars[name.Name] = &varDoc{Name: name.Name, Value: value, Doc: cleanDoc(firstDoc(vs.Doc, d.Doc))}
						}
					}
				}
			case *ast.FuncDecl:
				if !d.Name.IsExported() {
					continue
				}
				fd := &funcDoc{
					Name:    d.Name.Name,
					Doc:     cleanDoc(d.Doc),
					Params:  parseFieldList(d.Type.Params),
					Results: parseFieldList(d.Type.Results),
					File:    strings.TrimPrefix(file, "./"),
					Line:    fset.Position(d.Pos()).Line,
				}
				fd.Sign = signature(fd)
				if d.Recv != nil && len(d.Recv.List) > 0 {
					fd.Recv = receiverName(d.Recv.List[0].Type)
					set.Methods[fd.Recv] = append(set.Methods[fd.Recv], fd)
				} else {
					set.Funcs = append(set.Funcs, fd)
				}
			}
		}
	}
	for _, methods := range set.Methods {
		sort.Slice(methods, func(i, j int) bool { return methods[i].Name < methods[j].Name })
	}
	sort.Slice(set.Funcs, func(i, j int) bool { return set.Funcs[i].Name < set.Funcs[j].Name })
	return set, nil
}

func parseFields(st *ast.StructType) []fieldDoc {
	var fields []fieldDoc
	for _, field := range st.Fields.List {
		names := field.Names
		if len(names) == 0 {
			if ident := embeddedName(field.Type); ident != "" && ast.IsExported(ident) {
				names = []*ast.Ident{{Name: ident}}
			} else {
				continue
			}
		}
		for _, name := range names {
			if !name.IsExported() {
				continue
			}
			fields = append(fields, fieldDoc{
				Name: name.Name,
				Type: exprString(field.Type),
				Doc:  cleanDoc(field.Doc),
				JSON: jsonName(field.Tag),
			})
		}
	}
	return fields
}

func parseFieldList(list *ast.FieldList) []fieldDoc {
	if list == nil {
		return nil
	}
	var fields []fieldDoc
	for _, field := range list.List {
		typ := exprString(field.Type)
		if len(field.Names) == 0 {
			fields = append(fields, fieldDoc{Type: typ})
			continue
		}
		for _, name := range field.Names {
			fields = append(fields, fieldDoc{Name: name.Name, Type: typ})
		}
	}
	return fields
}

func renderPage(out *bytes.Buffer, docs *docSet, page pageSpec) {
	fmt.Fprintf(out, "# %s\n\n", page.Title)
	if page.Summary != "" {
		fmt.Fprintf(out, "%s\n\n", page.Summary)
	}
	if page.File == "constants.md" {
		renderVars(out, docs)
	}
	for _, name := range page.Types {
		td := docs.Types[name]
		if td == nil {
			continue
		}
		renderType(out, docs, page.File, td)
	}
	for _, recv := range page.MethodRecv {
		if methods := docs.Methods[recv]; len(methods) > 0 {
			fmt.Fprintf(out, "## %s Methods\n\n", splitName(recv))
			for _, method := range methods {
				renderFunc(out, docs, page.File, method)
			}
		}
	}
	if len(page.Functions) > 0 {
		fmt.Fprintln(out, "## Functions\n")
		for _, name := range page.Functions {
			for _, fn := range docs.Funcs {
				if fn.Name == name {
					renderFunc(out, docs, page.File, fn)
				}
			}
		}
	}
}

func renderVars(out *bytes.Buffer, docs *docSet) {
	names := make([]string, 0, len(docs.Vars))
	for name := range docs.Vars {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		v := docs.Vars[name]
		fmt.Fprintf(out, "<div class=\"api-field\" id=\"%s\" markdown=\"1\">\n\n", anchor(v.Name))
		fmt.Fprintf(out, "### `%s`\n\n", v.Name)
		if v.Value != "" {
			fmt.Fprintf(out, "<p><code>%s</code></p>\n\n", v.Value)
		}
		if v.Doc != "" {
			fmt.Fprintf(out, "%s\n\n", v.Doc)
		}
		fmt.Fprintln(out, "</div>\n")
	}
}

func renderType(out *bytes.Buffer, docs *docSet, currentPage string, td *typeDoc) {
	fmt.Fprintf(out, "<a id=\"%s\"></a>\n\n", anchor(td.Name))
	fmt.Fprintf(out, "## %s\n\n", splitName(td.Name))
	fmt.Fprintf(out, "<p class=\"api-signature\"><span class=\"api-kind\">%s</span> <code>clashy.%s</code></p>\n\n", td.Kind, td.Name)
	if td.Doc != "" {
		fmt.Fprintf(out, "%s\n\n", td.Doc)
	}
	if td.Kind != "struct" {
		if td.Aliases != "" {
			fmt.Fprintf(out, "<p><code>%s</code></p>\n\n", linkType(docs, currentPage, td.Aliases))
		}
		renderConsts(out, docs, td.Name)
		return
	}
	for _, field := range td.Fields {
		renderField(out, docs, currentPage, td.Name, field)
	}
	renderConsts(out, docs, td.Name)
}

func renderField(out *bytes.Buffer, docs *docSet, currentPage, parent string, field fieldDoc) {
	fmt.Fprintf(out, "<div class=\"api-field\" id=\"%s\" markdown=\"1\">\n\n", anchor(parent+"-"+field.Name))
	fmt.Fprintf(out, "### `%s`\n\n", field.Name)
	fmt.Fprintf(out, "<p><code>%s</code>", linkType(docs, currentPage, field.Type))
	if field.JSON != "" && field.JSON != "-" {
		fmt.Fprintf(out, " <span class=\"api-json\">json: %s</span>", field.JSON)
	}
	fmt.Fprintln(out, "</p>\n")
	if field.Doc != "" {
		fmt.Fprintf(out, "%s\n\n", field.Doc)
	}
	fmt.Fprintln(out, "</div>\n")
}

func renderConsts(out *bytes.Buffer, docs *docSet, typeName string) {
	consts := docs.Consts[typeName]
	if len(consts) == 0 {
		return
	}
	fmt.Fprintln(out, "### Values\n")
	for _, c := range consts {
		fmt.Fprintf(out, "<div class=\"api-field\" id=\"%s\" markdown=\"1\">\n\n", anchor(c.Name))
		fmt.Fprintf(out, "#### `%s`\n\n", c.Name)
		if c.Value != "" {
			fmt.Fprintf(out, "<p><code>%s</code></p>\n\n", c.Value)
		}
		if c.Doc != "" {
			fmt.Fprintf(out, "%s\n\n", c.Doc)
		}
		fmt.Fprintln(out, "</div>\n")
	}
}

func renderFunc(out *bytes.Buffer, docs *docSet, currentPage string, fn *funcDoc) {
	label := fn.Name
	if fn.Recv != "" {
		label = fn.Recv + "." + fn.Name
	}
	fmt.Fprintf(out, "<a id=\"%s\"></a>\n\n", anchor(label))
	fmt.Fprintf(out, "<div class=\"api-function\" markdown=\"1\">\n\n")
	fmt.Fprintf(out, "<p class=\"api-signature api-function-signature\"><code>%s</code></p>\n\n", functionDisplay(fn))
	if fn.Doc != "" {
		fmt.Fprintf(out, "%s\n\n", fn.Doc)
	}
	if len(fn.Params) > 0 {
		fmt.Fprintln(out, "<dl class=\"api-parameters\">")
		fmt.Fprintln(out, "<dt>Parameters:</dt><dd>")
		for _, p := range fn.Params {
			name := p.Name
			if name == "" {
				name = "value"
			}
			fmt.Fprintf(out, "<p><strong>%s</strong> (<code>%s</code>)</p>\n", name, linkType(docs, currentPage, p.Type))
		}
		fmt.Fprintln(out, "</dd>")
		fmt.Fprintln(out, "</dl>\n")
	}
	if len(fn.Results) > 0 {
		hasNamedResults := false
		for _, r := range fn.Results {
			if r.Name != "" {
				hasNamedResults = true
				break
			}
		}
		if hasNamedResults {
			fmt.Fprintln(out, "<dl class=\"api-parameters\">")
			fmt.Fprintln(out, "<dt>Returns:</dt><dd>")
			for _, r := range fn.Results {
				name := r.Name
				if name == "" {
					name = "value"
				}
				fmt.Fprintf(out, "<p><strong>%s</strong> (<code>%s</code>)</p>\n", name, linkType(docs, currentPage, r.Type))
			}
			fmt.Fprintln(out, "</dd>")
			fmt.Fprintln(out, "</dl>\n")
		}
		fmt.Fprintln(out, "<dl class=\"api-parameters\">")
		fmt.Fprintln(out, "<dt>Return type:</dt><dd>")
		for _, r := range fn.Results {
			fmt.Fprintf(out, "<code>%s</code> ", linkType(docs, currentPage, r.Type))
		}
		fmt.Fprintln(out, "</dd>")
		fmt.Fprintln(out, "</dl>\n")
	}
	fmt.Fprintln(out, "</div>\n")
}

func pages() []pageSpec {
	return []pageSpec{
		{
			Title:      "Client",
			File:       "client.md",
			Summary:    "Client construction, authentication, request behavior, and high-level Clash API methods.",
			Types:      []string{"Client", "ClientConfig", "SearchClansRequest", "RequestOptions", "HTTPClient"},
			Functions:  []string{"NewClient", "DefaultClientConfig", "NewHTTPClient"},
			MethodRecv: []string{"Client", "HTTPClient"},
		},
		{
			Title:      "Clans",
			File:       "clans.md",
			Summary:    "Clan profiles, clan members, clan labels, and clan search models.",
			Types:      []string{"Clan", "ClanMember", "PlayerClan", "ClanCapital", "ClanType"},
			MethodRecv: []string{"Clan"},
		},
		{
			Title:      "Players",
			File:       "players.md",
			Summary:    "Player profile models and helpers for achievements, units, spells, and labels.",
			Types:      []string{"Player", "LegendStatistics", "Achievement", "PlayerHouseElement"},
			MethodRecv: []string{"Player"},
		},
		{
			Title:      "Wars",
			File:       "wars.md",
			Summary:    "Classic war and Clan War League response models.",
			Types:      []string{"ClanWar", "WarClan", "ClanWarMember", "WarAttack", "ClanWarLogEntry", "ClanWarLeagueGroup", "ClanWarLeagueClan", "ExtendedCWLGroup", "WarRound", "WarState", "WarResult"},
			MethodRecv: []string{"ClanWar"},
		},
		{
			Title:      "Raids",
			File:       "raids.md",
			Summary:    "Clan Capital raid weekend logs, districts, attacks, and member totals.",
			Types:      []string{"RaidLogEntry", "RaidClan", "RaidDistrict", "RaidAttack", "RaidMember", "CapitalDistrict"},
			MethodRecv: []string{"RaidLogEntry", "RaidClan"},
		},
		{
			Title:   "Battle Logs",
			File:    "battle-logs.md",
			Summary: "Player battle logs and legend league group records.",
			Types:   []string{"BattleLogEntry", "Resource", "LeagueHistoryEntry", "LeagueTierGroup", "LeagueTierGroupMember", "LeagueTierGroupBattleLogEntry"},
		},
		{
			Title:      "Game Objects",
			File:       "game-objects.md",
			Summary:    "Troops, spells, heroes, pets, equipment, and parsed army-link models.",
			Types:      []string{"StaticUnit", "Troop", "Spell", "Hero", "Pet", "Equipment", "HeroLoadout", "ArmyRecipe", "TroopCount", "SpellCount", "AccountData"},
			Functions:  []string{"ParseArmyRecipe", "ParseAccountData"},
			MethodRecv: []string{"Troop", "Spell", "Hero", "Pet", "Equipment"},
		},
		{
			Title:     "Static Data",
			File:      "static-data.md",
			Summary:   "Embedded ClashKing static data, translations, and lookup helpers.",
			Types:     []string{"StaticData", "Translation"},
			Functions: []string{"LoadStaticData"},
		},
		{
			Title:   "Locations And Rankings",
			File:    "locations-rankings.md",
			Summary: "Locations, leagues, seasons, labels, and ranked clan/player wrappers.",
			Types:   []string{"Location", "League", "Season", "Label", "Icon", "Badge", "GoldPassSeason", "RankedClan", "RankedPlayer"},
		},
		{
			Title:   "Constants",
			File:    "constants.md",
			Summary: "Static-data IDs and display-order slices generated from ClashKing assets.",
		},
		{
			Title:   "Enums",
			File:    "enums.md",
			Summary: "Enum-like string and integer types used across the client.",
			Types:   []string{"Role", "VillageType", "LoadGameData", "ClanType", "WarRound", "WarState", "WarResult"},
		},
		{
			Title:      "Miscellaneous",
			File:       "miscellaneous.md",
			Summary:    "Shared timestamp, season, date, and utility helpers.",
			Types:      []string{"Timestamp", "TimeDelta", "SeasonWindow", "ChatLanguage"},
			Functions:  []string{"CorrectTag", "FromTimestamp", "GetSeasonID", "GenSeasonDate", "GenLegendDate", "GetSeasonStart", "GetSeasonEnd", "GetSeason", "GetSeasonByID", "GetClanGamesStart", "GetClanGamesEnd", "GetRaidWeekendStart", "GetRaidWeekendEnd", "WithoutRateLimit"},
			MethodRecv: []string{"Timestamp"},
		},
		{
			Title: "Exceptions",
			File:  "exceptions.md",
			Types: []string{"ClashOfClansException", "HTTPException", "InvalidArgument", "InvalidCredentials", "Forbidden", "PrivateWarLog", "NotFound", "Maintenance", "GatewayError"},
		},
	}
}

func linkType(docs *docSet, currentPage, typ string) string {
	if typ == "" {
		return ""
	}
	prefix, base, suffix := splitType(typ)
	if _, ok := docs.Types[base]; ok {
		return prefix + fmt.Sprintf(`<a href="%s">%s</a>`, typeHref(currentPage, base), base) + suffix
	}
	return typ
}

func typeHref(currentPage, name string) string {
	page := pageForType(name)
	target := anchor(name)
	if page == currentPage {
		return "#" + target
	}
	slug := strings.TrimSuffix(page, ".md")
	return "../" + slug + "/#" + target
}

func pageForType(name string) string {
	for _, page := range pages() {
		for _, typ := range page.Types {
			if typ == name {
				return page.File
			}
		}
	}
	return "clashy.md"
}

func splitType(typ string) (string, string, string) {
	prefix := ""
	for strings.HasPrefix(typ, "*") || strings.HasPrefix(typ, "[]") {
		if strings.HasPrefix(typ, "*") {
			prefix += "*"
			typ = strings.TrimPrefix(typ, "*")
		}
		if strings.HasPrefix(typ, "[]") {
			prefix += "[]"
			typ = strings.TrimPrefix(typ, "[]")
		}
	}
	base := typ
	suffix := ""
	if i := strings.IndexAny(typ, "[ "); i >= 0 {
		base = typ[:i]
		suffix = typ[i:]
	}
	return prefix, base, suffix
}

func exprString(expr ast.Expr) string {
	var out bytes.Buffer
	_ = goprinter.Fprint(&out, token.NewFileSet(), expr)
	return out.String()
}

func firstDoc(groups ...*ast.CommentGroup) *ast.CommentGroup {
	for _, group := range groups {
		if group != nil {
			return group
		}
	}
	return nil
}

func cleanDoc(group *ast.CommentGroup) string {
	if group == nil {
		return ""
	}
	return strings.TrimSpace(group.Text())
}

func jsonName(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}
	raw := strings.Trim(tag.Value, "`")
	for _, part := range strings.Split(raw, " ") {
		if strings.HasPrefix(part, "json:") {
			value := strings.Trim(strings.TrimPrefix(part, "json:"), `"`)
			return strings.Split(value, ",")[0]
		}
	}
	return ""
}

func embeddedName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return embeddedName(t.X)
	}
	return ""
}

func receiverName(expr ast.Expr) string {
	return embeddedName(expr)
}

func signature(fn *funcDoc) string {
	var out strings.Builder
	out.WriteString(fn.Name)
	out.WriteByte('(')
	for i, p := range fn.Params {
		if i > 0 {
			out.WriteString(", ")
		}
		if p.Name != "" {
			out.WriteString(p.Name)
			out.WriteByte(' ')
		}
		out.WriteString(p.Type)
	}
	out.WriteByte(')')
	if len(fn.Results) == 1 {
		out.WriteByte(' ')
		out.WriteString(fn.Results[0].Type)
	}
	if len(fn.Results) > 1 {
		out.WriteString(" (")
		for i, r := range fn.Results {
			if i > 0 {
				out.WriteString(", ")
			}
			if r.Name != "" {
				out.WriteString(r.Name)
				out.WriteByte(' ')
			}
			out.WriteString(r.Type)
		}
		out.WriteByte(')')
	}
	return out.String()
}

func functionDisplay(fn *funcDoc) string {
	var out strings.Builder
	out.WriteString("clashy.")
	if fn.Recv != "" {
		out.WriteString(fn.Recv)
		out.WriteByte('.')
	}
	out.WriteString(fn.Name)
	out.WriteByte('(')
	for i, p := range fn.Params {
		if i > 0 {
			out.WriteString(", ")
		}
		if p.Name != "" {
			out.WriteString(`<span class="api-param">`)
			out.WriteString(p.Name)
			out.WriteString(": ")
			out.WriteString(p.Type)
			out.WriteString(`</span>`)
		} else {
			out.WriteString(`<span class="api-param">`)
			out.WriteString(p.Type)
			out.WriteString(`</span>`)
		}
	}
	out.WriteByte(')')
	if len(fn.Results) > 0 {
		out.WriteString(`<span class="api-return-arrow"> -> </span>`)
		if len(fn.Results) == 1 {
			out.WriteString(`<span class="api-return">`)
			out.WriteString(fn.Results[0].Type)
			out.WriteString(`</span>`)
		} else {
			out.WriteByte('(')
			for i, r := range fn.Results {
				if i > 0 {
					out.WriteString(", ")
				}
				out.WriteString(`<span class="api-return">`)
				out.WriteString(r.Type)
				out.WriteString(`</span>`)
			}
			out.WriteByte(')')
		}
	}
	return out.String()
}

func anchor(name string) string {
	name = strings.ReplaceAll(name, ".", "-")
	return strings.ToLower(name)
}

func splitName(name string) string {
	var out []rune
	for i, r := range name {
		if i > 0 && unicode.IsUpper(r) {
			prev := rune(name[i-1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) {
				out = append(out, ' ')
			}
		}
		out = append(out, r)
	}
	return string(out)
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
