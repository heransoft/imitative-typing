package imitative_typing

import (
	"encoding/json"
	"fmt"
	"gopkg.in/olebedev/go-duktape.v2"
)

func RegisterFileInsert() {
	javaScriptContext.PushGoFunction(imitativeTypingConfig.GetJavaScriptTableNameForFile()+"."+imitativeTypingConfig.GetJavaScriptFunctionNameForFileInsert(),
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
		for id := range fileContext.JavaScript {
			ids = append(ids, id)
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
				panic(fmt.Sprintf("file(%s:%v) not exist JavaScript id(%s)",
					from,
					imitativeTypingContext.JavaScript,
					id))
			}
		}
		err := duktapeContext.PevalString(javaScript)
		if err != nil {
			panic(err)
		}
	}
	SaveFile(from, to)

}
