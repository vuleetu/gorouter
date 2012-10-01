package router

import (
    "regexp"
    "strings"
    "log"
)

const DEFAULT_ACTION string =  "Default"
const ACTION_SUFFIX string = "Action"

type RegRule struct{}

func (r RegRule) Parse(uri string) (string, string) {
    controller := ""
    action := DEFAULT_ACTION
    reg, err := regexp.Compile("^/([^/]*)")

    if err != nil {
        return controller, action
    }

    matches := reg.FindStringSubmatch(uri)

    controller = matches[1]

    if 0 == len(matches) {
        return controller, action
    }

    reg, err = regexp.Compile("^/[^/]+/([^/]+)")

    if err != nil {
        return controller, action
    }

    matches1 := reg.FindStringSubmatch(uri)

    if 0 == len(matches) {
        return controller, action
    }

    if 2 == len(matches1) {
        action = matches1[1]
    }

    action = strings.ToLower(action)
    log.Printf("Action is %s\n", action)
    action = strings.ToUpper(string(action[0])) + string(action[1:])

    return controller, action + ACTION_SUFFIX
}
