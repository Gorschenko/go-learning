package stats

import eventsbus "test/packages/events"

type StatsServiceDeps struct {
	EventBus        *eventsbus.EventBus
	StatsRepository *StatsRepository
}

type StatsService struct {
	EventBus        *eventsbus.EventBus
	StatsRepository *StatsRepository
}

func NewStatsService(deps *StatsServiceDeps) *StatsService {
	return &StatsService{
		EventBus:        deps.EventBus,
		StatsRepository: deps.StatsRepository,
	}
}

func (service *StatsService) AddClick() {
	for message := range service.EventBus.Subscribe() {
		if message.Type == eventsbus.EventLinkVisited {
			_, ok := message.Data.(uint)
			if !ok {
				continue
			}
			service.StatsRepository.AddClick(message.Data.(uint))
		}
	}
}
