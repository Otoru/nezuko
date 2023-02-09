package healthcheck

import (
	"go.uber.org/zap"
)

type Runner struct {
	Name       string
	Evaluation func() error
}

type HealthCheck struct {
	services []Runner
}

func (service *HealthCheck) Append(runner Runner) {
	service.services = append(service.services, runner)
}

func (service *HealthCheck) Validate(logger *zap.Logger) error {
	for _, item := range service.services {
		err := item.Evaluation()

		logger.Debug("Validating status of service", zap.String("service", item.Name), zap.Error(err))

		if err != nil {
			return err
		}
	}

	return nil
}

func New() *HealthCheck {
	instance := &HealthCheck{}
	return instance
}
