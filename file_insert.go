package imitative_typing

import (
	"encoding/json"
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
	"sort"
)

func RegisterFileInsert() {
	javaScriptContext.PushGoFunction(fmt.Sprintf("it.%s.%s",
		imitativeTypingConfig.GetJavaScriptTableNameForFile(),
		imitativeTypingConfig.GetJavaScriptFunctionNameForFileInsert()),
		func(duktapeContext *duktape.Context) int {
			from := duktapeContext.RequireString(0)
			if FileOrigin(from) {
				panic(fmt.Sprintf("not insert file from a origin file(%s)", from))
			}
			to := duktapeContext.RequireString(1)
			if FileOrigin(to) {
				panic(fmt.Sprintf("not insert file to a origin file(%s)", to))
			}
			idsJson := ""
			if duktapeContext.GetTop() > 2 {
				idsJson = duktapeContext.RequireString(2)
			}
			FileInsert(from, to, idsJson, duktapeContext)
			return 1
		})
}

func FileInsert(from, to, idsJson string, duktapeContext *duktape.Context) {
	ids := make([]string, 0)
	fileContext := GetOrCreateFileContext(from)
	if idsJson == "" {
		idForSortStruct := make([]*IDForSortStruct, 0)
		for id, order := range fileContext.JavaScriptOrder {
			idForSortStruct = append(idForSortStruct, &IDForSortStruct{
				id:    id,
				order: order,
			})
		}
		sort.Sort(IDForSort(idForSortStruct))
		for _, value := range idForSortStruct {
			ids = append(ids, value.id)
		}
	} else {
		err := json.Unmarshal([]byte(idsJson), &ids)
		if err != nil {
			panic(err)
		}
	}

	for _, id := range ids {
		javaScript, exist := fileContext.JavaScript[id]
		if !exist {
			javaScript, exist = imitativeTypingContext.JavaScript[id]
			if !exist {
				javaScript = ""
			}
		}
		if javaScript == "" {
			javaScript = fmt.Sprintf("it.%s.%s.%s()",
				imitativeTypingConfig.GetJavaScriptTableNameForSystem(),
				imitativeTypingConfig.GetJavaScriptTableNameForSystemFunctions(),
				id)
		}
		err := duktapeContext.PevalString(javaScript)
		if err != nil {
			panic(err)
		}

	}
	SaveFile(from, to)
}

type IDForSortStruct struct {
	id    string
	order int32
}

type IDForSort []*IDForSortStruct

func (m IDForSort) Len() int { return len(m) }
func (m IDForSort) Less(i, j int) bool {
	return m[i].order < m[j].order
}
func (m IDForSort) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
