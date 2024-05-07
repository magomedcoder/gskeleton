package router

type GrpcMethod struct {
	Name string
}

type GrpcMethodService struct {
	publicMethods []*GrpcMethod
}

func NewGrpMethodsService() *GrpcMethodService {
	return &GrpcMethodService{
		publicMethods: []*GrpcMethod{
			{
				Name: "/auth.AuthService/Login",
			},
		},
	}
}

func (s *GrpcMethodService) IsPublicMethod(method string) bool {
	isPublic := false
	for _, route := range s.publicMethods {
		if route.Name == method {
			isPublic = true
		}
	}
	return isPublic
}
