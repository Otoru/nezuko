package healthcheck

import (
	"go.uber.org/zap"
)

type runner struct {
	Name       string
	Evaluation func() error
}

type Service struct {
	services []runner
}

func (service *Service) Append(name string, evaluation func() error) {
	service.services = append(service.services, runner{Name: name, Evaluation: evaluation})
}

func (service *Service) Validate(logger *zap.Logger) error {
	for _, item := range service.services {
		err := item.Evaluation()

		logger.Debug("Validating status of service", zap.String("service", item.Name), zap.Error(err))

		if err != nil {
			return err
		}
	}

	return nil
}

func New() *Service {
	instance := &Service{}
	return instance
}
