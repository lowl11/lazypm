package skeleton_event

import "lazypm/src/data/models"

func (event *Event) CreateObjects(config *models.ProjectConfig) error {
	defer event.skeleton.Reset()
	for event.skeleton.Next() {
		obj := event.skeleton.Value()

		if obj.IsFolder {
			if err := event.createObject(&obj, config); err != nil {
				return err
			}
		} else {
			if err := event.createFile(&obj, config); err != nil {
				return err
			}
		}
	}

	return nil
}
