package message

func pipelineMessage() map[string]map[string]string {
	return map[string]map[string]string{
		"TeamId": {
			"required": "Please enter a team id",
		},
		"Name": {
			"required": "Please enter a pipeline name",
		},
		"Spec": {
			"required": "Please enter a pipeline spec",
		},
		"Status": {
			"required": "Please enter a pipeline status",
		},
		"Finished": {
			"required": "Please select a task to perform on finished",
		},
		"Failed": {
			"required": "Please select a task to perform on failure",
		},
		"Overlap": {
			"required": "Please select whether to repeat execution",
		},
	}
}
