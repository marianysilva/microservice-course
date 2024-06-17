package domain

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) SubjectDeletedEvent(subject *DeletedSubject) error {
	event := struct {
		Event       string
		SubjectUUID uuid.UUID
	}{
		Event:       "SubjectDeleted",
		SubjectUUID: subject.UUID,
	}

	marshalledSubject, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("service can't send delete subject event: %w", err)
	}
	s.queue.Publish(marshalledSubject)
	return nil
}
