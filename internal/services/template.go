package services

type TemplateService struct {
}

var templateSvc = &TemplateService{}

func GetTemplateSvc() *TemplateService {
	return templateSvc
}
