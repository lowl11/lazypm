package pkg_event

import "lazypm/src/data/models"

func (event *Event) Get() *models.Package {
	return event.pkg
}
