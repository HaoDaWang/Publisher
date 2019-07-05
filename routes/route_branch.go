package routes

import (
	"fmt"
	"net/http"
	"publisher/JSON"
	"publisher/utils"
)

func GetBranches(res http.ResponseWriter, req *http.Request){
	projectUrl := utils.GitProject("/Users/haodawang/Documents/repositories/edu-saas-workbench")
	data := JSON.H{
		"branches": projectUrl.GetRemoteBranchNames(),
	}

	fmt.Fprintln(res, data.Stringify())
}