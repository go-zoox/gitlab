package merge_request

type IsConflictRequest struct {
	ProjectID      int64 `json:"project_id"`
	MergeRequestID int64 `json:"merge_request_id"`
}

func IsConflict(req *IsConflictRequest) (bool, error) {
	mr, err := Get(&GetRequest{
		ProjectID:      req.ProjectID,
		MergeRequestID: req.MergeRequestID,
	})
	if err != nil {
		return false, err
	}

	if mr.HasConflicts {
		return true, nil
	}

	return false, nil
}
