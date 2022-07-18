package group

func Search(keyword string) (*[]Group, error) {
	return List(&ListRequest{
		Search: keyword,
	})
}
