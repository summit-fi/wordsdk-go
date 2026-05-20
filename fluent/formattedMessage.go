package fluent

type FormattedMessage struct {
	Value      *string
	Attributes map[string]string
}

// Text returns the value of the message, or the fallback if the message is nil or has no value.
func (m *FormattedMessage) Text(fallback string) string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return fallback
}

// Attr returns the value of the attribute with the given name, or the fallback if the attribute is not found.
func (m *FormattedMessage) Attr(name, fallback string) string {
	if m != nil {
		if v, ok := m.Attributes[name]; ok {
			return v
		}
	}
	return fallback
}
