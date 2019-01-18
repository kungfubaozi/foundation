package review

type Service interface {
}

type reviewService struct {
}

func (svc *reviewService) GetRepo() repository {
	return &reviewRepository{}
}
